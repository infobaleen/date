package date

// HolidaysByLocation is a map of holidays indexed by the location string returned by time.Location.String().
// This map is used by the Date.IsHoliday() method.
var HolidaysByLocation = map[string]Holidays{
	"Europe/Stockholm": {Fixed: []FixedHoliday{
		NewYearsEve, NewYearsDay, Epiphany, InternationalWorkersDay, ChristmasEve, ChristmasDay, SecondChristmasDay,
		{Month: 6, Day: 6, Name: "National day of Sweden"},
	}, Changing: []ChangingHoliday{
		EasterAndFriends,
	}},
}
