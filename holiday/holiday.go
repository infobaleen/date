package holiday

import (
	"sort"
	"time"

	"github.com/infobaleen/date"
)

// Calendar represents a collection of holidays
type Calendar struct {
	Changing []ChangingHoliday
	Fixed    []FixedHoliday
}

// Match returns the names of all holidays matching the specified date, sorted in increasing order.
func (cal *Calendar) Match(d date.Date) (matches []string) {
	for _, f := range cal.Fixed {
		if f.Match(d) {
			matches = append(matches, f.Name)
		}
	}
	for _, c := range cal.Changing {
		for _, name := range c.Match(d) {
			matches = append(matches, name)
		}
	}
	sort.Strings(matches)
	return
}

// ChangingHoliday encodes arbitrarily defined holidays.
type ChangingHoliday func(date.Date) []string

// FixedHoliday encodes holidays that repeat annually on the same date.
type FixedHoliday struct {
	// Name of holiday
	Name string
	// Annual date
	Day   int
	Month time.Month
	// Range of years specified by first year with and first year without holiday.
	// If FirstYear == EndYear: Range extends from negative infinity to positive infinity
	// If FirstYear > EndYear: Range extends from FirstYear to positive infinity
	FirstYear, EndYear int
}

// Match returns true if the holiday occurs on the specified date.
// In non-leap years February 29th matches March 1st.
func (f *FixedHoliday) Match(d date.Date) bool {
	var year = d.Year()
	return (f.FirstYear == f.EndYear || (f.FirstYear <= year && (year < f.EndYear || f.EndYear < f.FirstYear))) && d == date.NewDate(year, f.Month, f.Day)
}

// Match returns true and the holidays name if a holiday encoded in c occurs on the specified date.
func (c ChangingHoliday) Match(d date.Date) []string {
	return c(d)
}
