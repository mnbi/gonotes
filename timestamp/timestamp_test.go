// gonotes/timestamp/timestamp_test.go
package timestamp

import (
	"testing"
	"time"
)

func TestString(t *testing.T) {
	tests := []struct {
		id       int
		testcase string
		expected string
	}{
		{100, "2024-11-22 06:59:30", "20241122065930"},
	}

	for _, tc := range tests {
		tt, _ := time.Parse(time.DateTime, tc.testcase)
		ts := NewTimestamp(tt, 0)
		got := ts.String()
		if tc.expected != got {
			t.Fatalf("tests[%d] - String() wrong, expected=%s, got=%s",
				tc.id, tc.expected, got)
		}
	}
}

func TestTimeToStr(t *testing.T) {
	tests := []struct {
		id       int
		testcase string
		expected string
	}{
		{200, "2024-11-22 06:23:00", "20241122062300"},
	}

	for _, tc := range tests {
		tt, _ := time.Parse(time.DateTime, tc.testcase)
		got := timeToStr(tt)
		if tc.expected != got {
			t.Fatalf("tests[%d] - StampString wrong, expected=%s, got=%s",
				tc.id, tc.expected, got)
		}
	}
}
