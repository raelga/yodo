package util

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

var emptyList = list{
	ID:    0,
	Name:  "default",
	Tasks: make([]task, 0),
}

var taskMissing = task{ID: -1, Status: false, Description: "Missing task"}
var task1 = task{ID: 0, Status: false, Description: "task1"}

var task1List = list{
	ID:    0,
	Name:  "default",
	Tasks: []task{task1},
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
		got := AddTask(task1.Status, task1.Description)
		want := fmt.Sprintf("add task \"%s\" (%t) to file", task1.Description, task1.Status)
		assertCorrectMessage(t, got, want)
	})

	t.Run("AddedTaskIsAdded", func(t *testing.T) {
		taskList = emptyList
		_ = AddTask(task1.Status, task1.Description)
		got := taskList
		want := task1List
		assertCorrectList(t, got, want)
	})

}

func TestDoTask(t *testing.T) {

	t.Run("DoTaskMessageIsPrinted", func(t *testing.T) {
		taskList = task1List
		got := DoTask(0)
		want := fmt.Sprintf("Task %d done.\n", task1.ID)
		assertCorrectMessage(t, got, want)
	})

	t.Run("DoTaskFromEmptyListMessageIsPrinted", func(t *testing.T) {
		taskList = emptyList
		got := DoTask(0)
		want := fmt.Sprintf("Task %d not found.\n", task1.ID)
		assertCorrectMessage(t, got, want)
	})

	t.Run("DoTaskNotFoundListMessageIsPrinted", func(t *testing.T) {
		taskList = task1List
		got := DoTask(taskMissing.ID)
		want := fmt.Sprintf("Task %d not found.\n", taskMissing.ID)
		assertCorrectMessage(t, got, want)
	})

	t.Run("DoFromEmptyListIsEmpty", func(t *testing.T) {
		taskList = emptyList
		_ = DoTask(0)
		got := taskList
		want := emptyList
		assertCorrectList(t, got, want)
	})
}

func TestRemoveTask(t *testing.T) {

	t.Run("RemoveTaskMessageIsPrinted", func(t *testing.T) {
		taskList = task1List
		got := RemoveTask(0)
		want := fmt.Sprintf("Task %d removed.\n", task1.ID)
		assertCorrectMessage(t, got, want)
	})

	t.Run("RemoveTaskRemovesTheTask", func(t *testing.T) {
		taskList = task1List
		_ = RemoveTask(taskList.Tasks[0].ID)
		got := taskList
		want := emptyList
		assertCorrectList(t, got, want)
	})

	t.Run("RemoveTaskFromEmptyListMessageIsPrinted", func(t *testing.T) {
		taskList = emptyList
		got := RemoveTask(0)
		want := fmt.Sprintf("Task %d not found.\n", task1.ID)
		assertCorrectMessage(t, got, want)
	})

	t.Run("RemoveTaskNotFoundListMessageIsPrinted", func(t *testing.T) {
		taskList = task1List
		got := RemoveTask(taskMissing.ID)
		want := fmt.Sprintf("Task %d not found.\n", taskMissing.ID)
		assertCorrectMessage(t, got, want)
	})

	t.Run("RemoveFromEmptyListIsEmpty", func(t *testing.T) {
		taskList = emptyList
		_ = RemoveTask(0)
		got := taskList
		want := emptyList
		assertCorrectList(t, got, want)
	})
}

func TestGetTasks(t *testing.T) {

	t.Run("EmptyListMessageIsPrinted", func(t *testing.T) {
		taskList = emptyList
		got := GetTasks()
		want := "default:\nEmpty list\n"
		assertCorrectMessage(t, got, want)
	})

	t.Run("EmptyListIsEmpty", func(t *testing.T) {
		taskList = emptyList
		_ = GetTasks()
		got := taskList
		want := emptyList
		assertCorrectList(t, got, want)
	})
}
