package domain

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

func (s *TaskSpec) Title() string {
	return s.title
}

func (s *TaskSpec) Content() string {
	return s.content
}

type TaskID string

type Task struct {
	id      TaskID
	title   string
	content string
}

func NewTask(id TaskID, title, content string) *Task {
	return &Task{
		id:      id,
		title:   title,
		content: content,
	}
}

func (t *Task) ID() TaskID {
	return t.id
}

func (t *Task) Title() string {
	return t.title
}

func (t *Task) Content() string {
	return t.content
}

func (t *Task) Clone() *Task {
	return &Task{
		id:      t.id,
		title:   t.title,
		content: t.content,
	}
}

func (t *Task) SetSpec(spec *TaskSpec) *Task {
	ret := t.Clone()
	ret.title = spec.title
	ret.content = spec.content
	return ret
}
