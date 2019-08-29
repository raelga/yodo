package util

import (
	"errors"
	"fmt"
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
	fmt.Printf("Saving %v %s\n", list, listFilePath)
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
