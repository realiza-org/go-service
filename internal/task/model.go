package task

// Task represents a simple task with ID, Title, and Description.
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
