package habbits

// Habbit represents a habbit a user wants to set
type Habbit struct {
	HabbitID       int      `json:"habbit_id"`
	UserID         int      `json:"user_id"`
	Name           string   `json:"name"`
	Days           []string `json:"days"`
	CompletedToday bool     `json:"completed_today"`
	NextDue        DueDates `json:"due_dates"`
}

// CompleteHabbitBody represents the body expected by the completeHabbit request
type CompleteHabbitBody struct {
	HabbitID int `json:"habbit_id"`
}

// DueDates represents when a habbit is next due depending on state
type DueDates struct {
	NextDue               string `json:"next_due"`
	NextDueAfterCompleted string `json:"next_due_on_completed"`
}
