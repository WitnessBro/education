package models

import "time"

//TODO USer_id, created_at, notification_rules, deadline (time)?,
type Task struct {
	TaskDescription string `json:"task_description,omitempty"`
	//Id              string     `json:"id"`
	TaskStatus TaskStatus `json:"task_status"`
	Title      string     `json:"title"`
	CreatedAt  time.Time  `json:"created_at"`
	UserId     int        `json:"user_id"`
}

const (
	Completed  TaskStatus = "completed"
	InProgress TaskStatus = "in_progress"
	Pending    TaskStatus = "pending"
)

type TaskStatus string
