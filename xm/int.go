package xm

import (
	"errors"
	"math/rand"
	"strconv"
)

// SignedInteger 有符号整数类型
type SignedInteger interface {
	~int8 | ~int16 | ~int32 | ~int | ~int64
}

// UnsignedInteger 无符号整数类型
type UnsignedInteger interface {
	~uint8 | ~uint16 | ~uint32 | ~uint | ~uint64
}

// Integer 所有整数类型
type Integer interface {
	SignedInteger | UnsignedInteger
}

// Float 所有浮点数类型
type Float interface {
	~float32 | ~float64
}


// Int2String /*
func Int2String[T SignedInteger](paramValue T) string {
	v := int64(paramValue)
	return strconv.FormatInt(v, 10)
}

/*
I 简化整数转字符串
  - paramValue 要变成字符串的整数
  - return string
*/
func I[T SignedInteger](paramValue T) string {
	return Int2String(paramValue)
}

/*
Int2StringPad 整数转十进制字符串, 并指定最小位数，不足补0
  - paramValue 要变成字符串的整数
  - paramMinLen 最小的位数
  - return string
*/
func Int2StringPad[T SignedInteger](paramValue T, paramMinLen int) string {
	return StringPad(Int2String(paramValue), paramMinLen, "0")
}

/*
Int2StringBase 整数转指定进制字符串
  - paramValue 要变成字符串的整数
  - paramBase 指定的进制 改值有效访问2-36
  - return string
*/
func Int2StringBase[T SignedInteger](paramValue T, paramBase int) string {
	v := int64(paramValue)
	return strconv.FormatInt(v, paramBase)
}

/*
Int2StringBasePad 整数转指定进制字符串，并指定最小位数，不足补0
  - paramValue 要变成字符串的整数
  - paramBase 指定的进制 改值有效访问2-36
  - paramMinLen 最小的位数
  - return string
*/
func Int2StringBasePad[T SignedInteger](paramValue T, paramBase int, paramMinLen int) string {
	return StringPad(Int2StringBase(paramValue, paramBase), paramMinLen, "0")
}

/*
UInt2String 无符号整数转十进制字符串
  - paramValue 要变成字符串的整数
  - return string
*/
func UInt2String[T UnsignedInteger](paramValue T) string {
	v := uint64(paramValue)
	return strconv.FormatUint(v, 10)
}

/*
U 简化版无符号整数转十进制字符串
  - paramValue 要变成字符串的整数
  - return string
*/
func U[T UnsignedInteger](paramValue T) string {
	v := uint64(paramValue)
	return strconv.FormatUint(v, 10)
}

/*
UInt2StringPad 无符号整数转十进制字符串, 并指定最小位数，不足补0
  - paramValue 要变成字符串的整数
  - paramMinLen 最小的位数
  - return string
*/
func UInt2StringPad[T UnsignedInteger](paramValue T, paramMinLen int) string {
	return StringPad(UInt2String(paramValue), paramMinLen, "0")
}

/*
UInt2StringBase 无符号整数转指定进制字符串
  - paramValue 要变成字符串的整数
  - paramBase 指定的进制 改值有效访问2-36
  - return string
*/
func UInt2StringBase[T UnsignedInteger](paramValue T, paramBase int) string {
	v := uint64(paramValue)
	return strconv.FormatUint(v, paramBase)
}

/*
UInt2StringBasePad 无符号整数转指定进制字符串，并指定最小位数，不足补0
  - paramValue 要变成字符串的整数
  - paramBase 指定的进制 改值有效访问2-36
  - paramMinLen 最小的位数
  - return string
*/
func UInt2StringBasePad[T UnsignedInteger](paramValue T, paramBase int, paramMinLen int) string {
	return StringPad(UInt2StringBase(paramValue, paramBase), paramMinLen, "0")
}

/*
RandomIntScope 随机一个[paramMinValue, paramMaxValue]之间的整数
  - paramMinValue 最小值
  - paramMaxValue 最大值
  - return int 随机值
*/
func RandomIntScope(paramMinValue int, paramMaxValue int) int {
	if paramMinValue > paramMaxValue {
		paramMinValue, paramMaxValue = paramMaxValue, paramMinValue
	} else if paramMinValue == paramMaxValue {
		return paramMinValue
	}
	return rand.Intn(paramMaxValue-paramMinValue+1) + paramMinValue
}

// ReverseInt64 64位整数，十进制反转
func ReverseInt64(paramValue int64) int64 {
	var result int64
	num := paramValue
	for num != 0 {
		result = result*10 + num%10
		num /= 10
	}
	return result
}

// RandOneInArray 随机一个数组中的一个元素
func RandOneInArray[T any](paramArray []T) (*T, error) {
	cnt := len(paramArray)
	if cnt == 0 {
		err := errors.New("array is empty")
		return nil, err
	} else if cnt == 1 {
		return &paramArray[0], nil
	} else {
		idx := RandomIntScope(0, cnt-1)
		return &paramArray[idx], nil
	}
}

// RandNewByArray 用现有的数组生成一个新的随机数组
func RandNewByArray[T any](paramArray []T) []T {
	cnt := len(paramArray)
	retArr := append([]T{}, paramArray...)
	last := cnt - 1
	for i := 0; i < last; i++ {
		idx := RandomIntScope(i, last)
		if idx != i {
			retArr[i], retArr[idx] = retArr[idx], retArr[i]
		}
	}
	return retArr
}

// Str2Int 字符串转整数
func Str2Int[T Integer](paramStr string) (T, error) {
	v, err := strconv.ParseInt(paramStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return T(v), nil
}

// Str2UInt 字符串转无符号整数
func Str2UInt[T UnsignedInteger](paramStr string) (T, error) {
	v, err := strconv.ParseUint(paramStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return T(v), nil
}

// Str2Float 字符串转浮点数
func Str2Float(paramStr string) (float64, error) {
	return strconv.ParseFloat(paramStr, 64)
}

// Float2Str 浮点数转字符串
func Float2Str[T ~float32 | ~float64](paramFloat T) string {
	return strconv.FormatFloat(float64(paramFloat), 'f', -1, 64)
}
