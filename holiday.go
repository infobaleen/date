package date

// Fixed annual swedish holidays
// Years encodes the starting and end year
type fixedAnnualHoliday struct {
	date    Date // Starting date of annual holiday (first year)
	endYear int  // First year without the holiday (ignored if equal to start year)
	name    string
}

// Swedish fixed holidays
var fixedAnnualHolidays = [...]fixedAnnualHoliday{
	{date: Date{month: 1, day: 1}, name: "New years day"},
	{date: Date{month: 1, day: 6}, name: "Epiphany"},
	{date: Date{month: 5, day: 1}, name: "International worker's day"},
	{date: Date{month: 6, day: 6}, name: "National day of Sweden"},
	{date: Date{month: 12, day: 24}, name: "Christmas eve"},
	{date: Date{month: 12, day: 25}, name: "Christmas day"},
	{date: Date{month: 12, day: 26}, name: "Second day of Christmas"},
	{date: Date{month: 12, day: 31}, name: "New Year's Eve"},
}

func (f fixedAnnualHoliday) matches(d Date) bool {
	return d.month == f.date.month && d.day == f.date.day &&
		d.year >= f.date.year && (f.endYear == f.date.year || d.year < f.endYear)
}

func (d Date) isHoliday() bool {
	// Check fixed Annual Holidays
	for _, f := range fixedAnnualHolidays {
		if f.matches(d) {
			return true
		}
	}
	return false
}
