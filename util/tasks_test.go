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
	tasks: nil,
}

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
