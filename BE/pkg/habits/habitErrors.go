package habits

// HabitNotFoundError is raised when a habbit is not found in DB
type HabitNotFoundError struct {
	msg string
}

func (e *HabitNotFoundError) Error() string { return e.msg }
