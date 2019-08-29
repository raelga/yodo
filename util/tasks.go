package util

import (
	"fmt"
)

type task struct {
	ID          int
	Status      bool
	Description string
}

type list struct {
	ID    int
	Name  string
	Tasks []task
}

var defaultList = list{
	ID:    0,
	Name:  "default",
	Tasks: make([]task, 0),
}

var taskList = list{}

// AddTask adds a task to task list
func AddTask(status bool, description string) string {
	task := task{ID: 0, Status: status, Description: description}
	defaultList.Tasks = append(defaultList.Tasks, task)
	return fmt.Sprintf("add task \"%s\" (%t) to file", task.Description, task.Status)
}

// DoTask deletes a tast from the task list
func DoTask(taskID int) string {
	if defaultList.Tasks != nil {
		for i, task := range defaultList.Tasks {
			if task.ID == taskID {
				defaultList.Tasks[i].Status = true
				return fmt.Sprintf("Task %d done.\n", task.ID)
			}
		}
	}
	return fmt.Sprintf("Task %d not found.\n", taskID)
}

// RemoveTask deletes a tast from the task list
func RemoveTask(taskID int) string {
	if defaultList.Tasks != nil {
		for i, task := range defaultList.Tasks {
			if task.ID == taskID {
				defaultList.Tasks = append(defaultList.Tasks[:i], defaultList.Tasks[i+1:]...)
				return fmt.Sprintf("Task %d removed.\n", task.ID)
			}
		}
	}
	return fmt.Sprintf("Task %d not found.\n", taskID)
}

// GetTasks get task list
func GetTasks() string {
	out := fmt.Sprintf("%s:\n", defaultList.Name)
	if len(defaultList.Tasks) > 0 {
		for i, task := range defaultList.Tasks {
			out += fmt.Sprintf("- [%d] \"%s\" (%t)\n", i, task.Description, task.Status)
		}
	} else {
		out += "Empty list\n"
	}
	return out
}
