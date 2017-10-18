package holiday

import (
	"time"

	"github.com/infobaleen/date"
)

// ByLocation is a map of holiday calendars indexed by the location string returned by time.Location.String().
var ByLocation = map[string]Calendar{
	"Europe/Stockholm": {Fixed: []FixedHoliday{
		NewYearsEve, NewYearsDay, Epiphany, InternationalWorkersDay, ChristmasEve, ChristmasDay, SecondChristmasDay,
		{Month: 6, Day: 6, Name: "National day of Sweden"},
	}, Changing: []ChangingHoliday{
		EasterAndFriends, MidsummerFriday, MidsummerSaturday, allSaintsDaySweden,
	}},
}

// allSaintsDaySweden matches the Saturday betweem October 31st and November 6th.
var allSaintsDaySweden ChangingHoliday = func(d date.Date) []string {
	if d.Weekday() == time.Saturday {
		var year = d.Year()
		if date.NewDate(year, time.October, 30) < d && d < date.NewDate(year, time.November, 7) {
			return []string{"All Saints' Day"}
		}
	}
	return nil
}
