package habit

import (
	"time"

	"github.com/jackc/pgx/v4"
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
	ID        int      `json:"id"`
	UID       int      `json:"user_id"`
	Name      string   `json:"name"`
	Days      []string `json:"days"`
	Tags      []string `json:"tags"`
	Completed bool     `json:"completed"`
}

type HabitInfo struct {
	Streak int `json:"streak"`
}

type HabitCompletions struct {
	Completions []HabitCompletion `json:"completions"`
}

type HabitCompletion struct {
	HabitID int       `json:"habit_id"`
	Time    time.Time `json:"time"`
}

// Load returns a habit corresponding to the
// habit id and user id passed if one is found.
func Load(hid, uid int, c *pgx.Conn) (Habit, error) {
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
func Habits(uid int, c *pgx.Conn) ([]Habit, error) {

	habitRows, err := userHabits(uid, c)
	if err != nil {
		return []Habit{}, err
	}

	habits := []Habit{}

	for habitRows.Next() {
		h := Habit{}
		err = habitRows.Scan(&h.ID, &h.UID, &h.Name, &h.Days, &h.Tags, &h.Completed)
		if err != nil {
			return []Habit{}, err
		}
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
func (h *Habit) Update(c *pgx.Conn) error {

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
	return insertHabit(h, c)
}

// Delete removes a habit from the database
// doesnt check ownership as this is done in
// the Load function
func (h *Habit) Delete(c *pgx.Conn) error {
	err := deleteHabit(h, c)
	return err
}

// Complete add's an entry into the habit_completions
// table for a given habit if there isnt one already
// otherwise it errors
func (h *Habit) Complete(c *pgx.Conn) error {
	if h.Completed {
		return &Error{"can only complete a habit once a day"}
	}
	h.Completed = true
	return completeHabit(h, c)
}

// Completions get's all of the times this habit
// has been completed
func (h *Habit) Completions(c *pgx.Conn) (HabitCompletions, error) {

	completions, err := habitCompletions(h, c)
	defer completions.Close()
	if err != nil {
		return HabitCompletions{}, err
	}

	hCompletions := HabitCompletions{}
	for completions.Next() {
		hc := HabitCompletion{}
		err := completions.Scan(&hc.HabitID, &hc.Time)
		if err != nil {
			return hCompletions, err
		}

		hCompletions.Completions = append(hCompletions.Completions, hc)
	}

	return hCompletions, err
}

// Info gets information about the habit such
// as the number of times it has been completed in a row
func (h *Habit) Info(c *pgx.Conn) (HabitInfo, error) {

	hInfo := HabitInfo{}
	completions, err := h.Completions(c)
	if err != nil {
		return HabitInfo{}, err
	}

	hInfo.Streak = calculateStreak(h, completions)

	// consecutive completions

	// percent this month

	// percent overall

	return hInfo, err

}
