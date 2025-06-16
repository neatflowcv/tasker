package flow

import (
	"context"

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
	return s.repo.CreateTask(spec)
}

func (s *Service) ListTasks(ctx context.Context) ([]*domain.Task, error) {
	return s.repo.ListTasks()
}

func (s *Service) GetTask(ctx context.Context, id domain.TaskID) (*domain.Task, error) {
	return s.repo.GetTask(id)
}

func (s *Service) UpdateTask(ctx context.Context, id domain.TaskID, spec *domain.TaskSpec) (*domain.Task, error) {
	task, err := s.repo.GetTask(id)
	if err != nil {
		return nil, err
	}
	updatedTask := task.SetSpec(spec)
	return s.repo.UpdateTask(updatedTask)
}

func (s *Service) DeleteTask(ctx context.Context, id domain.TaskID) error {
	return s.repo.DeleteTask(id)
}
