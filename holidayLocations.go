package date

import "time"

// HolidaysByLocation is a map of holidays indexed by the location string returned by time.Location.String().
// This map is used by the Date.IsHoliday() method.
var HolidaysByLocation = map[string]Holidays{
	"Europe/Stockholm": {Fixed: []FixedHoliday{
		NewYearsEve, NewYearsDay, Epiphany, InternationalWorkersDay, ChristmasEve, ChristmasDay, SecondChristmasDay,
		{Month: 6, Day: 6, Name: "National day of Sweden"},
	}, Changing: []ChangingHoliday{
		EasterAndFriends, MidsummerFriday, allSaintsDaySweden,
	}},
}

// allSaintsDaySweden matches the Saturday betweem October 31st and November 6th.
var allSaintsDaySweden ChangingHoliday = func(year, month, day int) (string, bool) {
	dayBefore := NewDate(year, int(time.October), 30, nil)
	dayAfter := NewDate(year, int(time.November), 7, nil)
	d := NewDate(year, month, day, nil)
	if d.After(dayBefore) && d.Before(dayAfter) && d.IsWeekday(time.Saturday) {
		return "All Saints' Day", true
	}
	return "", false
}
