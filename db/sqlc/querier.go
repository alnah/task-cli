// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateTask(ctx context.Context, description string) (Task, error)
	DeleteTask(ctx context.Context, id int32) error
	ListTasks(ctx context.Context, arg ListTasksParams) ([]Task, error)
	ListTasksByStatus(ctx context.Context, arg ListTasksByStatusParams) ([]Task, error)
	UpdateTaskDescription(ctx context.Context, arg UpdateTaskDescriptionParams) (Task, error)
	UpdateTaskStatus(ctx context.Context, arg UpdateTaskStatusParams) (Task, error)
}

var _ Querier = (*Queries)(nil)
