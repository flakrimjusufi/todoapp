package model

import "time"

type Todo struct {
	ID       int       `json:"id"`
	Labels   string    `json:"labels"`
	Comments string    `json:"comments"`
	DueDate  time.Time `json:"due_date"`
}
