package task

import (
	"fmt"
	"time"
)

type Task struct {
	ID          string
	Title       string
	CreatedAt   time.Time
	Description string
	Completed   bool
}

func NewTask(title, description string) *Task {
	return &Task{
		Title:       title,
		CreatedAt:   time.Now(),
		Description: description,
	}
}

func (t *Task) Complete() {
	t.Completed = true
}

func (t *Task) Uncomplete() {
	t.Completed = false
}

func (t *Task) Update(title, description string) {
	t.Title = title
	t.Description = description
}

func (t *Task) Validate() error {
	if t.Title == "" {
		return fmt.Errorf("title is required")
	}
	return nil
}
