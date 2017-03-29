package date

import "time"

// Date represents a date with a resolution of one day.
type Date struct {
	day, month, year int
	loc              *time.Location
}

// NewDate returns a Date in the specified location
// If the location is not specified UTC is assumed
func NewDate(year, month, day int, loc *time.Location) Date {
	if loc == nil {
		loc = time.UTC
	}
	return Date{day: day, month: month, year: year, loc: loc}
}

// DateFromTime returns the Date at the specified time
func DateFromTime(t time.Time) Date {
	year, month, day := t.Date()
	return Date{year: year, month: int(month), day: day, loc: t.Location()}
}

// Time returns the time at the date at the specified time of day
func (d Date) Time(hour, minute, second, nanosecond int) time.Time {
	return time.Date(d.year, time.Month(d.month), d.day, hour, minute, second, nanosecond, d.loc)
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
	return DateFromTime(time.Now().In(loc))
}

// PreviousWeekday returns the clostest previous date which is at the specified weekday.
func (d Date) PreviousWeekday(day time.Weekday) Date {
	t := d.Time(0, 0, 0, 0)
	return DateFromTime(t.AddDate(0, 0, -int((t.Weekday()-day+6)%7)-1))
}

// Add returns a modified date. It follows the same rules as t.AddDate().
func (d Date) Add(year, month, day int) Date {
	t := d.Time(0, 0, 0, 0)
	return DateFromTime(t.AddDate(year, month, day))
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
	return d.Time(0, 0, 0, 0).Format("Mon Jan 2 2006 -0700 MST")
}

// IsHoliday returns true if the date is a holiday (may be on a weekend)
func (d Date) IsHoliday() bool {
	return d.isHoliday()
}

// IsWorkday returns true if the date is neither on a weekend, nor a holiday.
func (d Date) IsWorkday() bool {
	return !(d.IsWeekday(time.Sunday) || d.IsWeekday(time.Saturday) || d.IsHoliday())
}

// IsWeekday returns true if the date is on the specified weekday
func (d Date) IsWeekday(day time.Weekday) bool {
	return d.Time(0, 0, 0, 0).Weekday() == day
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
