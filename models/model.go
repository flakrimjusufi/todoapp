package models

import "time"

type Todo struct {
	ID       int       `gorm:"primary_key"`
	Labels   string    `json:"labels"`
	Comments string    `json:"comments"`
	DueDate  time.Time `json:"due_date"`
}
