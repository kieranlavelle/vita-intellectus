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

	if len(completionDays) == 0 {
		return 0
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

func consecutiveCompletions(h *Habit, hc HabitCompletions) int {

	consecutive := 0
	for i, c := range hc.Completions {
		if i == 0 {

			// if the last completion was today then skip
			if dateEqual(time.Now(), c.Time) {
				continue
			}

			// if the last completion was not yesterday
			// then the consecutive is 0
			yesterday := time.Now().Add(-time.Hour * 24)
			if !dateEqual(yesterday, c.Time) {
				return 0
			}

			consecutive++
			continue
		}

		last := hc.Completions[i-1]
		dayBeforeLast := last.Time.Add(-time.Hour * 24)

		if !dateEqual(dayBeforeLast, c.Time) {
			return consecutive
		}
		consecutive++
	}
	return consecutive
}

func completionPercentage(h *Habit, hc HabitCompletions) float32 {

	// Get's all of the days it should have been completed in the
	// past 28days
	dayMinus28 := time.Now().Add(-dayDuration * 28)
	habitDaysThisMonth := daysBetween(dayMinus28, time.Now(), h.Days)

	// Represents the days we were expected to complete the habit
	// and we completed it
	expectedCompletions := []time.Time{}
	for _, completion := range hc.Completions {
		weekday := strings.ToLower(completion.Time.Weekday().String())
		if stringContains(h.Days, weekday) {
			expectedCompletions = append(expectedCompletions, completion.Time)
		}
	}

	nExpected := float32(len(expectedCompletions))
	nActual := float32(len(habitDaysThisMonth))
	if nActual == 0 {
		return 0
	}

	return nExpected / nActual
}
