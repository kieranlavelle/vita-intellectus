package habit

// Error is a general error that is safe to return
// to the user
type Error struct {
	msg string
}

func (e *Error) Error() string { return e.msg }
