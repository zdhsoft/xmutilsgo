package xm

import (
	"testing"
	"time"
)

func Test_DateNum(t *testing.T) {
	timestamp := 1729206000
	tm := time.Unix(int64(timestamp), 0)
	dateNumUnix := FormatDateNum(tm)
	dateNumBJ := FormatDateNumForBeijing(tm)
	DestStr := "2024-10-18"
	DestNum := 20241018

	if DateNum2DateStr(dateNumUnix, "-") != DestStr {
		t.Errorf("DateNum2DateStr %s != %s", DateNum2DateStr(dateNumUnix, "-"), DestStr)
	}

	if dateNumBJ != DestNum {
		t.Errorf("dateNumBJ %d != %d", dateNumBJ, DestNum)
	}
	// t.Log("==> unix", dateNumUnix, "=>", DateNum2DateStr(dateNumUnix, "-"))
	// t.Log("==> beijing", dateNumBJ, "=>", DateNum2DateStr(20240101, "-"))
	//  ArraySort<[]int>(lst)
}
