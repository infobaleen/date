package holiday

import (
	"time"

	"github.com/infobaleen/date"
)

// Predefined fixed date holidays
var (
	NewYearsDay             = FixedHoliday{Month: 1, Day: 1, Name: "New years day"}
	NewYearsEve             = FixedHoliday{Month: 12, Day: 31, Name: "New Year's Eve"}
	Epiphany                = FixedHoliday{Month: 1, Day: 6, Name: "Epiphany"}
	InternationalWorkersDay = FixedHoliday{Month: 5, Day: 1, Name: "International worker's day"}
	ChristmasEve            = FixedHoliday{Month: 12, Day: 24, Name: "Christmas eve"}
	ChristmasDay            = FixedHoliday{Month: 12, Day: 25, Name: "Christmas day"}
	SecondChristmasDay      = FixedHoliday{Month: 12, Day: 26, Name: "Second day of Christmas"}
)

// EasterAndFriends matches holidays whose dates are defined relative to easter sunday (western easter).
// The matched holidays are: "Good Friday", Easter Sunday", "Easter Monday", "Ascension Day", "Pentecost".
var EasterAndFriends ChangingHoliday = func(d date.Date) []string {
	var year = d.Year()
	easterMonth, easterDay := westernEasterDate(year)
	var easter = date.NewDate(year, easterMonth, easterDay)
	switch d {
	case easter - 2:
		return []string{"Good Friday"}
	case easter:
		return []string{"Easter Sunday"}
	case easter + 1:
		return []string{"Easter Monday"}
	case easter + 39:
		return []string{"Ascension Day"}
	case easter + 49:
		return []string{"Pentecost"}
	}
	return nil
}

// MidsummerFriday matches the Friday during the period from June 19th to 25th.
var MidsummerFriday ChangingHoliday = func(d date.Date) []string {
	if d.Weekday() == time.Friday && d.Month() == time.June {
		var day = d.Day()
		if 19 <= day && day <= 25 {
			return []string{"Midsummer Eve"}
		}
	}
	return nil
}

// MidsummerSaturday matches the Saturday during the period from June 20th to 26th.
var MidsummerSaturday ChangingHoliday = func(d date.Date) []string {
	if d.Weekday() == time.Saturday && d.Month() == time.June {
		var day = d.Day()
		if 20 <= day && day <= 26 {
			return []string{"Midsummer Day"}
		}
	}
	return nil
}

// Here be dragons... ...and I am not very happy about that.
// This hopefully returns the western easter dates after year 0 correctly.
// Knuth documents some predecessor of this, which gives me hope.
// Help is appreciated if you think you know or can do better.
func westernEasterDate(year int) (month time.Month, day int) {
	a := year % 19
	if year >= 1583 {
		b := year / 100
		c := year % 100
		d := b / 4
		e := b % 4
		f := (b + 8) / 25
		g := (b - f + 1) / 3
		h := (19*a + b - d - g + 15) % 30
		i := c / 4
		k := c % 4
		l := (32 + 2*e + 2*i - h - k) % 7
		m := (a + 11*h + 22*l) / 451
		day = 21 + h + l - 7*m
	} else {
		b := year % 7
		c := year % 4
		d := (19*a + 15) % 30
		e := (2*c + 4*b - d + 34) % 7
		day = 21 + d + e
	}
	month = time.Month(3 + day/31)
	day = day%31 + 1
	return
}
