package xm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// DefaultHTTPTimeout 默认的 HTTP 请求超时时间
const DefaultHTTPTimeout = 60 * time.Second

// structToQueryParams 将带有json标记的结构体转换为url.Values
//   - paramInput 输入结构指或结构体指针
//   - paramAllocOmitempty=true时处理omitempty标记，否则忽略改标记
//   - 返回url.Values，错误信息
//
// 注意：
//   - url.Values的Encode()方法, 会对 key做排序
func StructToQueryParams(paramInput interface{}, paramAllocOmitempty bool) (url.Values, error) {
	values := url.Values{}
	v := reflect.ValueOf(paramInput)

	// 处理指针类型
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return values, nil
		}
		v = v.Elem()
	}

	// 仅处理结构体
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input must be a struct")
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)

		// 解析json标签
		jsonTag := fieldType.Tag.Get("json")
		jsonName, omitempty := parseJSONTag(jsonTag, fieldType.Name)

		// 跳过"-"标记字段
		if jsonName == "-" {
			continue
		}

		// 处理omitempty逻辑
		if paramAllocOmitempty && omitempty && isEmptyValue(fieldValue) {
			continue
		}

		// 值转换
		strValue, err := valueToString(fieldValue)
		if err != nil {
			return nil, fmt.Errorf("字段[%s]转换失败: %v", jsonName, err)
		}

		if strValue != "" {
			values.Add(jsonName, strValue)
		}
	}
	return values, nil
}

// parseJSONTag 解析json标签，返回字段名和omitempty标记
func parseJSONTag(paramTag, paramFieldName string) (string, bool) {
	if paramTag == "" {
		return strings.ToLower(paramFieldName), false
	}

	parts := strings.Split(paramTag, ",")
	name := parts[0]

	switch {
	case name == "-":
		return "-", false
	case name == "":
		name = strings.ToLower(paramFieldName)
	}

	omitempty := false
	for _, part := range parts[1:] {
		if strings.TrimSpace(part) == "omitempty" {
			omitempty = true
		}
	}

	return name, omitempty
}

// isEmptyValue 判断是否零值（支持嵌套指针）
func isEmptyValue(paramValue reflect.Value) bool {
	// 解除指针嵌套
	for paramValue.Kind() == reflect.Ptr || paramValue.Kind() == reflect.Interface {
		if paramValue.IsNil() {
			return true
		}
		paramValue = paramValue.Elem()
	}

	switch paramValue.Kind() {
	case reflect.String:
		return paramValue.Len() == 0
	case reflect.Array, reflect.Slice, reflect.Map:
		return paramValue.Len() == 0
	case reflect.Bool:
		return !paramValue.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return paramValue.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return paramValue.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return paramValue.Float() == 0
	case reflect.Struct:
		// 特殊处理time.Time
		if t, ok := paramValue.Interface().(time.Time); ok {
			return t.IsZero()
		}
		return false
	default:
		return false
	}
}

// 连接 Get 的请求参数
func joinGetParamsString(paramPath string, paramQueryString string) string {
	queryString := Trim(paramQueryString)
	if queryString != "" {
		return paramPath + "?" + queryString
	} else {
		return paramPath
	}
}

// valueToString 值类型转换（支持常见类型和嵌套指针）
func valueToString(paramValue reflect.Value) (string, error) {
	// 处理指针嵌套
	for paramValue.Kind() == reflect.Ptr || paramValue.Kind() == reflect.Interface {
		if paramValue.IsNil() {
			return "", nil
		}
		paramValue = paramValue.Elem()
	}

	switch paramValue.Kind() {
	case reflect.String:
		return paramValue.String(), nil
	case reflect.Bool:
		return strconv.FormatBool(paramValue.Bool()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(paramValue.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(paramValue.Uint(), 10), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(paramValue.Float(), 'f', -1, 64), nil
	case reflect.Struct:
		if t, ok := paramValue.Interface().(time.Time); ok {
			return t.Format(time.RFC3339), nil
		}
		return Struct2Json(paramValue.Interface())
	case reflect.Slice, reflect.Array:
		var parts []string
		for i := 0; i < paramValue.Len(); i++ {
			part, err := valueToString(paramValue.Index(i))
			if err != nil {
				return "", err
			}
			parts = append(parts, part)
		}
		return strings.Join(parts, ","), nil
	default:
		return Struct2Json(paramValue.Interface())
	}
}

// PostRequestByOrigin 原始的POST 请求，上传 JSON 数据并返回 JSON 响应
//   - paramURL 请求 URL 要求是完整路径，如 https://example.com/api/v1/test 或 https://example.com：8080/api/v1/test
//   - paramHeaders 请求头，map[string]string 类型 无 header 则传 nil
//   - paramBody 已经要发送的 JSON 数据，[]byte 类型
//   - paramTimeout 请求超时时间，如果为0则使用默认值60秒
//   - 返回值：响应数据，[]byte 类型
func PostRequestByOrigin(paramURL string, paramHeaders map[string]string, paramBody []byte, paramTimeout time.Duration) ([]byte, error) {
	// 创建一个新的 POST 请求
	req, err := http.NewRequest("POST", paramURL, bytes.NewBuffer(paramBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// 设置请求头
	for key, value := range paramHeaders {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/json")

	// 设置超时时间
	if paramTimeout == 0 {
		paramTimeout = DefaultHTTPTimeout
	}
	client := &http.Client{
		Timeout: paramTimeout,
	}

	// 执行请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应 body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	return respBody, nil
}

// GetRequestByOrigin 原始的Get 请求，返回 JSON 响应
//   - paramURL 请求 URL 要求是完整路径，如 https://example.com/api/v1/test 或 https://example.com：8080/api/v1/test
//   - paramHeaders 请求头，map[string]string 类型 无 header 则传 nil
//   - paramQueryString ?后面的查询参数，string 类型
//   - paramTimeout 请求超时时间，如果为0则使用默认值60秒
//   - 返回值：响应数据，[]byte 类型
func GetRequestByOrigin(paramURL string, paramHeaders map[string]string, paramQueryString string, paramTimeout time.Duration) ([]byte, error) {
	// 构造完整的 URL，包含查询参数
	finalURL := joinGetParamsString(paramURL, paramQueryString)
	// 创建一个新的 GET 请求
	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// 设置请求头
	for key, value := range paramHeaders {
		req.Header.Set(key, value)
	}

	// 设置超时时间
	if paramTimeout == 0 {
		paramTimeout = DefaultHTTPTimeout
	}
	client := &http.Client{
		Timeout: paramTimeout,
	}

	// 执行请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应 body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	return respBody, nil
}

// PostRequestBy2Map POST 请求，上传 JSON 数据并返回 map 响应
//   - paramURL 请求 URL 要求是完整路径，如 https://example.com/api/v1/test 或 https://example.com：8080/api/v1/test
//   - paramHeaders 请求头，map[string]string 类型 无 header 则传 nil
//   - paramBody 要发送的 JSON 数据，任意类型，会被转换为 JSON
//   - paramTimeout 请求超时时间，如果为0则使用默认值60秒
//   - 返回值：响应数据，map[string]interface{} 类型
func PostRequestBy2Map(paramURL string, paramHeaders map[string]string, paramBody interface{}, paramTimeout time.Duration) (map[string]interface{}, error) {
	// 将请求体转换为 JSON
	jsonData, err := json.Marshal(paramBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	// 发送请求
	respBody, err := PostRequestByOrigin(paramURL, paramHeaders, jsonData, paramTimeout)
	if err != nil {
		return nil, err
	}

	// 解析响应
	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}
	return result, nil
}

// PostRequestBy2Struct POST 请求，上传 JSON 数据并返回结构体响应
//   - paramURL 请求 URL 要求是完整路径，如 https://example.com/api/v1/test 或 https://example.com：8080/api/v1/test
//   - paramHeaders 请求头，map[string]string 类型 无 header 则传 nil
//   - paramBody 要发送的 JSON 数据，任意类型，会被转换为 JSON
//   - paramReply 响应数据的结构体指针
//   - paramTimeout 请求超时时间，如果为0则使用默认值60秒
func PostRequestBy2Struct(paramURL string, paramHeaders map[string]string, paramBody interface{}, paramReply interface{}, paramTimeout time.Duration) error {
	// 将请求体转换为 JSON
	jsonData, err := json.Marshal(paramBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err)
	}

	// 发送请求
	respBody, err := PostRequestByOrigin(paramURL, paramHeaders, jsonData, paramTimeout)
	if err != nil {
		return err
	}

	// 解析响应
	if err := json.Unmarshal(respBody, paramReply); err != nil {
		return fmt.Errorf("failed to unmarshal response: %v", err)
	}
	return nil
}

// GetRequestByMap2Map GET 请求，使用 map 参数并返回 map 响应
//   - paramURL 请求 URL 要求是完整路径，如 https://example.com/api/v1/test 或 https://example.com：8080/api/v1/test
//   - paramHeaders 请求头，map[string]string 类型 无 header 则传 nil
//   - params 查询参数，map[string]interface{} 类型
//   - paramTimeout 请求超时时间，如果为0则使用默认值60秒
//   - 返回值：响应数据，map[string]interface{} 类型
func GetRequestByMap2Map(paramURL string, paramHeaders map[string]string, params map[string]interface{}, paramTimeout time.Duration) (map[string]interface{}, error) {
	// 将参数转换为查询字符串
	values := url.Values{}
	for key, value := range params {
		values.Add(key, fmt.Sprintf("%v", value))
	}
	queryString := values.Encode()

	// 发送请求
	respBody, err := GetRequestByOrigin(paramURL, paramHeaders, queryString, paramTimeout)
	if err != nil {
		return nil, err
	}

	// 解析响应
	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}
	return result, nil
}

// GetRequestByMap2Struct GET 请求，使用 map 参数并返回结构体响应
//   - paramURL 请求 URL 要求是完整路径，如 https://example.com/api/v1/test 或 https://example.com：8080/api/v1/test
//   - paramHeaders 请求头，map[string]string 类型 无 header 则传 nil
//   - params 查询参数，map[string]string 类型
//   - paramReply 响应数据的结构体指针
//   - paramTimeout 请求超时时间，如果为0则使用默认值60秒
func GetRequestByMap2Struct(paramURL string, paramHeaders map[string]string, params map[string]string, paramReply interface{}, paramTimeout time.Duration) error {
	// 将参数转换为查询字符串
	values := url.Values{}
	for key, value := range params {
		values.Add(key, value)
	}
	queryString := values.Encode()

	// 发送请求
	respBody, err := GetRequestByOrigin(paramURL, paramHeaders, queryString, paramTimeout)
	if err != nil {
		return err
	}

	// 解析响应
	if err := json.Unmarshal(respBody, paramReply); err != nil {
		return fmt.Errorf("failed to unmarshal response: %v", err)
	}
	return nil
}

// GetRequestByStruct2Map GET 请求，使用结构体参数并返回 map 响应
//   - paramURL 请求 URL 要求是完整路径，如 https://example.com/api/v1/test 或 https://example.com：8080/api/v1/test
//   - paramHeaders 请求头，map[string]string 类型 无 header 则传 nil
//   - params 查询参数，结构体类型
//   - paramTimeout 请求超时时间，如果为0则使用默认值60秒
//   - 返回值：响应数据，map[string]interface{} 类型
func GetRequestByStruct2Map(paramURL string, paramHeaders map[string]string, params interface{}, paramTimeout time.Duration) (map[string]interface{}, error) {
	// 将结构体转换为查询字符串
	values, err := StructToQueryParams(params, true)
	if err != nil {
		return nil, fmt.Errorf("failed to convert struct to query params: %v", err)
	}
	queryString := values.Encode()

	// 发送请求
	respBody, err := GetRequestByOrigin(paramURL, paramHeaders, queryString, paramTimeout)
	if err != nil {
		return nil, err
	}

	// 解析响应
	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}
	return result, nil
}

// GetRequestByStruct2Struct GET 请求，使用结构体参数并返回结构体响应
//   - paramURL 请求 URL 要求是完整路径，如 https://example.com/api/v1/test 或 https://example.com：8080/api/v1/test
//   - paramHeaders 请求头，map[string]string 类型 无 header 则传 nil
//   - params 查询参数，结构体类型
//   - paramReply 响应数据的结构体指针
//   - paramTimeout 请求超时时间，如果为0则使用默认值60秒
func GetRequestByStruct2Struct(paramURL string, paramHeaders map[string]string, params interface{}, paramReply interface{}, paramTimeout time.Duration) error {
	// 将结构体转换为查询字符串
	values, err := StructToQueryParams(params, true)
	if err != nil {
		return fmt.Errorf("failed to convert struct to query params: %v", err)
	}
	queryString := values.Encode()

	// 发送请求
	respBody, err := GetRequestByOrigin(paramURL, paramHeaders, queryString, paramTimeout)
	if err != nil {
		return err
	}

	// 解析响应
	if err := json.Unmarshal(respBody, paramReply); err != nil {
		return fmt.Errorf("failed to unmarshal response: %v", err)
	}
	return nil
}
