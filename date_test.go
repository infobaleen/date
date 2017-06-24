package date

import (
	"testing"
	"time"
)

const ThursdayJune15th2017 Date = 736494

func TestNewDate(t *testing.T) {
	d := NewDate(2017, time.June, 15)
	if d != ThursdayJune15th2017 {
		t.Errorf("%d != %d", d, ThursdayJune15th2017)
	}
}

func TestWeekday(t *testing.T) {
	var ref = []struct {
		Date
		time.Weekday
	}{
		{ThursdayJune15th2017, time.Thursday},
		{0, time.Monday},
		{-1, time.Sunday},
	}
	for i := range ref {
		if ref[i].Date.Weekday() != ref[i].Weekday {
			t.Errorf("%d: %s != %s", i, ref[i].Date.Weekday(), ref[i].Weekday)
		}
	}
}
