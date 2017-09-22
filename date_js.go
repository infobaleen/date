// +build js

package date

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jsbuiltin"
)

// Supports strings, Date and Moment from Moments.js
func NewDateFromJSObject(jsDate *js.Object) Date {
	const format = "2006-01-02"
	var str string
	if jsbuiltin.TypeOf(jsDate) == jsbuiltin.TypeString {
		str = jsDate.String()
	} else {
		str = jsDate.Call("toISOString").String()[:10]
	}
	var d, err = ParseDate(format, str)
	if err != nil {
		panic(err.Error())
	}
	return d
}
