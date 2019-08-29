// Package util > file
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
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// parse
func (list *list) Parse(data []byte) error {
	return yaml.Unmarshal(data, list)
}

// LoadTasks loads a list from a file
func LoadTasks(listFilePath string) error {

	if _, err := os.Stat(listFilePath); err != nil {
		if err := createTaskFile(listFilePath, defaultList); err != nil {
			return err
		}
	}

	listFile, err := ioutil.ReadFile(listFilePath)
	if err != nil {
		return err
	}

	if err := taskList.Parse(listFile); err != nil {
		return err
	}

	return nil
}

// SaveTasks saves the file to disk
func SaveTasks(listFilePath string) error {
	return saveTasksFile(listFilePath, taskList)
}

// SaveTasksFile saves the task file to disk
func saveTasksFile(listFilePath string, list list) error {
	bytes, err := yaml.Marshal(list)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(listFilePath, bytes, 0644)
}

// createTaskFile
func createTaskFile(listFilePath string, list list) error {

	if list.Name == "" {
		return errors.New("the list should have a name")
	}

	if err := os.MkdirAll(filepath.Dir(listFilePath), 0755); err != nil {
		return err
	}

	return saveTasksFile(listFilePath, list)

}
