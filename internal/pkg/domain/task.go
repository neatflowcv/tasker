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

func (ts *TaskSpec) Title() string {
	return ts.title
}

func (ts *TaskSpec) Content() string {
	return ts.content
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
