package habit

import (
	"strings"
	"time"
)

func testEq(a, b []string) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func stringContains(s []string, val string) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}
	return false
}

func intContains(s []int, val int) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}
	return false
}

const dayDuration time.Duration = time.Duration(24) * time.Hour

func daysBetween(from time.Time, to time.Time, expectedDays []string) []time.Time {
	days := []time.Time{}
	reversed := []time.Time{}

	for !from.After(to) {
		if stringContains(expectedDays, strings.ToLower(from.Weekday().String())) {
			days = append(days, from)
		}
		from = from.Add(dayDuration)
	}

	for i := range days {
		n := days[len(days)-1-i]
		reversed = append(reversed, n)
	}

	return reversed
}

func dateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
