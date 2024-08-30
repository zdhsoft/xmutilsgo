package xm

import (
	"testing"
	"time"
)

func Test_ParseDateTimeForBeijingTimestamp(t *testing.T) {
	bjDateTime := "2024-08-30"
	dt, err := ParseDateTimeForBeijingMillis(bjDateTime)
	if err != nil {
		t.Error(err)
	}
	dest := int64(1724947200000)
	destSecond := dest / 1000
	if dt != dest {
		t.Errorf("北京时间：%s 的时间戳不是%d (毫秒)", bjDateTime, dest)
	}

	dtSecond, errSecond := ParseDateTimeForBeijingSecond(bjDateTime)
	if errSecond != nil {
		t.Error(errSecond)
	}

	if dtSecond != destSecond {
		t.Errorf("北京时间：%s 的时间戳不是%d (秒)", bjDateTime, destSecond)
	}
}

func Test_Timestamp2Beijing(t *testing.T) {
	destDateTime := "2024-08-30 18:18:18"
	destDate := "2024-08-30"
	destTime := "18:18:18"
	destCompactDateTime := "20240830181818"
	destCompactDate := "20240830"
	destCompactTime := "181818"

	dtTimestamp := int64(1725013098)
	bjDate := Timestamp2BeijingDate(dtTimestamp)
	bjTime := Timestamp2BeijingTime(dtTimestamp)
	bjDateTime := Timestamp2BeijingDateTime(dtTimestamp)
	t.Log(bjDate)
	t.Log(bjTime)
	t.Log(bjDateTime)
	if bjDate != destDate {
		t.Errorf("(%s)%d => Date != %s", destDateTime, dtTimestamp, bjDate)
	}
	if bjTime != destTime {
		t.Errorf("(%s)%d => Time != %s", destDateTime, dtTimestamp, bjTime)
	}

	if bjDateTime != destDateTime {
		t.Errorf("(%s)%d => DateTime != %s", destDateTime, dtTimestamp, bjDateTime)
	}

	stTime := time.Unix(dtTimestamp, 0)

	bjDate = BeijingDateString(stTime)
	bjTime = BeijingTimeString(stTime)
	bjDateTime = BeijingDateTimeString(stTime)
	t.Log(bjDate)
	t.Log(bjTime)
	t.Log(bjDateTime)
	if bjDate != destDate {
		t.Errorf("(%s)%d => Date != %s", destDateTime, dtTimestamp, bjDate)
	}
	if bjTime != destTime {
		t.Errorf("(%s)%d => Time != %s", destDateTime, dtTimestamp, bjTime)
	}

	if bjDateTime != destDateTime {
		t.Errorf("(%s)%d => DateTime != %s", destDateTime, dtTimestamp, bjDateTime)
	}
	bjDate = BeijingCompactDateString(stTime)
	bjTime = BeijingCompactTimeString(stTime)
	bjDateTime = BeijingCompactDateTimeString(stTime)
	t.Log(bjDate)
	t.Log(bjTime)
	t.Log(bjDateTime)
	if bjDate != destCompactDate {
		t.Errorf("(%s)%d => Compact Date != %s", destDateTime, dtTimestamp, bjDate)
	}
	if bjTime != destCompactTime {
		t.Errorf("(%s)%d => Compact Time != %s", destDateTime, dtTimestamp, bjTime)
	}

	if bjDateTime != destCompactDateTime {
		t.Errorf("(%s)%d => Compact DateTime != %s", destDateTime, dtTimestamp, bjDateTime)
	}
}
