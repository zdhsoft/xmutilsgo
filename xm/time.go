package xm

//仅演示获取当天指定时间的时间戳
import (
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
)

var beijingLoc *time.Location

func init() {
	beijingLoc, _ = time.LoadLocation("Asia/Shanghai")
}

// TDateTime 日期时间
type TDateTime struct {
	// 时间戳 单位毫秒
	dt_timestamp int64
	// 类型
	dt_type int8
}

// IsUTC 是否是格林威治时间
func (t *TDateTime) IsUTC() bool {
	return t.dt_type == DT_TYPE_UTC
}

// IsBeijing 是否是北京时间
func (t *TDateTime) IsBeijing() bool {
	return t.dt_type == DT_TYPE_BEIJING
}

// GetBeijingTimezone 取北京时区
func GetBeijingTimezone() int {
	return TIMEZONE_BEIJING
}

// ToBeijing 将时间变成北京时间戳
func (t *TDateTime) ToBeijing() *TDateTime {
	st := TDateTime{dt_timestamp: t.dt_timestamp, dt_type: t.dt_type}
	if st.IsUTC() {
		st.dt_timestamp -= MILLIS_BY_TIMEZONE_BEIJING
		st.dt_type = DT_TYPE_BEIJING
	}
	return &st
}

// SelfToBeijing 将自己变成北京时间戳
func (t *TDateTime) SelfToBeijing() *TDateTime {
	if t.IsUTC() {
		t.dt_timestamp -= MILLIS_BY_TIMEZONE_BEIJING
		t.dt_type = DT_TYPE_UTC
	}
	return t
}

// ToUTC 变成UTC时间戳
func (t *TDateTime) ToUTC() *TDateTime {
	st := TDateTime{dt_timestamp: t.dt_timestamp, dt_type: t.dt_type}
	if st.IsBeijing() {
		st.dt_timestamp += MILLIS_BY_TIMEZONE_BEIJING
		st.dt_type = DT_TYPE_UTC
	}
	return &st
}

// SelfToUTC 将自己变成UTC时间戳
func (t *TDateTime) SelfToUTC() *TDateTime {
	if t.IsBeijing() {
		t.dt_timestamp += MILLIS_BY_TIMEZONE_BEIJING
		t.dt_type = DT_TYPE_UTC
	}
	return t
}

// ToBeijingZeroTime 取对应北京0点的utc时间戳
func (t *TDateTime) ToBeijingZeroTime() *TDateTime {
	st := t.ToBeijing()
	st.dt_timestamp -= st.dt_timestamp % MILLIS_BY_DAY
	return st.ToUTC()
}

// GetMillis 取当前时间的时间戳
func (t *TDateTime) GetMillis() int64 {
	return t.dt_timestamp
}

// GetSecond 取当前时间的秒数
func (t *TDateTime) GetSecond() int64 {
	return (t.dt_timestamp - t.dt_timestamp%MILLIS_BY_SECOND) / MILLIS_BY_SECOND
}

// GetType 取当前时间的类型
func (t *TDateTime) GetType() int8 {
	return t.dt_type
}

// MakeDateTime 生成当前时间的时间对象
func MakeDateTime() *TDateTime {
	st := TDateTime{dt_timestamp: time.Now().UnixMilli(), dt_type: DT_TYPE_UTC}
	return &st
}

// MakeBeijingDateTime 生成北京时间的时间对象
func MakeBeijingDateTime() *TDateTime {
	st := TDateTime{dt_timestamp: time.Now().UnixMilli(), dt_type: DT_TYPE_UTC}
	return st.SelfToBeijing()
}

func DateTimeFromMillis(paramMillis int64) *TDateTime {
	st := TDateTime{dt_timestamp: paramMillis, dt_type: DT_TYPE_UTC}
	return &st
}

// MakeFromBeijingDate 取指定日期的北京时间 单位毫秒
func MakeFromBeijingDate(paramDate string) (*TDateTime, error) {
	stNow, err := ParseDateTimeForBeijingMillis(paramDate)
	if err == nil {
		st := TDateTime{dt_timestamp: stNow, dt_type: DT_TYPE_UTC}
		return &st, nil
	} else {
		return nil, err
	}
}

// GetNowMillis 取当前时间戳(毫秒)
func GetNowMillis() int64 {
	return time.Now().UnixMilli()
}

// GetNowSecond 取当前时间戳(秒）
func GetNowSecond() int64 {
	return time.Now().Unix()
}

// ParseDateTimeForBeijingMillis 解析北京格式的日期 单位毫秒
func ParseDateTimeForBeijingMillis(paramDate string) (int64, error) {
	//获取北京时区
	//时区定义参考： https://jp.cybozu.help/general/zh/admin/list_systemadmin/list_localization/timezone.html
	loc, _ := time.LoadLocation("Asia/Shanghai")
	startTime, err := time.ParseInLocation("2006-1-2", paramDate, loc)
	if err != nil {
		return 0, err
	}
	return startTime.UnixMilli(), nil
}

// ParseDateTimeForBeijingSecond 解析北京格式的日期 单位秒
func ParseDateTimeForBeijingSecond(paramDate string) (int64, error) {
	stNow, err := ParseDateTimeForBeijingMillis(paramDate)
	stNow = (stNow - stNow%MILLIS_BY_SECOND) / MILLIS_BY_SECOND
	return stNow, err
}

// Timestamp2BeijingDate 将时间戳转为北京时区的日期 时间戳单位秒
func Timestamp2BeijingDate(paramTimestamp int64) string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.Unix(paramTimestamp, 0).In(loc).Format("2006-01-02")
}

// BeijingFormat 将时间转换为北京时区的格式字符串
func BeijingFormat(paramDateTime time.Time, paramFormat string) string {
	return paramDateTime.In(beijingLoc).Format(paramFormat)
}

// BeijingDateString 将时间转换为北京时区的日期字符串
func BeijingDateString(paramDateTime time.Time) string {
	return BeijingFormat(paramDateTime, "2006-01-02")
}

// BeijingDateTimeString 将时间转换为北京时区的日期时间字符串
func BeijingDateTimeString(paramDateTime time.Time) string {
	return BeijingFormat(paramDateTime, "2006-01-02 15:04:05")
}

// BeijingTimeString 将时间转换为北京时区的时间字符串
func BeijingTimeString(paramDateTime time.Time) string {
	return BeijingFormat(paramDateTime, "15:04:05")
}

// BeijingCompactDateString 将时间转换为北京时区的日期字符串(压缩版)
func BeijingCompactDateString(paramDateTime time.Time) string {
	return BeijingFormat(paramDateTime, "20060102")
}

// BeijingCompactDateTimeString 将时间转换为北京时区的日期时间字符串(压缩版)
func BeijingCompactDateTimeString(paramDateTime time.Time) string {
	return BeijingFormat(paramDateTime, "20060102150405")
}

// BeijingCompactTimeString 将时间转换为北京时区的时间字符串(压缩版)
func BeijingCompactTimeString(paramDateTime time.Time) string {
	return BeijingFormat(paramDateTime, "150405")
}

// GetTimeOperationDayString 获取指定时间根据天数计算后得到的时间
func GetTimeOperationDayString(paramDateTime time.Time, paramDays int) string {
	paramDateTime = paramDateTime.AddDate(0, 0, paramDays)
	return BeijingFormat(paramDateTime, "20060102")
}

// DiffDaysFromTimestamp 计算两个时间戳对应的北京时间相差的天数，单位毫秒
func DiffDaysFromTimestamp(paramT1 int64, paramT2 int64) int64 {
	dt1 := DateTimeFromMillis(paramT1).ToBeijingZeroTime()
	dt2 := DateTimeFromMillis(paramT2).ToBeijingZeroTime()
	return (dt1.GetMillis() - dt2.GetMillis()) / MILLIS_BY_DAY
}

// DiffDaysFromTimestampSecond 计算两个时间戳对应的北京时间相差的天数，单位秒
func DiffDaysFromTimestampSecond(paramT1 int64, paramT2 int64) int64 {
	return DiffDaysFromTimestamp(paramT1*MILLIS_BY_SECOND, paramT2*MILLIS_BY_SECOND)
}

// DiffDaysFromTime 计算两个time.Time对应的北京时间相差的天数
func DiffDaysFromTime(paramT1 time.Time, paramT2 time.Time) int64 {
	return DiffDaysFromTimestamp(paramT1.UnixMilli(), paramT2.UnixMilli())
}

func IsSameDayFromTime(paramT1 time.Time, paramT2 time.Time) bool {
	return DiffDaysFromTime(paramT1, paramT2) == 0
}

func IsSameDayFromTimestamp(paramT1 int64, paramT2 int64) bool {
	return DiffDaysFromTimestamp(paramT1, paramT2) == 0
}

func IsSameDayFromTimestampSecond(paramT1 int64, paramT2 int64) bool {
	return DiffDaysFromTimestampSecond(paramT1, paramT2) == 0
}

// TimestampSecond2Time 时间戳转为time.Time 时间戳单位秒
func TimestampSecond2Time(paramTimestamp int64) time.Time {
	return time.Unix(paramTimestamp, 0)

}

// TimestampMillis2Time 时间戳转为time.Time 时间戳单位毫秒
func TimestampMillis2Time(paramMillis int64) time.Time {
	return time.UnixMilli(paramMillis / MILLIS_BY_SECOND)
}
