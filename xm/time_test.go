package xm

import (
	"testing"
	"time"
)

func Test_ParseDateTimeForBeijingTimestamp(t *testing.T) {
	bjDateTime := "2024-08-30"
	dt, err := ParseDateForBeijingMillis(bjDateTime)
	if err != nil {
		t.Error(err)
	}
	dest := int64(1724947200000)
	destSecond := dest / 1000
	if dt != dest {
		t.Errorf("北京时间：%s 的时间戳不是%d (毫秒)", bjDateTime, dest)
	}

	dtSecond, errSecond := ParseDateForBeijingSecond(bjDateTime)
	if errSecond != nil {
		t.Error(errSecond)
	}

	if dtSecond != destSecond {
		t.Errorf("北京时间：%s 的时间戳不是%d (秒)", bjDateTime, destSecond)
	}
}

func Test_Days(t *testing.T) {
	t1 := "2024-08-30 18:18:18"
	t2 := "2024-08-30 09:28:38"
	t3 := "2024-09-30 18:18:18"

	tt1, _ := ParseDateTimeForBeijingSecond(t1)
	tt2, _ := ParseDateTimeForBeijingSecond(t2)
	tt3, _ := ParseDateTimeForBeijingSecond(t3)

	if !IsSameDayFromTimestampSecond(tt1, tt2) {
		t.Errorf("%s 和 %s 因该是同一天！", t1, t2)
	}

	days := DiffDaysFromTimestampSecond(tt1, tt3)
	days2 := DiffDaysFromTimestampSecond(tt2, tt3)

	if days < 0 {
		days = -days
	}
	if days2 < 0 {
		days2 = -days2
	}

	t.Logf("计算出来%s 和 %s 相差%d天", t1, t3, days)
	t.Logf("计算出来%s 和 %s 相差%d天", t2, t3, days2)

	if days != 31 {
		t.Errorf("计算出来%s 和 %s = %d 相差不是31天", t1, t3, days)
	}
	if days2 != 31 {
		t.Errorf("计算出来%s 和 %s = %d 相差不是31天", t2, t3, days2)
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
