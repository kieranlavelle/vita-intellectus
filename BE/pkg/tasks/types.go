package tasks

type DisplayableError struct {
	s string
}

func (e *DisplayableError) Error() string {
	return e.s
}
