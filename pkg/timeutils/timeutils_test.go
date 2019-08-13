package timeutils

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestNowInSecond(t *testing.T) {
	timeNowInSeconds := time.Now().Unix()
	ts := NowInSecond()
	if math.Abs(float64(ts-int64(timeNowInSeconds))) > 1 {
		t.Errorf("timestamp now in second are not equavent. timeNowInSeconds=%d, ts=%d", timeNowInSeconds, ts)
	}
}

func TestNowMonth(t *testing.T) {
	month := NowMonth()
	fmt.Printf("month: %d", month)

}

func TestParseTime(t *testing.T) {
	const pATTERN_FULL string = "2006-01-02 15:04:05"
	timeValue := "2019-08-12 19:57:23"
	ts, _ := time.Parse(pATTERN_FULL, timeValue)
	result, _ := ParseTime(pATTERN_FULL, "2019-08-12 19:57:23")
	fmt.Printf("result: %s", result)
	if ts != result {
		t.Errorf("ParseTime gives the wrong result: %s, expected: %s", result, ts)
	}
}

func TestParseTimeInLocation(t *testing.T) {
	const pATTERN_FULL string = "2006-01-02 15:04:05"
	timeValue := "2019-08-12 19:57:23"
	ts, _ := time.ParseInLocation(pATTERN_FULL, timeValue, time.Local)
	result, _ := ParseTimeInLocation(pATTERN_FULL, "2019-08-12 19:57:23", time.Local)
	fmt.Printf("result: %s", result)
	if ts != result {
		t.Errorf("ParseTime gives the wrong result: %s, expected: %s", result, ts)
	}
}
func TestLoadLocation(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Shanghai")
	value := "2019-08-12 19:57:23"
	result, _ := ParseTimeInLocation(pATTERN, value, location)
	fmt.Printf("result: %s", result)
	srcTime := time.Date(2019, 8, 12, 19, 57, 23, 0, location)
	if srcTime != result {
		t.Errorf("ParseTime gives the wrong result: %s, expected: %s", result, srcTime)
	}
}
