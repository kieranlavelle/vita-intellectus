package habit

import (
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var allDays = []string{
	"monday",
	"tuesday",
	"wednesday",
	"thursday",
	"friday",
	"saturday",
	"sunday",
}

var indexToDay = map[int]string{
	0: "monday",
	1: "tuesday",
	2: "wednesday",
	3: "thursday",
	4: "friday",
	5: "saturday",
	6: "sunday",
}

var dayToIndex = map[string]int{
	"monday":    0,
	"tuesday":   1,
	"wednesday": 2,
	"thursday":  3,
	"friday":    4,
	"saturday":  5,
	"sunday":    6,
}

// Habit represents a habbit a user wants to set
type Habit struct {
	ID            int                    `json:"id"`
	UID           int                    `json:"user_id"`
	Name          string                 `json:"name"`
	Days          []string               `json:"days"`
	Tags          []string               `json:"tags"`
	SelectedStats []string               `json:"selected_stats"`
	Description   string                 `json:"description"`
	Completed     bool                   `json:"completed"`
	Statistics    map[string]interface{} `json:"statistics"`
}

type HabitCompletions struct {
	Completions []HabitCompletion `json:"completions"`
}

type HabitCompletion struct {
	HabitID int       `json:"habit_id"`
	Time    time.Time `json:"time"`
	Notes   string    `json:"notes"`
}

// Load returns a habit corresponding to the
// habit id and user id passed if one is found.
func Load(hid, uid int, c *pgxpool.Pool) (Habit, error) {
	h := Habit{ID: hid, UID: uid}
	err := getHabit(&h, c)
	switch err {
	case pgx.ErrNoRows:
		return h, &Error{"failed to find habit"}
	}
	return h, err
}

// New creates an empty habit with the default
// configuration
func New(uid int) Habit {
	return Habit{UID: uid, Days: allDays, Tags: []string{}}
}

// Habits returns all of the habits for the caller
func Habits(uid int, c *pgxpool.Pool) ([]Habit, error) {

	habitRows, err := userHabits(uid, c)
	defer habitRows.Close()
	if err != nil {
		return []Habit{}, err
	}

	habits := []Habit{}

	for habitRows.Next() {
		h := Habit{}
		err = habitRows.Scan(
			&h.ID, &h.UID, &h.Name, &h.Days,
			&h.Tags, &h.Description, &h.SelectedStats,
			&h.Completed,
		)
		if err != nil {
			return []Habit{}, err
		}

		info, err := h.Info(c)
		if err != nil {
			return []Habit{}, err
		}
		h.Statistics = info

		habits = append(habits, h)
	}

	return habits, err
}

func (h *Habit) updateable(h1 Habit) bool {

	if !testEq(h.Days, h1.Days) {
		return false
	}

	return true
}

// Update inserts a habit into the database or
// updates it if it already exists
func (h *Habit) Update(c *pgxpool.Pool) error {

	// see if this habit exists
	exists := true
	dbHabit := Habit{ID: h.ID, UID: h.UID}

	err := getHabit(&dbHabit, c)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			exists = false
		default:
			return err
		}
	}

	if exists {
		if h.updateable(dbHabit) {
			return updateHabit(h, c)
		}
		return &Error{"cant update property [days] of habit"}
	}

	// if there are no days then default to all
	if len(h.Days) == 0 {
		h.Days = allDays
	}

	if len(h.SelectedStats) > 2 {
		return &Error{"a habit can only have 2 selected statistics."}
	}

	return insertHabit(h, c)
}

// Delete removes a habit from the database
// doesnt check ownership as this is done in
// the Load function
func (h *Habit) Delete(c *pgxpool.Pool) error {
	err := deleteHabit(h, c)
	return err
}

// Complete add's an entry into the habit_completions
// table for a given habit if there isnt one already
// otherwise it errors
func (h *Habit) Complete(notes string, c *pgxpool.Pool) error {
	if h.Completed {
		return &Error{"can only complete a habit once a day"}
	}
	h.Completed = true
	return completeHabit(h, notes, c)
}

// UnComplete removes an entry from the habit_completions
// table for a given habit if there is one already
// otherwise it errors
func (h *Habit) UnComplete(c *pgxpool.Pool) error {
	if !h.Completed {
		return &Error{"Habit is not completed"}
	}
	h.Completed = false
	return unCompleteHabit(h, c)
}

// Completions get's all of the times this habit
// has been completed
func (h *Habit) Completions(c *pgxpool.Pool) (HabitCompletions, error) {

	completions, err := habitCompletions(h, c)
	defer completions.Close()
	if err != nil {
		return HabitCompletions{}, err
	}

	hCompletions := HabitCompletions{}
	for completions.Next() {
		hc := HabitCompletion{}
		err := completions.Scan(&hc.HabitID, &hc.Time, &hc.Notes)
		if err != nil {
			return hCompletions, err
		}

		hCompletions.Completions = append(hCompletions.Completions, hc)
	}

	return hCompletions, err
}

// Info gets information about the habit such
// as the number of times it has been completed in a row
func (h *Habit) Info(c *pgxpool.Pool) (map[string]interface{}, error) {

	habitStatistics := map[string]interface{}{}
	completions, err := h.Completions(c)
	if err != nil {
		return habitStatistics, err
	}

	for statisticName, value := range statistics {
		habitStatistics[statisticName] = value(h, completions)
	}
	return habitStatistics, err
}
