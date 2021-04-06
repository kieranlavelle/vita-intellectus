package tasks

import (
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kieranlavelle/vita-intellectus/pkg/helpers"
)

// Tast represents a task a user owns
type Task struct {
	ID          int                    `json:"id"`
	UID         int                    `json:"user_id"`
	Name        string                 `json:"name"`
	State       string                 `json:"state"`
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
		if t.Recurring {
			t.Days = []string{
				"monday",
				"tuesday",
				"wednesday",
				"thursday",
				"friday",
				"saturday",
				"sunday",
			}
		} else {
			t.Days = make([]string, 0)
		}
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

func Tasks(uid int, filter string, date time.Time, c *pgxpool.Pool) ([]*Task, error) {
	tasks, err := getTasks(uid, c)
	if err != nil {
		return tasks, err
	}

	if strings.ToLower(filter) != "all" && strings.ToLower(filter) != "due" {
		err = &DisplayableError{s: "please provide a valid filter"}
		return tasks, err
	}

	filteredTasks := make([]*Task, 0)
	for _, task := range tasks {

		// if this is a fixed date task and the passed date is not
		// the required task date then skip it
		if !task.Recurring && !helpers.DateEquals(task.Date, date) {
			continue
		}

		task.SetState(date, c)

		if task.State == "not-created" {
			continue
		}

		if strings.ToLower(filter) == "all" {
			filteredTasks = append(filteredTasks, task)
		} else if task.State != "not-due" && strings.ToLower(filter) == "due" {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return filteredTasks, err
}

func (t *Task) Delete(c *pgxpool.Pool) error {
	return deleteTask(t, c)
}

func (t *Task) Edit(newTask *Task, date time.Time, c *pgxpool.Pool) (*Task, error) {

	if newTask.Name != "" {
		t.Name = newTask.Name
	}

	t.Description = newTask.Description

	err := updateTask(t, c)
	if err != nil {
		return t, err
	}

	return t, nil

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

func (t *Task) UnComplete(date time.Time, c *pgxpool.Pool) (*Task, error) {
	if t.State != "completed" {
		err := &DisplayableError{s: fmt.Sprintf("habit not completed on: %v", date)}
		return t, err
	}

	err := unCompleteTask(t, date, c)
	if err != nil {
		return t, err
	}

	t.SetState(date, c)
	return t, err
}

func (t *Task) SetState(date time.Time, c *pgxpool.Pool) error {

	createdY, createdMonth, createdD := t.DateCreated.Date()
	y, month, d := date.Date()

	createdM := int(createdMonth)
	m := int(month)

	// if the passed date comes before the task was created
	// then the task is not due.
	if (y < createdY) || (m < createdM) || (d < createdD) {
		t.State = "not-created"
		return nil
	}

	if t.Recurring {
		day := strings.ToLower(date.Weekday().String())

		// for this task there is an entry in completions on the passed
		// date for the task so it is completed
		completed, err := checkAbsoluteCompletion(t.ID, d, m, y, c)
		if err != nil {
			return err
		} else if completed {
			t.State = "completed"
			return err
		}

		if helpers.DateEquals(date, time.Now()) {
			// if not completed and is/was due then the state is due
			for _, requiredDay := range t.Days {
				if !completed && strings.ToLower(requiredDay) == day {
					t.State = "due"
					return err
				}
			}
		} else if helpers.DateInPast(date, time.Now()) {
			for _, requiredDay := range t.Days {
				if !completed && strings.ToLower(requiredDay) == day {
					t.State = "missed"
					return err
				}
			}
		}

		t.State = "not-due"
		return err

	} else {

		completed, err := checkAbsoluteCompletion(t.ID, d, m, y, c)
		if err != nil {
			return err
		}

		// task is completed
		if completed {
			t.State = "completed"
			return err
		}

		if helpers.DateEquals(t.Date, time.Now()) {
			// if not completed and is/was due then the state is due
			t.State = "due"
			return err
		} else if helpers.DateInPast(date, time.Now()) {
			t.State = "missed"
			return err
		}

		t.State = "not-due"
		return err
	}
}

func (t *Task) numCompletions(date time.Time, c *pgxpool.Pool) (int, error) {
	return getNumCompletions(t, date, c)
}
