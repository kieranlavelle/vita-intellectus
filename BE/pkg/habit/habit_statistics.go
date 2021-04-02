package habit

import (
	"strings"
	"time"
)

type StatisticsFunc = func(*Habit, HabitCompletions) int

var statistics = map[string]StatisticsFunc{
	"streak":                    calculateStreak,
	"consecutive":               consecutiveCompletions,
	"days_since_last_completed": daysSinceLastCompletion,
	"total_completions":         numCompletions,
}

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

func numCompletions(h *Habit, hc HabitCompletions) int {
	return len(hc.Completions)
}

func daysSinceLastCompletion(h *Habit, hc HabitCompletions) int {
	if len(hc.Completions) == 0 {
		// this habit has never been completed
		return -1
	}

	lastCompletion := hc.Completions[len(hc.Completions)-1]
	daysSince := time.Now().Sub(lastCompletion.Time).Hours() / 24
	if daysSince > 0 && (time.Now().Day() != lastCompletion.Time.Day()) {
		return 1
	}
	return int(daysSince)
}
