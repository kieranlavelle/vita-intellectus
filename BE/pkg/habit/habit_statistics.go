package habit

import (
	"strings"
	"time"
)

func calculateStreak(h *Habit, completions HabitCompletions) int {

	streak := 0

	// This gets us the completion times that fall on the days that the user
	// has specified as required days in h.Days
	completionDays := []time.Time{}
	for _, completion := range completions.Completions {
		weekday := strings.ToLower(completion.Time.Weekday().String())
		if stringContains(h.Days, weekday) {
			completionDays = append(completionDays, completion.Time)
		}
	}

	expectedDays := daysBetween(completionDays[len(completionDays)-1], time.Now(), h.Days)

	for i, expected := range expectedDays {
		if i >= len(completionDays) {
			break
		} else if !dateEqual(completionDays[i], expected) {
			break
		}
		streak++
	}
	return streak
}
