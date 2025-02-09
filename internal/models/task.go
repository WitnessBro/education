package models

//TODO USer_id, created_at, notification_rules, deadline (time)?,
type Task struct {
	Description *string    `json:"description,omitempty"`
	Id          string     `json:"id"`
	Status      TaskStatus `json:"status"`
	Title       string     `json:"title"`
}

const (
	Completed  TaskStatus = "completed"
	InProgress TaskStatus = "in_progress"
	Pending    TaskStatus = "pending"
)

// TaskStatus defines model for Task.Status.
type TaskStatus string
