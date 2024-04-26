package xm

import "encoding/json"

// 生成紧凑的Json字符串
func Struct2Json(paramStruct interface{}) (string, error) {
	bytes, err := json.Marshal(paramStruct)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// 生成格式化的Json字符串
func Struct2JsonIndent(obj interface{}) (string, error) {
	bytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// 字符串转struct
func Json2Struct(jsonStr string, pStruct interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), pStruct)
	if err != nil {
		return err
	}
	return nil
}

// 字符串截取
// 当字符串超过指定的长度，则截取指定的长度(中文算一个字符)
func TruncateString(paramValue string, paramMaxLen int) string {
	if len(paramValue) > paramMaxLen {
		runes := []rune(paramValue)
		runesCnt := len(runes)
		if runesCnt > paramMaxLen {
			return string(runes[:paramMaxLen])
		} else {
			return paramValue
		}
	}
	return paramValue
}
