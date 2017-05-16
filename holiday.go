package date

import "time"

// Holidays contains a specification of holidays, which can be matched with dates
type Holidays struct {
	Changing []ChangingHoliday
	Fixed    []FixedHoliday
}

// ChangingHoliday encodes arbitrarily defined holidays, which can be matched with dates
type ChangingHoliday func(year int, month time.Month, day int) (string, bool)

// FixedHoliday encodes holidays that repeat annually on the same date
type FixedHoliday struct {
	Month              time.Month
	Day                int
	FirstYear, EndYear int // First year with and without the holiday (EndYear is ignored if it is equal to FirstYear)
	Name               string
}

// Match returns true if the holiday occurs on the specified date.
func (f *FixedHoliday) Match(d Date) bool {
	return d.Month() == f.Month && d.Day() == f.Day &&
		d.Year() >= f.FirstYear && (f.EndYear == f.FirstYear || d.Year() < f.EndYear)
}

// Match returns true and the holidays name if a holiday encoded in c occurs on the specified date.
func (c ChangingHoliday) Match(d Date) (string, bool) {
	return c(d.Year(), d.Month(), d.Day())
}

// Find returns a holiday matching the specified date if there are any.
func (h *Holidays) Find(d Date) (string, bool) {
	for _, f := range h.Fixed {
		if f.Match(d) {
			return f.Name, true
		}
	}
	for i := range h.Changing {
		if name, ok := h.Changing[i].Match(d); ok {
			return name, true
		}
	}
	return "", false
}
