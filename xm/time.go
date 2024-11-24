package xm

//仅演示获取当天指定时间的时间戳
import (
	"errors"
	"regexp"
	"time"
)

const (
	// 每天的小时数
	HOUR_BY_DAY = 24
	// 每秒的毫秒数
	MILLIS_BY_SECOND = 1000
	// 每分钟的毫秒数
	MILLIS_BY_MINUTE = 60 * MILLIS_BY_SECOND
	// 每小时的毫秒数
	MILLIS_BY_HOUR = 60 * MILLIS_BY_MINUTE
	// 每天的毫秒数
	MILLIS_BY_DAY = MILLIS_BY_HOUR * HOUR_BY_DAY
	// 每分钟的秒数
	SECOND_BY_MINUTE = 60
	// 每小时的秒数
	SECOND_BY_HOUR = 60 * SECOND_BY_MINUTE
	// 每天的秒数
	SECOND_BY_DAY = SECOND_BY_HOUR * HOUR_BY_DAY
	// 每小时的分钟数
	MINUTE_BY_HOUR = 60
	// 每天的分钟数
	MINUTE_BY_DAY = MINUTE_BY_HOUR * HOUR_BY_DAY
	// 北京时间时区
	TIMEZONE_BEIJING = -8
	// 北京时区的毫秒数
	MILLIS_BY_TIMEZONE_BEIJING = TIMEZONE_BEIJING * MILLIS_BY_HOUR
	// DateTime类型 格林威治时间
	DT_TYPE_UTC = 0
	// DateTime类型 北京时间
	DT_TYPE_BEIJING = 8
	// 北京时区名称
	ZONE_NAME_BEGIJING = "Asia/Shanghai"

	// 参数类型 日期 YYYY-MM-DD
	PARAM_TYPE_DATE = 1
	// 参数类型 日期时间 YYYY-MM-DD hh:mm:ss
	PARAM_TYPE_DATETIME = 2
	// 参数类型 错误
	PARAM_TYPE_ERROR = 0
)

var (
	dateRegex     = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	dateTimeRegex = regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`)
)

// 时区定义参考： https://jp.cybozu.help/general/zh/admin/list_systemadmin/list_localization/timezone.html
var beijingLoc *time.Location

func init() {
	beijingLoc, _ = time.LoadLocation(ZONE_NAME_BEGIJING)
}

// GetBeijingLocation 获取北京时区的Location
func GetBeijingLocation() *time.Location {
	return beijingLoc
}

// IsDateFormat 判断是否是日期格式 YYYY-MM-DD
func IsDateFormat(paramDate string) bool {
	return dateRegex.MatchString(paramDate)
}

// IsDateTimeFormat 判断是否是日期时间格式 YYYY-MM-DD hh:mm:ss
func IsDateTimeFormat(paramDateTime string) bool {
	return dateTimeRegex.MatchString(paramDateTime)
}

// GetNowMillis 取当前时间戳(毫秒)
func GetNowMillis() int64 {
	return time.Now().UnixMilli()
}

// GetNowSecond 取当前时间戳(秒）
func GetNowSecond() int64 {
	return time.Now().Unix()
}

// ParseDateTimeForBeijingMillis 解析北京格式的日期时间 单位毫秒
//   - paramDate 日期时间字符串 格式：2021-01-01 15:4:5
func ParseDateTimeForBeijingMillis(paramDate string) (int64, error) {
	startTime, err := time.ParseInLocation("2006-1-2 15:4:5", paramDate, beijingLoc)
	if err != nil {
		return 0, err
	}
	return startTime.UnixMilli(), nil
}

// ParseDateTimeForBeijingTime 解析北京格式的日期时间
//   - paramDate 日期时间字符串 格式：2021-01-01 15:4:5
func ParseDateTimeForBeijingTime(paramDate string) (time.Time, error) {
	return time.ParseInLocation("2006-1-2 15:4:5", paramDate, beijingLoc)
}

// ParseDateForBeijingTime 解析北京格式的日期时间
//   - paramDate 日期时间字符串 格式：2021-01-01
func ParseDateForBeijingTime(paramDate string) (time.Time, error) {
	return time.ParseInLocation("2006-1-2", paramDate, beijingLoc)
}

// ParseDateForBeijingMillis 解析北京格式的日期 单位毫秒
//   - paramDate 日期字符串 格式：2021-01-01
func ParseDateForBeijingMillis(paramDate string) (int64, error) {
	startTime, err := time.ParseInLocation("2006-1-2", paramDate, beijingLoc)
	if err != nil {
		return 0, err
	}
	return startTime.UnixMilli(), nil
}

// ParseDateTimeForBeijingSecond 解析北京格式的日期 单位秒
//   - paramDate 日期时间字符串 格式：2021-01-01 15:4:5
func ParseDateTimeForBeijingSecond(paramDate string) (int64, error) {
	stNow, err := ParseDateTimeForBeijingMillis(paramDate)
	stNow /= MILLIS_BY_SECOND // (stNow - stNow%MILLIS_BY_SECOND) / MILLIS_BY_SECOND
	return stNow, err
}

// ParseDateForBeijingSecond 解析北京格式的日期 单位秒
//   - paramDate 日期字符串 格式：2021-01-01
func ParseDateForBeijingSecond(paramDate string) (int64, error) {
	stNow, err := ParseDateForBeijingMillis(paramDate)
	stNow /= MILLIS_BY_SECOND // (stNow - stNow%MILLIS_BY_SECOND) / MILLIS_BY_SECOND
	return stNow, err
}

// Timestamp2BeijingDate 将时间戳转为北京时区的日期字符串 时间戳单位秒 YYYY-MM-DD
//   - paramTimestamp 时间戳 单位秒
//   - return 北京时区的日期字符串 格式：2021-01-01
func Timestamp2BeijingDate(paramTimestamp int64) string {
	return time.Unix(paramTimestamp, 0).In(beijingLoc).Format("2006-01-02")
}

// Timestamp2BeijingTime 将时间戳转为北京时区的时间字符串 时间戳单位秒 hh:mm:ss
//   - paramTimestamp 时间戳 单位秒
//   - return 北京时区的时间字符串 格式：15:04:05
func Timestamp2BeijingTime(paramTimestamp int64) string {
	return time.Unix(paramTimestamp, 0).In(beijingLoc).Format("15:04:05")
}

// Timestamp2BeijingDateTime 将时间戳转为北京时区的时间字符串 时间戳单位秒 YYYY-MM-DD hh:mm:ss
//   - paramTimestamp 时间戳 单位秒
//   - return 北京时区的日期时间字符串 格式：2021-01-01 15:04:05
func Timestamp2BeijingDateTime(paramTimestamp int64) string {
	return time.Unix(paramTimestamp, 0).In(beijingLoc).Format("2006-01-02 15:04:05")
}

// BeijingFormat 将时间转换为北京时区的格式字符串
//   - paramDateTime 时间
//   - paramFormat 格式化字符串 例如："2006-01-02 15:04:05"
//   - return 北京时区的格式化字符串
func BeijingFormat(paramDateTime time.Time, paramFormat string) string {
	return paramDateTime.In(beijingLoc).Format(paramFormat)
}

// BeijingDateString 将时间转换为北京时区的日期字符串 YYYY-MM-DD
//   - paramDateTime 时间
//   - return 北京时区的日期字符串 格式：2021-01-01
func BeijingDateString(paramDateTime time.Time) string {
	return BeijingFormat(paramDateTime, "2006-01-02")
}

// BeijingDateTimeString 将时间转换为北京时区的日期时间字符串 YYYY-MM-DD hh:mm:ss
//   - paramDateTime 时间
//   - return 北京时区的日期时间字符串 格式：2021-01-01 15:04:05
func BeijingDateTimeString(paramDateTime time.Time) string {
	return BeijingFormat(paramDateTime, "2006-01-02 15:04:05")
}

// BeijingTimeString 将时间转换为北京时区的时间字符串 hh:mm:ss
//   - paramDateTime 时间
//   - return 北京时区的时间字符串 格式：15:04:05
func BeijingTimeString(paramDateTime time.Time) string {
	return BeijingFormat(paramDateTime, "15:04:05")
}

// BeijingCompactDateString 将时间转换为北京时区的日期字符串(压缩版)
//   - paramDateTime 时间
//   - return 北京时区的日期字符串(压缩版) 格式：20210101
func BeijingCompactDateString(paramDateTime time.Time) string {
	return BeijingFormat(paramDateTime, "20060102")
}

// BeijingCompactDateTimeString 将时间转换为北京时区的日期时间字符串(压缩版)
//   - paramDateTime 时间
//   - return 北京时区的日期时间字符串(压缩版) 格式：20210101150405
func BeijingCompactDateTimeString(paramDateTime time.Time) string {
	return BeijingFormat(paramDateTime, "20060102150405")
}

// BeijingCompactTimeString 将时间转换为北京时区的时间字符串(压缩版)
//   - paramDateTime 时间
//   - return 北京时区的时间字符串(压缩版) 格式：150405
func BeijingCompactTimeString(paramDateTime time.Time) string {
	return BeijingFormat(paramDateTime, "150405")
}

// GetTimeOperationDayString 获取指定时间根据天数计算后得到的时间 YYYYMMDD格式字符串
//   - paramDateTime 时间
//   - paramDays 天数
func GetTimeOperationDayString(paramDateTime time.Time, paramDays int) string {
	paramDateTime = paramDateTime.AddDate(0, 0, paramDays)
	return BeijingFormat(paramDateTime, "20060102")
}

// DiffDaysFromTimestamp 计算两个时间戳对应的北京时间相差的天数，单位毫秒
//   - paramT1 时间戳1 单位毫秒
//   - paramT2 时间戳2 单位毫秒
//   - return 相差的天数 0表示同一天
func DiffDaysFromTimestamp(paramT1 int64, paramT2 int64) int64 {
	dt1 := DateTimeFromMillis(paramT1).ToBeijingZeroTime()
	dt2 := DateTimeFromMillis(paramT2).ToBeijingZeroTime()
	return (dt1.GetMillis() - dt2.GetMillis()) / MILLIS_BY_DAY
}

// DiffDaysFromTimestampSecond 计算两个时间戳对应的北京时间相差的天数，单位秒
//   - paramT1 时间戳1 单位秒
//   - paramT2 时间戳2 单位秒
//   - return 相差的天数 0表示同一天
func DiffDaysFromTimestampSecond(paramT1 int64, paramT2 int64) int64 {
	return DiffDaysFromTimestamp(paramT1*MILLIS_BY_SECOND, paramT2*MILLIS_BY_SECOND)
}

// DiffDaysFromTime 计算两个time.Time对应的北京时间相差的天数
//   - paramT1 时间1
//   - paramT2 时间2
//   - return 相差的天数 0表示同一天
func DiffDaysFromTime(paramT1 time.Time, paramT2 time.Time) int64 {
	return DiffDaysFromTimestamp(paramT1.UnixMilli(), paramT2.UnixMilli())
}

// IsSameDayFromTime 判断两个time.Time是否同一天
//   - paramT1 时间1
//   - paramT2 时间2
//   - return true表示同一天 false表示不同天
func IsSameDayFromTime(paramT1 time.Time, paramT2 time.Time) bool {
	return DiffDaysFromTime(paramT1, paramT2) == 0
}

// IsSameDayFromTimestamp 判断两个时间戳是否同一天
//   - paramT1 时间戳1 单位毫秒
//   - paramT2 时间戳2 单位毫秒
//   - return true表示同一天 false表示不同天
func IsSameDayFromTimestamp(paramT1 int64, paramT2 int64) bool {
	return DiffDaysFromTimestamp(paramT1, paramT2) == 0
}

// IsSameDayFromTimestampSecond 判断两个时间戳是否同一天
//   - paramT1 时间戳1 单位秒
//   - paramT2 时间戳2 单位秒
//   - return true表示同一天 false表示不同天
func IsSameDayFromTimestampSecond(paramT1 int64, paramT2 int64) bool {
	return DiffDaysFromTimestampSecond(paramT1, paramT2) == 0
}

// TimestampSecond2Time 时间戳转为time.Time 时间戳单位秒
//   - paramTimestamp 时间戳 单位秒
//   - return time.Time
func TimestampSecond2Time(paramTimestamp int64) time.Time {
	return time.Unix(paramTimestamp, 0)
}

// TimestampMillis2Time 时间戳转为time.Time 时间戳单位毫秒
//   - paramMillis 时间戳 单位毫秒
//   - return time.Time
func TimestampMillis2Time(paramMillis int64) time.Time {
	return time.UnixMilli(paramMillis)
}

// 参数的日期时间类
type ParamDateTime struct {
	dateType int       // 参数的日期类型 PARAM_TYPE_DATE(1)：日期(YYYY-MM-DD) ，PARAM_TYPE_DATETIME(2)：日期时间(YYYY-MM-DD hh:mm:ss)
	param    string    // 参数的日期字符串 原始的字符串
	value    time.Time // 参数的日期时间
}

// 参数的日期时间类型
func (p *ParamDateTime) GetDateType() int {
	return p.dateType
}

// 参数是不是空字符串
func (p *ParamDateTime) IsEmpty() bool {
	return p.param == ""
}

// 是否是日期
func (p *ParamDateTime) IsDate() bool {
	return p.dateType == PARAM_TYPE_DATE
}

// 是否是错误的
func (p *ParamDateTime) IsError() bool {
	return p.dateType == PARAM_TYPE_ERROR
}

// 是否是时间
func (p *ParamDateTime) IsDateTime() bool {
	return p.dateType == PARAM_TYPE_DATETIME
}

// 是否是错误 不是正确类型是，返回error,正确类型返回nil
func (p *ParamDateTime) Error() error {
	if p.dateType == PARAM_TYPE_ERROR {
		return errors.New("invalid datetime param format:" + p.param)
	}
	return nil
}

// 是否是是日期 不是日期返回error,正确类型返回nil
func (p *ParamDateTime) ErrorForDate() error {
	if p.dateType != PARAM_TYPE_DATE {
		return errors.New("invalid date format:" + p.param)
	}
	return nil
}

// 是否是是日期时间 不是日期时间返回error,正确类型返回nil
func (p *ParamDateTime) ErrorForDateTime() error {
	if p.dateType != PARAM_TYPE_DATETIME {
		return errors.New("invalid date time format:" + p.param)
	}
	return nil
}

// 获取参数的时间(要先判断是否是正确类型)
func (p *ParamDateTime) GetTime() time.Time {
	return p.value
}

// 获取参数的日期字符串（要先判断是否是正确类型）
func (p *ParamDateTime) GetDateString() string {
	return p.value.Format("2006-01-02")
}

// 获取参数的日期时间字符串（要先判断是否是正确类型）
func (p *ParamDateTime) GetDateTimeString() string {
	return p.value.Format("2006-01-02 15:04:05")
}

// 获取参数日期的当日最小值字符串（要先判断是否是正确类型）
func (p *ParamDateTime) MinDateTimeString() string {
	return p.value.Format("2006-01-02") + " 00:00:00"
}

// 获取参数日期的当日最大值字符串（要先判断是否是正确类型）
func (p *ParamDateTime) MaxDateTimeString() string {
	return p.value.Format("2006-01-02") + " 23:59:59"
}

// 获取YYYYMMDD格式的日期整数
func (p *ParamDateTime) DateNum() int {
	return p.value.Year()*10000 + int(p.value.Month())*100 + p.value.Day()
}

// 获取时间戳(单位秒)
func (p *ParamDateTime) GetSecond() int64 {
	return p.value.Unix()
}

// 获取时间戳(单位毫秒)
func (p *ParamDateTime) GetMillis() int64 {
	return p.value.UnixMilli()
}

// 新的日期时间参数（默认时区）
func NewParamDateTime(param string, local *time.Location) *ParamDateTime {
	p := &ParamDateTime{
		dateType: PARAM_TYPE_ERROR,
		param:    Trim(param),
	}
	if p.IsEmpty() {
		return p
	}
	if local == nil {
		local = time.Local
	}
	if IsDateFormat(param) {
		p.dateType = PARAM_TYPE_DATE
		v, err := time.ParseInLocation("2006-1-2", param, local)
		if err == nil {
			p.value = v
		} else {
			p.dateType = PARAM_TYPE_ERROR
		}
	} else if IsDateTimeFormat(param) {
		p.dateType = PARAM_TYPE_DATETIME
		v, err := time.ParseInLocation("2006-1-2 15:4:5", param, local)
		if err == nil {
			p.value = v
		} else {
			p.dateType = PARAM_TYPE_ERROR
		}
	}
	return p
}

// 新的北京时区日期时间参数
func NewBeijingParamDateTime(param string) *ParamDateTime {
	p := &ParamDateTime{
		dateType: PARAM_TYPE_ERROR,
		param:    Trim(param),
	}
	if p.IsEmpty() {
		return p
	}
	if IsDateFormat(param) {
		p.dateType = PARAM_TYPE_DATE
		v, err := time.ParseInLocation("2006-1-2", param, beijingLoc)
		if err == nil {
			p.value = v
		} else {
			p.dateType = PARAM_TYPE_ERROR
		}
	} else if IsDateTimeFormat(param) {
		p.dateType = PARAM_TYPE_DATETIME
		v, err := time.ParseInLocation("2006-1-2 15:4:5", param, beijingLoc)
		if err == nil {
			p.value = v
		} else {
			p.dateType = PARAM_TYPE_ERROR
		}
	}
	return p
}
