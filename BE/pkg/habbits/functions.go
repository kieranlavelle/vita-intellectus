package habbits

import (
	"strings"
	"time"
)

var weekdayOrdering = map[string]int{
	"monday":    0,
	"tuesday":   1,
	"wednesday": 2,
	"thursday":  3,
	"friday":    4,
	"saturday":  5,
	"sunday":    6,
}

var days = [...]string{
	"monday",
	"tuesday",
	"wednesday",
	"thursday",
	"friday",
	"saturday",
	"sunday",
}

func contains(days []string, day string) bool {
	for _, sliceDay := range days {
		if sliceDay == day {
			return true
		}
	}
	return false
}

func (h Habbit) setDueDates() Habbit {

	dueDates := DueDates{}

	today := strings.ToLower(time.Now().Weekday().String())
	tomorrow := strings.ToLower(time.Now().Add(time.Hour * 24).Weekday().String())

	todaysIndex := weekdayOrdering[today]

	// Convert the Habbit's days into indexes
	daysAfterToday, daysBeforeToday := []int{}, []int{}
	for _, day := range h.Days {
		if index := weekdayOrdering[day]; index < todaysIndex {
			daysBeforeToday = append(daysBeforeToday, index)
		} else {
			daysAfterToday = append(daysAfterToday, index)
		}
	}

	weekFromToday := append(daysAfterToday, daysBeforeToday...)
	if len(weekFromToday) == 1 {
		dueDates.NextDue = days[weekFromToday[0]]
		dueDates.NextDueAfterCompleted = days[weekFromToday[0]]
	} else {
		dueDates.NextDue = days[weekFromToday[0]]
		dueDates.NextDueAfterCompleted = days[weekFromToday[1]]
	}

	if dueDates.NextDue == today {
		dueDates.NextDue = "Today"
	} else if dueDates.NextDue == tomorrow {
		dueDates.NextDue = "Tomorrow"
	}

	if dueDates.NextDueAfterCompleted == tomorrow {
		dueDates.NextDueAfterCompleted = "Tomorrow"
	}

	dueDates.NextDue = strings.Title(dueDates.NextDue)
	dueDates.NextDueAfterCompleted = strings.Title(dueDates.NextDueAfterCompleted)

	h.NextDue = dueDates
	return h
}
