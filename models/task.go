package models

import "time"

type Task struct {
	ID          int
	Title       string
	Description string
	Priority    int
	DueDateTime time.Time
}
