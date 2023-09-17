package task

import "context"

type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
	Update(ctx context.Context, task *Task) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*Task, error)
	GetAll(ctx context.Context) ([]*Task, error)
}
