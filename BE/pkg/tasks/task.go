package tasks

import (
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func weekdayToString(weekday int) string {
	switch weekday {
	case 1:
		return "monday"
	case 2:
		return "tuesday"
	case 3:
		return "wednesday"
	case 4:
		return "thursday"
	case 5:
		return "friday"
	case 6:
		return "saturday"
	case 7:
		return "sunday"
	}

	// this is a bad idea but CBA
	return "monday"
}

// Tast represents a task a user owns
type Task struct {
	ID          int                    `json:"id"`
	UID         int                    `json:"user_id"`
	Name        string                 `json:"name"`
	State       string                 `json:"state"`
	Tags        []string               `json:"tags"`
	Description string                 `json:"description"`
	Recurring   bool                   `json:"recurring"`
	Days        []string               `json:"days,omitempty"`
	Date        time.Time              `json:"date"`
	DateCreated time.Time              `json:"date_created"`
	Extra       map[string]interface{} `json:"extra"`
}

func New(t *Task, c *pgxpool.Pool) (*Task, error) {

	// initialise days and tags if they're nil so they're
	// empty slices instead of nil slices for DB insert
	if t.Days == nil {
		t.Days = make([]string, 0)
	}
	if t.Tags == nil {
		t.Tags = make([]string, 0)
	}

	newTask, err := createEmptyTask(t, c)
	if err != nil {
		return newTask, err
	}

	err = newTask.SetState(time.Now(), c)
	return newTask, err
}

func Load(id, uid int, date time.Time, c *pgxpool.Pool) (*Task, error) {
	t, err := getTask(id, uid, c)
	if err != nil {
		return t, err
	}

	err = t.SetState(date, c)
	return t, err
}

func Tasks(uid int, date time.Time, c *pgxpool.Pool) ([]*Task, error) {
	tasks, err := getTasks(uid, c)
	if err != nil {
		return tasks, err
	}

	for _, task := range tasks {
		task.SetState(date, c)
	}

	return tasks, err
}

func (t *Task) Complete(notes string, date time.Time, c *pgxpool.Pool) (*Task, error) {
	if t.State == "completed" {
		err := &DisplayableError{s: fmt.Sprintf("habit already completed on: %v", date)}
		return t, err
	}

	err := completeTask(t, notes, date, c)
	if err != nil {
		return t, err
	}

	t.State = "completed"

	return t, err
}

func (t *Task) SetState(date time.Time, c *pgxpool.Pool) error {

	createdY, createdMonth, createdD := t.DateCreated.Date()
	currentY, currentMonth, currentD := time.Now().Date()
	taskY, taskMonth, taskD := t.Date.Date()
	y, month, d := date.Date()

	createdM := int(createdMonth)
	taskM := int(taskMonth)
	currentM := int(currentMonth)
	m := int(month)

	if (y < createdY) || (m < createdM) || (d < createdD) {
		t.State = "not-due"
		return nil
	}

	if t.Recurring {
		weekday := strings.ToLower(date.Weekday().String())

		completed, err := checkAbsoluteCompletion(t.ID, d, m, y, c)
		if err != nil {
			return err
		}

		// for this task there is an entry in completions on the given date
		// so it is completed
		if completed {
			t.State = "completed"
			return err
		}

		// if not completed and is/was due then the state is due
		for _, value := range t.Days {

			// if date == today
			if (currentY == y) && (currentM == m) && (currentD == d) {
				if !completed && strings.ToLower(value) == weekday {
					t.State = "due"
					return err
				}
			} else if (currentY < y) && (currentM < m) && (currentD < d) {
				t.State = "not-due"
				return err
			} else {
				// it wasn't completed and the date is in the past where
				// it was due
				if !completed && strings.ToLower(value) == weekday {
					t.State = "missed"
					return err
				}
			}
		}

		return err

	} else {

		completed, err := checkAbsoluteCompletion(t.ID, d, m, y, c)
		if err != nil {
			return err
		}

		// task isnt due
		if (y > taskY) || (m > taskM) || (d > taskD) {
			t.State = "not-due"
			return err
		}

		// task is completed
		if completed {
			t.State = "completed"
			return err
		}

		// task is due
		if (taskY == y) && (taskM == m) && (taskD == d) {
			t.State = "due"
			return err
		}

		// task was missed
		if (y < taskY) || (m < taskM) || (d < taskD) {
			t.State = "missed"
			return err
		}
	}

	return nil
}

// func Load(id, uid int, c *pgxpool.Pool) (Task, error) {

// 	t := Task{}
// }
