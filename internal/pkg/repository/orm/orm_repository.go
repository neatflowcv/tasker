package orm

import (
	"errors"

	"github.com/neatflowcv/tasker/internal/pkg/domain"
	"github.com/neatflowcv/tasker/internal/pkg/repository/core"
	"github.com/oklog/ulid/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ core.Repository = (*Repository)(nil)

type TaskModel struct {
	ID          string `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
}

func (TaskModel) TableName() string {
	return "tasks"
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(dsn string) *Repository {
	var config gorm.Config

	db, err := gorm.Open(postgres.Open(dsn), &config)
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&TaskModel{}) //nolint:exhaustruct
	if err != nil {
		panic(err)
	}

	return &Repository{db: db}
}

func (r *Repository) CreateTask(spec *domain.TaskSpec) (*domain.Task, error) {
	taskID := ulid.Make()
	taskModel := TaskModel{
		ID:          taskID.String(),
		Title:       spec.Title(),
		Description: spec.Description(),
	}

	err := r.db.Create(&taskModel).Error
	if err != nil {
		return nil, err
	}

	return domain.NewTask(
		domain.TaskID(taskModel.ID),
		taskModel.Title,
		taskModel.Description,
	), nil
}

func (r *Repository) ListTasks() ([]*domain.Task, error) {
	var taskModels []TaskModel
	if err := r.db.Find(&taskModels).Error; err != nil {
		return nil, err
	}

	tasks := make([]*domain.Task, len(taskModels))
	for i, model := range taskModels {
		tasks[i] = domain.NewTask(
			domain.TaskID(model.ID),
			model.Title,
			model.Description,
		)
	}

	return tasks, nil
}

func (r *Repository) GetTask(id domain.TaskID) (*domain.Task, error) {
	var taskModel TaskModel

	err := r.db.First(&taskModel, "id = ?", string(id)).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrTaskNotFound
		}

		return nil, err
	}

	return domain.NewTask(
		domain.TaskID(taskModel.ID),
		taskModel.Title,
		taskModel.Description,
	), nil
}

func (r *Repository) UpdateTask(task *domain.Task) (*domain.Task, error) {
	taskModel := TaskModel{
		ID:          string(task.ID()),
		Title:       task.Title(),
		Description: task.Description(),
	}

	if err := r.db.Save(&taskModel).Error; err != nil {
		return nil, err
	}

	return domain.NewTask(
		domain.TaskID(taskModel.ID),
		taskModel.Title,
		taskModel.Description,
	), nil
}

func (r *Repository) DeleteTask(id domain.TaskID) error {
	result := r.db.Delete(&TaskModel{ //nolint:exhaustruct
		ID: string(id),
	})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return core.ErrTaskNotFound
	}

	return nil
}
