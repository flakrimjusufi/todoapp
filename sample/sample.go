package sample

import (
	"time"
	"todoapp/model"
)

type Sample struct{}

func (s *Sample) GetSampleData() ([]model.Todo, error) {
	todoList := []model.Todo{
		{
			ID:       1,
			Labels:   "Clean dishes",
			Comments: "Don't forget the garbage too..",
			DueDate:  time.Date(2021, 10, 28, 7, 40, 0, 0, time.UTC),
		},
		{
			ID:       2,
			Labels:   "Send kids to school",
			Comments: "Don't forget the Umbrella..",
			DueDate:  time.Date(2021, 10, 28, 8, 10, 0, 0, time.UTC),
		},
		{
			ID:       3,
			Labels:   "Go to work",
			Comments: "Do some small-talks with colleagues before starting the work..",
			DueDate:  time.Date(2021, 10, 28, 8, 30, 0, 0, time.UTC),
		},
	}

	return todoList, nil
}
