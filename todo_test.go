package todo

import (
	"fmt"
	"testing"
)

func TestAddTodo(t *testing.T) {
	todo := Todo{Tasks: []*Task{}, Current: 0}
	todo.AddTask("first task")
	todo.AddTask("second task")

	for todo.Next() {
		fmt.Println(todo.Get())
	}

	todo.UpdateTask(1, "AAAA")
	todo.SetTaskStatus(1)

	for todo.Next() {
		fmt.Println(todo.Get())
	}

	t.Error("finished")
}
