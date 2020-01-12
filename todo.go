package todo

// Todo struct
type Todo struct {
	Tasks   []*Task
	Current int
}

// Task struct
type Task struct {
	ID      int32
	Content string
	Done    bool
}

// Next get next task
func (t *Todo) Next() bool {
	if t.Current < len(t.Tasks) {
		t.Current++
		return true
	}
	t.Current = 0
	return false
}

// Get Task
func (t *Todo) Get() *Task {
	return t.Tasks[t.Current-1]
}

// AddTask add
func (t *Todo) AddTask(content string) {
	id := int32(len(t.Tasks) + 1)
	task := &Task{ID: id, Content: content, Done: false}
	t.Tasks = append(t.Tasks, task)
}

// UpdateTask update a task
func (t *Todo) UpdateTask(id int32, content string) {
	for index, task := range t.Tasks {
		if task.ID == id {
			t.Tasks[index].Content = content
		}
	}
}

// SetTaskStatus set task staus
func (t *Todo) SetTaskStatus(id int32) {
	for index, task := range t.Tasks {
		if task.ID == id {
			t.Tasks[index].Done = !t.Tasks[id-1].Done
		}
	}
}

// DeleteTask delete a task
func (t *Todo) DeleteTask(id int32) {
	out := []*Task{}
	for _, task := range t.Tasks {
		if task.ID != id {
			out = append(out, task)
		}
	}
	t.Tasks = out
}
