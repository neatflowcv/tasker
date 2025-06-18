package flow

import (
	"context"
	"fmt"

	"github.com/neatflowcv/tasker/internal/pkg/domain"
	"github.com/neatflowcv/tasker/internal/pkg/repository/core"
)

type Service struct {
	repo core.Repository
}

func NewService(repo core.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTask(ctx context.Context, spec *domain.TaskSpec) (*domain.Task, error) {
	task, err := s.repo.CreateTask(spec)
	if err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	return task, nil
}

func (s *Service) ListTasks(ctx context.Context) ([]*domain.Task, error) {
	tasks, err := s.repo.ListTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}

	return tasks, nil
}

func (s *Service) GetTask(ctx context.Context, id domain.TaskID) (*domain.Task, error) {
	task, err := s.repo.GetTask(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	return task, nil
}

func (s *Service) UpdateTask(ctx context.Context, id domain.TaskID, spec *domain.TaskSpec) (*domain.Task, error) {
	task, err := s.repo.GetTask(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	updatedTask := task.SetSpec(spec)

	updatedTask, err = s.repo.UpdateTask(updatedTask)
	if err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	return updatedTask, nil
}

func (s *Service) DeleteTask(ctx context.Context, id domain.TaskID) error {
	err := s.repo.DeleteTask(id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	return nil
}
