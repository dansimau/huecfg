package hue

import (
	"fmt"
	"strings"
	"time"
)

const absoluteTimeFormat = "2006-01-02T15:04:05"

// AbsoluteTime is a custom time struct that supports JSON marshalling from
// the Hue Bridge format.
type AbsoluteTime struct {
	time.Time
}

func (t *AbsoluteTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return
	}
	t.Time, err = time.Parse(absoluteTimeFormat, s)
	return
}

func (t *AbsoluteTime) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", t.Time.Format(absoluteTimeFormat))), nil
}
