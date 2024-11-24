package xm

import (
	"fmt"
	"strconv"
	"time"
)

// 将Time转换为日期号码， 该日期的格式为YYYYMMDD的8位整数
//   - 这个存在不确定性，主要有是这个使用的是默认时区造成的
//   - 建议使用FormatDateNumForBeijing函数指定了时区
func FormatDateNum(date time.Time) int {
	year, month, day := date.Date()
	return year*10000 + int(month)*100 + day
}

// 将Time转换为北京时区日期号码， 该日期的格式为YYYYMMDD的8位整数,
func FormatDateNumForBeijing(date time.Time) int {
	bj := date.In(beijingLoc)
	year, month, day := bj.Date()
	return year*10000 + int(month)*100 + day
}

// 将日期号码转换为Time， 该日期的格式为YYYYMMDD的8位整数
func ParseDateNum(dateNum int) (time.Time, error) {
	return time.Parse("20060102", strconv.Itoa(dateNum))
}

// 将日期号码转换为Time， 该日期的格式为YYYYMMDD的8位整数, 该日期的时区为北京时区
func ParseDateNumForBeijing(dateNum int) (time.Time, error) {
	return time.ParseInLocation("20060102", strconv.Itoa(dateNum), beijingLoc)
}

// 取指定日期的后一天的日期号码 如：20210101 -> 20210102
func NextDateNum(dateNum int) (int, error) {
	date, err := ParseDateNum(dateNum)
	if err != nil {
		return 0, err
	}
	nextDate := date.AddDate(0, 0, 1)
	return FormatDateNum(nextDate), nil
}

// 取指定日期的后n天的日期号码 如：20210101 -> 20210102
func NextNDaysDateNum(dateNum int, nDays int) (int, error) {
	date, err := ParseDateNum(dateNum)
	if err != nil {
		return 0, err
	}
	nextDate := date.AddDate(0, 0, nDays)
	return FormatDateNum(nextDate), nil
}

// 取指定日期的后一天的日期号码 如：20210101 -> 20210102
func NextDateNumForBeijing(dateNum int) (int, error) {
	date, err := ParseDateNumForBeijing(dateNum)
	if err != nil {
		return 0, err
	}
	nextDate := date.AddDate(0, 0, 1)
	return FormatDateNumForBeijing(nextDate), nil
}

// 取当前日期的日期号码
func GetNowDateNum() int {
	return FormatDateNum(time.Now())
}

// 取当前日期的北京时区日期号码
func GetNowDateNumForBeijing() int {
	return FormatDateNumForBeijing(time.Now())
}

// 日期号码转日期字符串 YYYY-MM-DD
func DateNum2DateStr(dateNum int, splitChar string) string {
	Day := dateNum % 100
	dateNum /= 100
	Month := dateNum % 100
	Year := dateNum / 100
	return fmt.Sprintf("%04d%s%02d%s%02d", Year, splitChar, Month, splitChar, Day)
}
