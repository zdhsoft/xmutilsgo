package xm

import (
	"net/url"
	"strings"
)

/*
StringPad 字符串位数不足补指定字符
  - paramStr 源字符串
  - paramMinLen 最小长度
  - paramChar 指定的替换字符
  - return string 返回补足后的字符串
*/
func StringPad(paramStr string, paramMinLen int, paramChar string) string {
	strLen := len(paramStr)
	if strLen < paramMinLen {
		return strings.Repeat("0", paramMinLen-strLen) + paramStr
	}
	return paramStr
}

/*
StringCat 链接字符串
  - paramArgs 要链接的字符串列表
  - return string 返回链接后的字符串
*/
func StringCat(paramArgs ...string) string {
	builder := &strings.Builder{}
	builder.Grow(builder.Len() + 128)
	for _, s := range paramArgs {
		builder.WriteString(s)
	}
	return builder.String()
}

// Trim 去除前后空格
func Trim(paramStr string) string {
	return strings.TrimSpace(paramStr)
}

/*
ReplaceDomain 替换链接中的域名
链接可以绝对地址，也可以是相对地址
如果newDomain中，没有带https:// 或http:// 则默认视为https:// 如果带了，则以带入的
  - 示例: [原:/api] [新域名:www.google.com]=> https://www.google.com/api
  - 示例: [原:/api?aaaa=bbb] [新域名:http://www.google.com]=> http://www.google.com/api?aaaa=bbb
  - 示例: [原:https://www.com.cn/api?name=999&c=测试] [新域名:https://www.google.com]=> https://www.google.com/api?name=999&c=测试
*/
func ReplaceDomain(paramOldURL string, paramNewDomain string) string {
	u, err := url.Parse(paramOldURL)
	if err != nil {
		return ""
	}
	h, err1 := url.Parse(paramNewDomain)
	if err1 != nil {
		return ""
	}
	if h.Scheme != "" {
		u.Scheme = h.Scheme
		u.Host = h.Host
	} else {
		u.Scheme = "https"
		u.Host = paramNewDomain
	}
	return u.String()
}

// TruncateString 字符串截取
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
