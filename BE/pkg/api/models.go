package api

type TaskCompletion struct {
	Notes string `json:"notes"`
}

type EditTaskModel struct {
	Name        string   `json:"name"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
}
