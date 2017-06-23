package date

import "time"

// Date represents a date with a resolution of one day.
type Date int32

// ParseDate parses a formatted string and returns the value it represents.
// The layout defines the format by showing how the reference date
//	Mon Jan 2 2006
// woud be represented. See the documentation of time.ParseInLocation
// for more in depth documentation.
func ParseDate(layout, value string) (Date, error) {
	tim, err := time.ParseInLocation(layout, value, nil)
	if err != nil {
		return 0, err
	}
	return NewDateFromTime(tim), nil
}

// NewDate returns a Date in the specified location
// If the location is not specified UTC is assumed
func NewDate(year int, month time.Month, day int) Date {
	now := time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Unix()
	if now%(24*60*60) != 0 {
		panic("this should never happen")
	}
	zero := time.Time{}.Unix()
	return Date((now - zero) / (24 * 60 * 60))
}

// NewDateFromTime returns the Date at the specified time
func NewDateFromTime(t time.Time) Date {
	year, month, day := t.Date()
	return NewDate(year, month, day)
}

// Time returns the time at the date at the specified time of day
func (d Date) Time(hour, minute, second, nanosecond int, loc *time.Location) time.Time {
	return time.Date(0, 0, int(d), hour, minute, second, nanosecond, loc)
}

// Unix returns the date as a Unix time, the number of seconds elapsed since January 1, 1970 UTC
func (d Date) Unix(loc *time.Location) int64 {
	return d.Time(0, 0, 0, 0, loc).Unix()
}

// TodayIn returns the current date in the specified timezone
func TodayIn(loc *time.Location) Date {
	return NewDateFromTime(time.Now().In(loc))
}

// PreviousWeekday returns the clostest previous date which is at the specified weekday.
func (d Date) PreviousWeekday(day time.Weekday) Date {
	return d - Date((d.Weekday()-day+6)%7) - 1
}

// Add returns a modified date. It follows the same rules as t.AddDate().
func (d Date) Add(year, month, day int) Date {
	t := d.Time(0, 0, 0, 0, nil)
	return NewDateFromTime(t.AddDate(year, month, day))
}

// Before returns true if the first moment of the receiver date occurs before the specified reference.
func (d Date) Before(ref Date) bool {
	return d < ref
}

// After returns true if the first moment of the receiver date occurs after the specified reference.
func (d Date) After(ref Date) bool {
	return d > ref
}

// String returns a human readable string
func (d Date) String() string {
	return d.Time(0, 0, 0, 0, nil).Format("2006-01-02")
}

// MarshalJSON implements a the json.Marshaler interface
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte("\"" + d.String() + "\""), nil
}

// Weekday returns the weekday of the date
func (d Date) Weekday() time.Weekday {
	return d.Time(0, 0, 0, 0, nil).Weekday() // TODO: We can do that ourselves.
}

// ISOWeek returns the ISO 8601 year and week number in which the date occurs.
// Week ranges from 1 to 53. Jan 01 to Jan 03 of year n might belong to week 52 or 53 of year n-1, and Dec 29 to Dec 31 might belong to week 1 of year n+1.
func (d Date) ISOWeek() (year, week int) {
	return d.Time(0, 0, 0, 0, nil).ISOWeek() // TODO: We can do that ourselves.
}

// Year returns the year
func (d Date) Year() int {
	return d.Time(0, 0, 0, 0, nil).Year() // TODO: We can do that ourselves.
}

// Month returns the month of the year [1,12]
func (d Date) Month() time.Month {
	return d.Time(0, 0, 0, 0, nil).Month() // TODO: We can do that ourselves.
}

// Day returns the day of the month [1:31]
func (d Date) Day() int {
	return d.Time(0, 0, 0, 0, nil).Day() // TODO: We can do that ourselves.
}

// IsWeekday returns true if the date is on the specified weekday
func (d Date) IsWeekday(day time.Weekday) bool {
	return d.Weekday() == day
}

// Sub returns d-e, the number of days that elapsed since e until d.
// If d occured before e, the result is negative.
func (d Date) Sub(e Date) int32 {
	return int32(d - e)
}
