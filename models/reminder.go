package models

import "time"

type Reminder struct {
	ID          int       `json:"id"`
	TaskID      int       `json:"task_id"`
	Description string    `json:"description"`
	DueDateTime time.Time `json:"due_date_time"`
}
