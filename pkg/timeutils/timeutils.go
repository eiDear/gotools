/**
* @Author       :jerry
* @Version      :1.0.0
* @Date         :下午7:12 19-8-12
* @Description  : timeutils to help you
				(1) get time in specified type,
				(2) and format time to specified pattern,
				(3) parse time with specified pattern.
*/
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
documentation for ANSIC and the other constants defined by package  time
*/
func NowTimeStringWithinPattern(pattern string) string {
	now := time.Now()

	resultStr := now.Format(pattern)

	return resultStr
}

/**
return formatted string from specified parameter @param t using pattern.
*/
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

/**
 ParseTime parses a formatted string and returns the time value it represents.
 The layout defines the format by showing how the reference time,
 defined to be
	Mon Jan 2 15:04:05 -0700 MST 2006
 would be interpreted if it were the value; it serves as an example of
 the input format. The same interpretation will then be made to the
 input string.

 Predefined layouts ANSIC, UnixDate, RFC3339 and others describe standard
 and convenient representations of the reference time. For more information
 about the formats and the definition of the reference time, see the
 documentation for ANSIC and the other constants defined by this package.
 Also, the executable example for Time.Format demonstrates the working
 of the layout string in detail and is a good reference.

 Elements omitted from the value are assumed to be zero or, when
 zero is impossible, one, so parsing "3:04pm" returns the time
 corresponding to Jan 1, year 0, 15:04:00 UTC (note that because the year is
 0, this time is before the zero Time).
 Years must be in the range 0000..9999. The day of the week is checked
 for syntax but it is otherwise ignored.

 In the absence of a time zone indicator, Parse returns a time in UTC.

 When parsing a time with a zone offset like -0700, if the offset corresponds
 to a time zone used by the current location (Local), then Parse uses that
 location and zone in the returned time. Otherwise it records the time as
 being in a fabricated location with time fixed at the given zone offset.

 When parsing a time with a zone abbreviation like MST, if the zone abbreviation
 has a defined offset in the current location, then that offset is used.
 The zone abbreviation "UTC" is recognized as UTC regardless of location.
 If the zone abbreviation is unknown, Parse records the time as being
 in a fabricated location with the given zone abbreviation and a zero offset.
 This choice means that such a time can be parsed and reformatted with the
 same layout losslessly, but the exact instant used in the representation will
 differ by the actual zone offset. To avoid such problems, prefer time layouts
 that use a numeric zone offset, or use ParseInLocation.
*/
func ParseTime(pattern string, value string) (time.Time, error) {
	t, err := time.Parse(pattern, value)
	return t, err
}

/**
ParseTimeInLocation is like Parse but differs in two important ways.
First, in the absence of time zone information, Parse interprets a time as UTC;
ParseInLocation interprets the time as in the given location.
Second, when given a zone offset or abbreviation, Parse tries to match it
against the Local location; ParseInLocation uses the given location.
*/
func ParseTimeInLocation(pattern, value string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(pattern, value, loc)
}
