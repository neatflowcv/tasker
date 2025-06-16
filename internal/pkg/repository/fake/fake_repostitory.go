package fake

import (
	"fmt"

	"github.com/neatflowcv/tasker/internal/pkg/domain"
	"github.com/neatflowcv/tasker/internal/pkg/repository/core"
)

var _ core.Repository = (*Repository)(nil)

type Repository struct {
	tasks   map[string]*domain.Task
	counter int
}

// NewRepository creates a new fake repository
func NewRepository() *Repository {
	return &Repository{
		tasks:   make(map[string]*domain.Task),
		counter: 0,
	}
}

// CreateTask implements core.Repository.
func (r *Repository) CreateTask(spec *domain.TaskSpec) (*domain.Task, error) {
	r.counter++
	id := domain.TaskID(fmt.Sprintf("task-%d", r.counter))

	task := domain.NewTask(id, spec.Title(), spec.Description())
	r.tasks[string(id)] = task

	return task, nil
}

// DeleteTask implements core.Repository.
func (r *Repository) DeleteTask(id domain.TaskID) error {
	if _, exists := r.tasks[string(id)]; !exists {
		return core.ErrTaskNotFound
	}

	delete(r.tasks, string(id))
	return nil
}

// GetTask implements core.Repository.
func (r *Repository) GetTask(id domain.TaskID) (*domain.Task, error) {
	task, exists := r.tasks[string(id)]
	if !exists {
		return nil, core.ErrTaskNotFound
	}

	return task, nil
}

// ListTasks implements core.Repository.
func (r *Repository) ListTasks() ([]*domain.Task, error) {
	tasks := make([]*domain.Task, 0, len(r.tasks))

	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// UpdateTask implements core.Repository.
func (r *Repository) UpdateTask(task *domain.Task) (*domain.Task, error) {
	if _, exists := r.tasks[string(task.ID())]; !exists {
		return nil, core.ErrTaskNotFound
	}

	r.tasks[string(task.ID())] = task
	return task, nil
}
