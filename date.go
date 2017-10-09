package date

import (
	"time"
)

// Date represents a date with a resolution of one day.
type Date int32

// Limits
const (
	Min = -1 << 31
	Max = 1<<31 - 1
)

// First returns the earliest provided date
func First(d Date, ds ...Date) Date {
	for _, de := range ds {
		if d > de {
			d = de
		}
	}
	return d
}

// Last returns the latest provided date
func Last(d Date, ds ...Date) Date {
	for _, de := range ds {
		if d < de {
			d = de
		}
	}
	return d
}

// ParseDate parses a formatted string and returns the value it represents.
// The layout defines the format by showing how the reference date
//	Mon Jan 2 2006
// woud be represented. See the documentation of time.ParseInLocation
// for more in depth documentation.
func ParseDate(layout, value string) (Date, error) {
	tim, err := time.ParseInLocation(layout, value, time.UTC)
	if err != nil {
		return 0, err
	}
	return NewDate(tim.Date()), nil
}

// NewDate returns a Date in the specified location
// If the location is not specified UTC is assumed
func NewDate(year int, month time.Month, day int) Date {
	seconds := time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Unix() - time.Time{}.Unix()
	return Date((seconds) / (24 * 60 * 60))
}

// Today returns the current date in your local timezone (see time.Now() for details).
func Today() Date {
	return NewDate(time.Now().Date())
}

// Time returns the time at the date at the specified time of day
func (d Date) Time(hour, minute, second, nanosecond int, loc *time.Location) time.Time {
	if loc == nil {
		loc = time.UTC
	}
	return time.Time{}.AddDate(0, 0, int(d))
}

// PreviousWeekday returns the clostest previous date which is at the specified weekday.
func (d Date) PreviousWeekday(day time.Weekday) Date {
	return d - Date((d.Weekday()-day+6)%7) - 1
}

// Add returns a modified date. It follows the same rules as t.AddDate().
func (d Date) Add(year, month, day int) Date {
	return NewDate(d.Time(0, 0, 0, 0, nil).AddDate(year, month, day).Date())
}

// String returns a human readable string
func (d Date) String() string {
	return d.Time(0, 0, 0, 0, nil).Format("2006-01-02")
}

// Weekday returns the weekday of the date
func (d Date) Weekday() time.Weekday {
	// Sunday is 0 and Day 0 is a Monday, so we would like to do (d+1)%7).
	// Unfortunately % is calculated towards 0, instead of negative infinity, so we use ((d+1)%7+7)%7,
	// which can be optimized to:
	return time.Weekday((d%7 + 8) % 7)
}

// ISOWeek returns the ISO 8601 year and week number in which the date occurs.
// Week ranges from 1 to 53. Jan 01 to Jan 03 of year n might belong to week 52 or 53 of year n-1, and Dec 29 to Dec 31 might belong to week 1 of year n+1.
func (d Date) ISOWeek() (year, week int) {
	return time.Time{}.AddDate(0, 0, int(d)).ISOWeek()
}

// Year returns the year
func (d Date) Year() int {
	return time.Time{}.AddDate(0, 0, int(d)).Year()
}

// Month returns the month of the year [1,12]
func (d Date) Month() time.Month {
	return time.Time{}.AddDate(0, 0, int(d)).Month()
}

// Day returns the day of the month [1:31]
func (d Date) Day() int {
	return time.Time{}.AddDate(0, 0, int(d)).Day()
}

// Format returns a textual representation of the time value formatted
// according to layout, which defines the format by showing how the reference
// time, defined to be
//	Mon Jan 2 15:04:05 -0700 MST 2006
// would be displayed if it were the value; it serves as an example of the
// desired output. See time.Time.Format for details.
func (d Date) Format(layout string) string {
	return d.Time(0, 0, 0, 0, nil).Format(layout)
}
