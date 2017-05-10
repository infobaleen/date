package date

import "time"

// Date represents a date with a resolution of one day.
type Date struct {
	day, month, year int
	loc              *time.Location
}

// ParseDate parses a formatted string and returns the value it represents.
// The layout defines the format by showing how the reference date
//	Mon Jan 2 2006
// woud be represented. See the documentation of time.ParseInLocation
// for more in depth documentation.
func ParseDate(layout, value string, loc *time.Location) (Date, error) {
	tim, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		return Date{}, err
	}
	return NewDateFromTime(tim), nil
}

// NewDate returns a Date in the specified location
// If the location is not specified UTC is assumed
func NewDate(year, month, day int, loc *time.Location) Date {
	if loc == nil {
		loc = time.UTC
	}
	return Date{day: day, month: month, year: year, loc: loc}
}

// NewDateFromTime returns the Date at the specified time
func NewDateFromTime(t time.Time) Date {
	year, month, day := t.Date()
	return NewDate(year, int(month), day, t.Location())
}

// Time returns the time at the date at the specified time of day
func (d Date) Time(hour, minute, second, nanosecond int) time.Time {
	return time.Date(d.year, time.Month(d.month), d.day, hour, minute, second, nanosecond, d.loc)
}

// Unix returns the date as a Unix time, the number of seconds elapsed since January 1, 1970 UTC
func (d Date) Unix() int64 {
	return d.Time(0, 0, 0, 0).Unix()
}

// In returns a new Date in the specified location. The day, month and year remain unchanged.
func (d Date) In(loc *time.Location) Date {
	d.loc = loc
	return d
}

// Location returns the location of the date.
func (d Date) Location() *time.Location {
	return d.loc
}

// TodayIn returns the current date in the specified timezone
func TodayIn(loc *time.Location) Date {
	return NewDateFromTime(time.Now().In(loc))
}

// PreviousWeekday returns the clostest previous date which is at the specified weekday.
func (d Date) PreviousWeekday(day time.Weekday) Date {
	t := d.Time(0, 0, 0, 0)
	return NewDateFromTime(t.AddDate(0, 0, -int((t.Weekday()-day+6)%7)-1))
}

// Add returns a modified date. It follows the same rules as t.AddDate().
func (d Date) Add(year, month, day int) Date {
	t := d.Time(0, 0, 0, 0)
	return NewDateFromTime(t.AddDate(year, month, day))
}

// Before returns true if the first moment of the receiver date occurs before the specified reference.
func (d Date) Before(ref Date) bool {
	return d.Time(0, 0, 0, 0).Before(ref.Time(0, 0, 0, 0))
}

// After returns true if the first moment of the receiver date occurs after the specified reference.
func (d Date) After(ref Date) bool {
	return d.Time(0, 0, 0, 0).After(ref.Time(0, 0, 0, 0))
}

// String returns a human readable string
func (d Date) String() string {
	return d.Time(0, 0, 0, 0).Format("2006-01-02Z07:00")
}

// MarshalJSON implements a the json.Marshaler interface
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte("\"" + d.String() + "\""), nil
}

// IsHoliday returns true if the date is a holiday (may be on a weekend)
func (d Date) IsHoliday() bool {
	_, ok := d.HolidayName()
	return ok
}

// HolidayName returns the name of the holiday and true if the date is a holiday (may be on a weekend)
func (d Date) HolidayName() (string, bool) {
	h, ok := HolidaysByLocation[d.loc.String()]
	if ok {
		return h.Find(d)
	}
	return "", false
}

// IsWorkday returns true if the date is neither on a weekend, nor a holiday.
func (d Date) IsWorkday() bool {
	return !(d.IsWeekday(time.Sunday) || d.IsWeekday(time.Saturday) || d.IsHoliday())
}

// Weekday returns the weekday of the date
func (d Date) Weekday() time.Weekday {
	return d.Time(0, 0, 0, 0).Weekday()
}

// ISOWeek returns the ISO 8601 year and week number in which the date occurs.
// Week ranges from 1 to 53. Jan 01 to Jan 03 of year n might belong to week 52 or 53 of year n-1, and Dec 29 to Dec 31 might belong to week 1 of year n+1.
func (d Date) ISOWeek() (year, week int) {
	return d.Time(0, 0, 0, 0).ISOWeek()
}

// IsWeekday returns true if the date is on the specified weekday
func (d Date) IsWeekday(day time.Weekday) bool {
	return d.Weekday() == day
}

// PreviousHoliday returns the closest previous holiday (may be on a weekend)
func (d Date) PreviousHoliday() Date {
	for {
		d = d.Add(0, 0, -1)
		if d.IsHoliday() {
			return d
		}
	}
}

// PreviousNonWorkday returns the closest previous day that is either a holiday or on a weekend
func (d Date) PreviousNonWorkday() Date {
	for {
		d = d.Add(0, 0, -1)
		if !d.IsWorkday() {
			return d
		}
	}
}

// PreviousWorkday returns the closest previous day that is neither on a weekend, nor a holiday.
func (d Date) PreviousWorkday() Date {
	for {
		d = d.Add(0, 0, -1)
		if d.IsWorkday() {
			return d
		}
	}
}
