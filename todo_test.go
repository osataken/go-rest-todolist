package todolist

import (
	"testing"
)

func Test_create_todo_item_should_increase_size_of_todo_list_by_1(t *testing.T) {
	toDoList := ToDoList{}
	expected := 1

	toDoList.CreateToDoItem("Topic", "Description")

	result := toDoList.Size()

	if result != expected {
		t.Errorf("Expected %V but was  %V", expected, result)
	}
}