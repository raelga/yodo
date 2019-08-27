package util

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

var emptyList = list{
	id:    0,
	name:  "default",
	tasks: make([]task, 0),
}

var taskMissing = task{id: -1, status: false, description: "Missing task"}
var task1 = task{id: 0, status: false, description: "task1"}

var task1List = list{
	id:    0,
	name:  "default",
	tasks: []task{task1},
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func assertCorrectList(t *testing.T, got, want list) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(fmt.Sprintf("%s", pretty.Diff(got, want)))
	}
}

func TestAddTask(t *testing.T) {

	t.Run("AddMessageIsPrinted", func(t *testing.T) {
		got := AddTask(task1.status, task1.description)
		want := fmt.Sprintf("add task \"%s\" (%t) to file", task1.description, task1.status)
		assertCorrectMessage(t, got, want)
	})

	t.Run("AddedTaskIsAdded", func(t *testing.T) {
		defaultList = emptyList
		_ = AddTask(task1.status, task1.description)
		got := defaultList
		want := task1List
		assertCorrectList(t, got, want)
	})

}

func TestDoTask(t *testing.T) {

	t.Run("DoTaskMessageIsPrinted", func(t *testing.T) {
		defaultList = task1List
		got := DoTask(0)
		want := fmt.Sprintf("Task %d done.\n", task1.id)
		assertCorrectMessage(t, got, want)
	})

	t.Run("DoTaskFromEmptyListMessageIsPrinted", func(t *testing.T) {
		defaultList = emptyList
		got := DoTask(0)
		want := fmt.Sprintf("Task %d not found.\n", task1.id)
		assertCorrectMessage(t, got, want)
	})

	t.Run("DoTaskNotFoundListMessageIsPrinted", func(t *testing.T) {
		defaultList = task1List
		got := DoTask(taskMissing.id)
		want := fmt.Sprintf("Task %d not found.\n", taskMissing.id)
		assertCorrectMessage(t, got, want)
	})

	t.Run("DoFromEmptyListIsEmpty", func(t *testing.T) {
		defaultList = emptyList
		_ = DoTask(0)
		got := defaultList
		want := emptyList
		assertCorrectList(t, got, want)
	})
}

func TestRemoveTask(t *testing.T) {

	t.Run("RemoveTaskMessageIsPrinted", func(t *testing.T) {
		defaultList = task1List
		got := RemoveTask(0)
		want := fmt.Sprintf("Task %d removed.\n", task1.id)
		assertCorrectMessage(t, got, want)
	})

	t.Run("RemoveTaskFromEmptyListMessageIsPrinted", func(t *testing.T) {
		defaultList = emptyList
		got := RemoveTask(0)
		want := fmt.Sprintf("Task %d not found.\n", task1.id)
		assertCorrectMessage(t, got, want)
	})

	t.Run("RemoveTaskNotFoundListMessageIsPrinted", func(t *testing.T) {
		defaultList = task1List
		got := RemoveTask(taskMissing.id)
		want := fmt.Sprintf("Task %d not found.\n", taskMissing.id)
		assertCorrectMessage(t, got, want)
	})

	t.Run("RemoveFromEmptyListIsEmpty", func(t *testing.T) {
		defaultList = emptyList
		_ = RemoveTask(0)
		got := defaultList
		want := emptyList
		assertCorrectList(t, got, want)
	})
}

func TestGetTasks(t *testing.T) {

	t.Run("EmptyListMessageIsPrinted", func(t *testing.T) {
		defaultList = emptyList
		got := GetTasks()
		want := "default:\nEmpty list\n"
		assertCorrectMessage(t, got, want)
	})

	t.Run("EmptyListIsEmpty", func(t *testing.T) {
		defaultList = emptyList
		_ = GetTasks()
		got := defaultList
		want := emptyList
		assertCorrectList(t, got, want)
	})
}
