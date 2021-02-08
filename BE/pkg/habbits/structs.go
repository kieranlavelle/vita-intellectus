package habbits

// Habbit represents a habbit a user wants to set
type Habbit struct {
	HabbitID       int      `json:"habbit_id"`
	UserID         int      `json:"user_id"`
	Name           string   `json:"name"`
	Days           []string `json:"days"`
	CompletedToday bool     `json:"completed_today"`
}

// CompleteHabbitBody represents the body expected by the completeHabbit request
type CompleteHabbitBody struct {
	HabbitID int `json:"habbit_id"`
}
