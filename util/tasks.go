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
func AddTask(status bool, description string) string {
	task := task{id: 0, status: status, description: description}
	defaultList.tasks = append(defaultList.tasks, task)
	return fmt.Sprintf("add task \"%s\" (%t) to file", task.description, task.status)
}

// DoTask deletes a tast from the task list
func DoTask(taskID int) string {
	if defaultList.tasks != nil {
		for i, task := range defaultList.tasks {
			if task.id == taskID {
				defaultList.tasks[i].status = true
				return fmt.Sprintf("Task %d done.\n", task.id)
			}
		}
	}
	return fmt.Sprintf("Task %d not found.\n", taskID)
}

// RemoveTask deletes a tast from the task list
func RemoveTask(taskID int) string {
	if defaultList.tasks != nil {
		for i, task := range defaultList.tasks {
			if task.id == taskID {
				defaultList.tasks = append(defaultList.tasks[:i], defaultList.tasks[i+1:]...)
				return fmt.Sprintf("Task %d removed.\n", task.id)
			}
		}
	}
	return fmt.Sprintf("Task %d not found.\n", taskID)
}

// GetTasks get task list
func GetTasks() string {
	out := fmt.Sprintf("%s:\n", defaultList.name)
	if len(defaultList.tasks) > 0 {
		for i, task := range defaultList.tasks {
			out += fmt.Sprintf("- [%d] \"%s\" (%t)\n", i, task.description, task.status)
		}
	} else {
		out += "Empty list\n"
	}
	return out
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
