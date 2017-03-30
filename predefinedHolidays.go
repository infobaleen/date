package date

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
var EasterAndFriends ChangingHoliday = func(year, month, day int) (string, bool) {
	eMonth, eDay := westernEasterDate(year)
	e := NewDate(year, eMonth, eDay, nil)
	d := NewDate(year, month, day, nil)
	switch d {
	case e.Add(0, 0, -2):
		return "Good Friday", true
	case e:
		return "Easter Sunday", true
	case e.Add(0, 0, 1):
		return "Easter Monday", true
	case e.Add(0, 0, 39):
		return "Ascension Day", true
	case e.Add(0, 0, 49):
		return "Pentecost", true
	}
	return "", false
}

// Here be dragons... ...and I am not very happy about that.
// This hopefully returns the western easter dates after year 0 correctly.
// Knuth documents some predecessor of this, which gives me hope.
// Help is appreciated if you think you know or can do better.
func westernEasterDate(year int) (month, day int) {
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
	month = 3 + day/31
	day = day%31 + 1
	return
}
