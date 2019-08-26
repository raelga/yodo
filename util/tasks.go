package util

import (
	"fmt"
	"os"
	"path/filepath"
)

type task struct {
	id          int
	status      bool
	description string
}

type list struct {
	id    int
	name  string
	tasks []task
}

var defaultList = list{
	id:    0,
	name:  "default",
	tasks: nil,
}

// AddTask adds a task to task list
func AddTask() {
	fmt.Println("add task to file")
}

// RemoveTask deletes a tast from the task list
func RemoveTask() {
	fmt.Println("del task from file")
}

// GetTasks get task list
func GetTasks() {
	fmt.Println("get tasks from file")
}

// createTaskFile
func createTaskFile(taskFile string) {

	if err := os.MkdirAll(filepath.Dir(taskFile), 0755); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if defaultCfgFile, err := os.Create(taskFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		defaultCfgFile.Close()
	}

}
