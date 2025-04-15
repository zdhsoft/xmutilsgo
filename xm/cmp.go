package xm

import "time"

// CanOrderType 所有可比较的类型
type CanOrderType interface {
	Integer | Float | ~string
}

// IntCmp 比较两个整数
// CanOrderType 所有可比较的类型主要有整数、浮点数、字符串
// 返回值 1: paramValue1 > paramValue2, -1: paramValue1 < paramValue2, 0: paramValue1 == paramValue2
func Cmp[T CanOrderType](paramValue1, paramValue2 T) int {
	if paramValue1 > paramValue2 {
		return 1
	} else if paramValue1 < paramValue2 {
		return -1
	} else {
		return 0
	}
}

// CmpTime 比较两个时间
// 返回值 1: paramValue1 > paramValue2, -1: paramValue1 < paramValue2, 0: paramValue1 == paramValue2
func CmpTime(paramValue1, paramValue2 time.Time) int {
	if paramValue1.After(paramValue2) {
		return 1
	} else if paramValue1.Before(paramValue2) {
		return -1
	} else {
		return 0
	}
}

// CmpBool 比较两个布尔值  false < true
// 返回值 1: paramValue1 > paramValue2, -1: paramValue1 < paramValue2, 0: paramValue1 == paramValue2
func CmpBool(paramValue1, paramValue2 bool) int {
	if paramValue1 == paramValue2 {
		return 0
	} else if paramValue1 {
		return 1
	} else {
		return -1
	}
}
