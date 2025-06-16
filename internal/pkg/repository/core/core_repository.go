package core

import "github.com/neatflowcv/tasker/internal/pkg/domain"

type Repository interface {
	CreateTask(spec *domain.TaskSpec) (*domain.Task, error)
	ListTasks() ([]*domain.Task, error)
	GetTask(id domain.TaskID) (*domain.Task, error)
	UpdateTask(task *domain.Task) (*domain.Task, error)
	DeleteTask(id domain.TaskID) error
}
