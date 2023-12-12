package xm

import (
	"math/rand"
	"strconv"
)

// 有符号整数类型
type SignedInteger interface {
	~int8 | ~int16 | ~int32 | ~int | ~int64
}

// 无符号整数类型
type UnsignedInterger interface {
	~uint8 | ~uint16 | ~uint32 | ~uint | ~uint64
}

// 所有整数类型
type Integer interface {
	SignedInteger | UnsignedInterger
}

/*
整数转十进制字符串
  - paramValue 要变成字符串的整数
  - return string
*/
func Int2String[T SignedInteger](paramValue T) string {
	v := int64(paramValue)
	return strconv.FormatInt(v, 10)
}

/*
整数转十进制字符串, 并指定最小位数，不足补0
  - paramValue 要变成字符串的整数
  - paramMinLen 最小的位数
  - return string
*/
func Int2StringPad[T SignedInteger](paramValue T, paramMinLen int) string {
	return StringPad(Int2String(paramValue), paramMinLen, "0")
}

/*
整数转指定进制字符串
  - paramValue 要变成字符串的整数
  - paramBase 指定的进制 改值有效访问2-36
  - return string
*/
func Int2StringBase[T SignedInteger](paramValue T, paramBase int) string {
	v := int64(paramValue)
	return strconv.FormatInt(v, paramBase)
}

/*
整数转指定进制字符串，并指定最小位数，不足补0
  - paramValue 要变成字符串的整数
  - paramBase 指定的进制 改值有效访问2-36
  - paramMinLen 最小的位数
  - return string
*/
func Int2StringBasePad[T SignedInteger](paramValue T, paramBase int, paramMinLen int) string {
	return StringPad(Int2StringBase(paramValue, paramBase), paramMinLen, "0")
}

/*
无符号整数转十进制字符串
  - paramValue 要变成字符串的整数
  - return string
*/
func UInt2String[T UnsignedInterger](paramValue T) string {
	v := int64(paramValue)
	return strconv.FormatInt(v, 10)
}

/*
无符号整数转十进制字符串, 并指定最小位数，不足补0
  - paramValue 要变成字符串的整数
  - paramMinLen 最小的位数
  - return string
*/
func UInt2StringPad[T UnsignedInterger](paramValue T, paramMinLen int) string {
	return StringPad(UInt2String(paramValue), paramMinLen, "0")
}

/*
无符号整数转指定进制字符串
  - paramValue 要变成字符串的整数
  - paramBase 指定的进制 改值有效访问2-36
  - return string
*/
func UInt2StringBase[T UnsignedInterger](paramValue T, paramBase int) string {
	v := int64(paramValue)
	return strconv.FormatInt(v, paramBase)
}

/*
无符号整数转指定进制字符串，并指定最小位数，不足补0
  - paramValue 要变成字符串的整数
  - paramBase 指定的进制 改值有效访问2-36
  - paramMinLen 最小的位数
  - return string
*/
func UInt2StringBasePad[T UnsignedInterger](paramValue T, paramBase int, paramMinLen int) string {
	return StringPad(UInt2StringBase(paramValue, paramBase), paramMinLen, "0")
}

/*
随机一个[minValue, maxValue]之间的整数
  - minValue 最小值
  - maxValue 最大值
  - return int 随机值
*/
func RandomIntScope(minValue int, maxValue int) int {
	if minValue > maxValue {
		minValue, maxValue = maxValue, minValue
	} else if minValue == maxValue {
		return minValue
	}
	return rand.Intn(maxValue-minValue+1) + minValue
}

// 64位整数，十进制反转
func ReverseInt64(paramValue int64) int64 {
	var result int64
	num := paramValue
	for num != 0 {
		result = result*10 + num%10
		num /= 10
	}
	return result
}
