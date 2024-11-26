// gonotes/timestamp/timestamp.go
package timestamp

import (
	"fmt"
	"time"
)

type StampString string

type Timestamp struct {
	Time        time.Time
	Suffix      int
	StampString StampString
}

func NewTimestamp(t time.Time, suffix int) Timestamp {
	stampStr := timeToStr(t)
	if suffix != 0 {
		stampStr += fmt.Sprintf("%03d", suffix)
	}

	ts := Timestamp{Time: t, Suffix: suffix, StampString: StampString(stampStr)}
	return ts
}

func (ts *Timestamp) String() string {
	return string(ts.StampString)
}

func timeToStr(t time.Time) string {
	timeStr := t.Format(time.DateTime)
	return filterDelimiter(timeStr)
}

func filterDelimiter(timeStr string) string {
	runes := []rune(timeStr)
	result := make([]rune, 0, len(runes))
	for _, r := range runes {
		switch r {
		case '-', ' ', ':':
			continue
		default:
			result = append(result, r)
		}
	}
	return string(result)
}
