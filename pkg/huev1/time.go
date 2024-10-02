package huev1

import (
	"fmt"
	"strings"
	"time"
)

const (
	AbsoluteTimeFormat              = "2006-01-02T15:04:05"
	AbsoluteTimeHumanReadableFormat = "2006-01-02 15:04:05"
)

// AbsoluteTime is a custom time struct that supports JSON marshalling from
// the Hue Bridge format.
type AbsoluteTime struct {
	time.Time
}

func (t AbsoluteTime) String() string {
	if t.Time.IsZero() {
		return ""
	}
	return t.Time.Format(AbsoluteTimeHumanReadableFormat)
}

func (t *AbsoluteTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "none" {
		t.Time = time.Time{}
		return
	}
	t.Time, err = time.Parse(AbsoluteTimeFormat, s)
	return
}

func (t *AbsoluteTime) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", t.Time.Format(AbsoluteTimeFormat))), nil
}
