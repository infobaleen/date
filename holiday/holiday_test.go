package holiday

import (
	"sort"
	"testing"
	"time"

	"github.com/infobaleen/date"
)

func Test(t *testing.T) {
	var holidays = []struct {
		Date date.Date
		Name string
	}{
		{date.NewDate(2017, time.April, 16), "Easter Sunday"},
		{date.NewDate(2017, time.December, 24), "Christmas eve"},
	}
	var cal = ByLocation["Europe/Stockholm"]
	for _, h := range holidays {
		var matches = cal.Match(h.Date)
		var idx = sort.SearchStrings(matches, h.Name)
		if idx == len(matches) || matches[idx] != h.Name {
			t.Errorf("%s did not match %s", h.Date, h.Name)
		}
	}
}
