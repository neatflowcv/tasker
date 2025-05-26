package domain

import "time"

type TaskID string

type TaskSpec struct {
	title   string
	content string
}

func NewTaskSpec(title, content string) *TaskSpec {
	return &TaskSpec{
		title:   title,
		content: content,
	}
}

type Task struct {
	id         TaskID
	spec       *TaskSpec
	createdAt  time.Time
	archivedAt *time.Time // 삭제된 경우에만 설정
}

func NewTask(id TaskID, spec *TaskSpec, createdAt time.Time, archivedAt *time.Time) *Task {
	return &Task{
		id:         id,
		spec:       spec,
		createdAt:  createdAt,
		archivedAt: archivedAt,
	}
}
