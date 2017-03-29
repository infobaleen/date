package date

import (
	"time"
)

// TzStockholm represents the central european timezone.
// This is equivalent to CET or CEST as obvserved in Sweden, depending on the date.
var TzStockholm *time.Location

// Set up timezones
func init() {
	var err error
	TzStockholm, err = time.LoadLocation("Europe/Stockholm")
	if err != nil {
		panic("Could not load Europe/Stockholm timezone information")
	}
}
