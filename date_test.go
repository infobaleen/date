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

func TestTime(t *testing.T) {
	var dRef, hRef, mRef, sRef, nRef = ThursdayJune15th2017, 1, 2, 3, 4
	var dhmsn = dRef.Time(hRef, mRef, sRef, nRef, time.UTC)
	var d, h, m, s, n = NewDate(dhmsn.Date()), dhmsn.Hour(), dhmsn.Minute(), dhmsn.Second(), dhmsn.Nanosecond()
	if d != dRef || h != hRef || m != mRef || s != sRef || n != nRef {
		t.Error(d, h, m, s, n, "!=", dRef, hRef, mRef, sRef, nRef)
	}
}
