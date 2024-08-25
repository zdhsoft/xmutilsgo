package xm

import "time"

// TDateTime 日期时间
type TDateTime struct {
	// 时间戳 单位毫秒
	dtTimestamp int64 `json:"timestamp"`
	// 类型
	dtType int8 `json:"type"`
}

// IsUTC 是否是格林威治时间
func (t *TDateTime) IsUTC() bool {
	return t.dtType == DT_TYPE_UTC
}

// IsBeijing 是否是北京时间
func (t *TDateTime) IsBeijing() bool {
	return t.dtType == DT_TYPE_BEIJING
}

// GetBeijingTimezone 取北京时区
func GetBeijingTimezone() int {
	return TIMEZONE_BEIJING
}

// ToBeijing 将时间变成北京时间戳
func (t *TDateTime) ToBeijing() *TDateTime {
	st := TDateTime{dtTimestamp: t.dtTimestamp, dtType: t.dtType}
	if st.IsUTC() {
		st.dtTimestamp -= MILLIS_BY_TIMEZONE_BEIJING
		st.dtType = DT_TYPE_BEIJING
	}
	return &st
}

// SelfToBeijing 将自己变成北京时间戳
func (t *TDateTime) SelfToBeijing() *TDateTime {
	if t.IsUTC() {
		t.dtTimestamp -= MILLIS_BY_TIMEZONE_BEIJING
		t.dtType = DT_TYPE_UTC
	}
	return t
}

// ToUTC 变成UTC时间戳
func (t *TDateTime) ToUTC() *TDateTime {
	st := TDateTime{dtTimestamp: t.dtTimestamp, dtType: t.dtType}
	if st.IsBeijing() {
		st.dtTimestamp += MILLIS_BY_TIMEZONE_BEIJING
		st.dtType = DT_TYPE_UTC
	}
	return &st
}

// SelfToUTC 将自己变成UTC时间戳
func (t *TDateTime) SelfToUTC() *TDateTime {
	if t.IsBeijing() {
		t.dtTimestamp += MILLIS_BY_TIMEZONE_BEIJING
		t.dtType = DT_TYPE_UTC
	}
	return t
}

// ToBeijingZeroTime 取对应北京0点的utc时间戳
func (t *TDateTime) ToBeijingZeroTime() *TDateTime {
	st := t.ToBeijing()
	st.dtTimestamp -= st.dtTimestamp % MILLIS_BY_DAY
	return st.ToUTC()
}

// GetMillis 取当前时间的时间戳
func (t *TDateTime) GetMillis() int64 {
	return t.dtTimestamp
}

// GetSecond 取当前时间的秒数
func (t *TDateTime) GetSecond() int64 {
	return (t.dtTimestamp - t.dtTimestamp%MILLIS_BY_SECOND) / MILLIS_BY_SECOND
}

// GetType 取当前时间的类型
func (t *TDateTime) GetType() int8 {
	return t.dtType
}

// MakeDateTime 生成当前时间的时间对象
func MakeDateTime() *TDateTime {
	st := TDateTime{dtTimestamp: time.Now().UnixMilli(), dtType: DT_TYPE_UTC}
	return &st
}

// MakeBeijingDateTime 生成北京时间的时间对象
func MakeBeijingDateTime() *TDateTime {
	st := TDateTime{dtTimestamp: time.Now().UnixMilli(), dtType: DT_TYPE_UTC}
	return st.SelfToBeijing()
}

func DateTimeFromMillis(paramMillis int64) *TDateTime {
	st := TDateTime{dtTimestamp: paramMillis, dtType: DT_TYPE_UTC}
	return &st
}

// MakeFromBeijingDate 取指定日期的北京时间 单位毫秒
func MakeFromBeijingDate(paramDate string) (*TDateTime, error) {
	stNow, err := ParseDateTimeForBeijingMillis(paramDate)
	if err == nil {
		st := TDateTime{dtTimestamp: stNow, dtType: DT_TYPE_UTC}
		return &st, nil
	} else {
		return nil, err
	}
}
