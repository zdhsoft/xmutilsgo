package xm

import "encoding/json"

// Struct2Json 生成紧凑的Json字符串
func Struct2Json(paramStruct interface{}) (string, error) {
	bytes, err := json.Marshal(paramStruct)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Struct2JsonIndent 生成格式化的Json字符串 - 注意 paramStruct必须是指针类型
func Struct2JsonIndent(paramStruct interface{}) (string, error) {
	bytes, err := json.MarshalIndent(paramStruct, "", "  ")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Json2Struct 字符串转struct - 注意 paramStruct必须是指针类型
func Json2Struct(jsonStr string, paramStruct interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), paramStruct)
	if err != nil {
		return err
	}
	return nil
}

// Struct2JsonString 生成紧凑的Json字符串 如果有错误返回空字符串
func Struct2JsonString(paramStruct interface{}) string {
	str, err := Struct2Json(paramStruct)
	if err != nil {
		return ""
	} else {
		return str
	}
}
