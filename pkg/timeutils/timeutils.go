package timeutils

import (
	"time"
)

const pATTERN string = "2006-01-02 15:04:05"

/**
return Now time-object within Golang builtin type `time.Time`.
*/
func NowTime() time.Time {

	return time.Now()
}

/**
return formatted string from Now time-object within Golang builtin type `time.Time`.
Predefined layouts ANSIC, UnixDate, RFC3339 and others describe standard
and convenient representations of the reference time. For more information
about the formats and the definition of the reference time, see the
documentation for ANSIC and the other constants defined by package @time
*/
func NowTimeStringWithinPattern(pattern string) string {
	now := time.Now()

	resultStr := now.Format(pattern)

	return resultStr
}
func FormetTime(t time.Time, pattern string) string {

	return t.Format(pattern)
}

/**
return timestamp now int second
*/
func NowInSecond() int64 {
	timeNowInNanoSecond := time.Now().Unix()

	return timeNowInNanoSecond
}

/**
return timestamp now in millisecond
*/
func NowInMillisecond() int64 {
	timeNowInNanoSecond := time.Now().UnixNano()

	return timeNowInNanoSecond / 1e6
}

/**
return timestamp now in nanosecond
*/
func NowInNanosecond() int64 {
	timeNowInNanoSecond := time.Now().UnixNano()

	return timeNowInNanoSecond
}

/**
return the YEAR from system clock in int.
e.g, Now is 2019-08-10 10:29:56, then return 2019
*/
func NowYear() int {
	year := time.Now().Year()

	return year
}

/**
return the MONTH from system clock in int.
	1——"January",
	2——"February",
	3——"March",
	4——"April",
	5——"May",
	6——"June",
	7——"July",
	8——"August",
	9——"September",
	10——"October",
	11——"November",
	12——"December",
e.g, Now is 2019-08-10 10:29:56, then return 8
*/
func NowMonth() int {
	month := time.Now().Month()

	return int(month)
}

/**
return the MONTH from system clock in STRING.
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
e.g, Now is 2019-08-10 10:29:56, then return "August"
*/
func NowMonthString() string {
	month := time.Now().Month()

	return month.String()
}

/**
return the Day from system clock.
e.g, Now is 2019-08-10 10:29:56, then return 10
*/
func NowDay() int {
	return time.Now().Day()
}

/**
return the Hour from system clock.
e.g, Now is 2019-01-10 10:29:56, then return 10
*/
func NowHour() int {
	return time.Now().Hour()
}

/**
return the Minute from system clock.
e.g, Now is 2019-01-10 10:29:56, then return 29
*/
func NowMinute() int {
	return time.Now().Minute()
}

/**
return the Second from system clock.
e.g, Now is 2019-01-10 10:29:56, then return 56
*/
func NowSecond() int {
	return time.Now().Second()
}
