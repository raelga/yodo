// Package util > tasks
/*
Copyright © 2019 Rael Garcia <rael@rael.io>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
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
	task := task{ID: len(taskList.Tasks), Status: status, Description: description}
	taskList.Tasks = append(taskList.Tasks, task)
	return fmt.Sprintf("add task \"%s\" (%t) to file", task.Description, task.Status)
}

// DoTask deletes a task from the task list
func DoTask(taskID int) string {
	if taskList.Tasks != nil {
		for i, task := range taskList.Tasks {
			if task.ID == taskID {
				taskList.Tasks[i].Status = true
				return fmt.Sprintf("Task %d done.\n", task.ID)
			}
		}
	}
	return fmt.Sprintf("Task %d not found.\n", taskID)
}

// RemoveTask deletes a task from the task list
func RemoveTask(taskID int) string {
	if taskList.Tasks != nil {
		for i, task := range taskList.Tasks {
			if task.ID == taskID {
				taskList.Tasks = append(taskList.Tasks[:i], taskList.Tasks[i+1:]...)
				return fmt.Sprintf("Task %d removed.\n", task.ID)
			}
		}
	}
	return fmt.Sprintf("Task %d not found.\n", taskID)
}

// GetTasks get task list
func GetTasks() string {
	out := fmt.Sprintf("%s:\n", taskList.Name)
	if len(taskList.Tasks) > 0 {
		for i, task := range taskList.Tasks {
			out += fmt.Sprintf("- [%d] \"%s\" (%t)\n", i, task.Description, task.Status)
		}
	} else {
		out += "Empty list\n"
	}
	return out
}
