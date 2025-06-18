package domain

type TaskSpec struct {
	title       string
	description string
}

func NewTaskSpec(title, description string) *TaskSpec {
	return &TaskSpec{
		title:       title,
		description: description,
	}
}

func (s *TaskSpec) Title() string {
	return s.title
}

func (s *TaskSpec) Description() string {
	return s.description
}

type TaskID string

type Task struct {
	id          TaskID
	title       string
	description string
}

func NewTask(id TaskID, title, description string) *Task {
	return &Task{
		id:          id,
		title:       title,
		description: description,
	}
}

func (t *Task) ID() TaskID {
	return t.id
}

func (t *Task) Title() string {
	return t.title
}

func (t *Task) Description() string {
	return t.description
}

func (t *Task) Clone() *Task {
	return &Task{
		id:          t.id,
		title:       t.title,
		description: t.description,
	}
}

func (t *Task) SetSpec(spec *TaskSpec) *Task {
	ret := t.Clone()
	ret.title = spec.title
	ret.description = spec.description

	return ret
}
