package habbits

// HabbitNotFoundError is raised when a habbit is not found in DB
type HabbitNotFoundError struct {
	msg string
}

func (e *HabbitNotFoundError) Error() string { return e.msg }
