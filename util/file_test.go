// Package util > file_test
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
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/kr/pretty"
	"gopkg.in/yaml.v2"
)

func Test_createTaskFile(t *testing.T) {
	type args struct {
		listFilePath string
		list         list
	}

	var emptyList = list{
		ID:    0,
		Name:  "default",
		Tasks: make([]task, 0),
	}

	var task1List = list{
		ID:    0,
		Name:  "default",
		Tasks: []task{task1},
	}

	tempFile, err := ioutil.TempFile("/tmp", "test-list.*.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "defaultListValidPath", args: args{listFilePath: tempFile.Name(), list: defaultList}, wantErr: false},
		{name: "task1ListValidPath", args: args{listFilePath: tempFile.Name(), list: task1List}, wantErr: false},
		{name: "empyListValidPath", args: args{listFilePath: tempFile.Name(), list: emptyList}, wantErr: false},
		{name: "newListValidPath", args: args{listFilePath: tempFile.Name(), list: list{}}, wantErr: true},
		{name: "defaultListBadPath", args: args{listFilePath: "/fake-folder/dep.yaml", list: defaultList}, wantErr: true},
		{name: "defaultListFolderPath", args: args{listFilePath: "/fake-folder/", list: defaultList}, wantErr: true},
		{name: "emptyFilePath", args: args{listFilePath: "", list: defaultList}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run("errorHandling/"+tt.name, func(t *testing.T) {
			if err := createTaskFile(tt.args.listFilePath, tt.args.list); (err != nil) != tt.wantErr {
				t.Errorf("createTaskFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		t.Run("savedState/"+tt.name, func(t *testing.T) {
			if err := createTaskFile(tt.args.listFilePath, tt.args.list); err != nil {
				return
			}

			bytes, err := ioutil.ReadFile(tt.args.listFilePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("createTaskFile() read error = %v, wantErr %v", err, tt.wantErr)
			}

			savedList := list{}
			if err := yaml.Unmarshal(bytes, &savedList); (err != nil) != tt.wantErr {
				t.Errorf("createTaskFile() unmarshal error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.args.list, savedList) {
				t.Errorf("createTaskFile() drift error: %v", pretty.Diff(tt.args.list, savedList))
			}

		})
	}

}

func Test_saveTasksFile(t *testing.T) {
	type args struct {
		listFilePath string
		list         list
	}

	var emptyList = list{
		ID:    0,
		Name:  "default",
		Tasks: make([]task, 0),
	}

	var task1List = list{
		ID:    0,
		Name:  "default",
		Tasks: []task{task1},
	}

	tempFile, err := ioutil.TempFile("/tmp", "test-list.*.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "defaultListValidPath", args: args{listFilePath: tempFile.Name(), list: defaultList}, wantErr: false},
		{name: "task1ListValidPath", args: args{listFilePath: tempFile.Name(), list: task1List}, wantErr: false},
		{name: "empyListValidPath", args: args{listFilePath: tempFile.Name(), list: emptyList}, wantErr: false},
		{name: "defaultListBadPath", args: args{listFilePath: "/fake-folder/dep.yaml", list: defaultList}, wantErr: true},
		{name: "defaultListFolderPath", args: args{listFilePath: "/fake-folder/", list: defaultList}, wantErr: true},
		{name: "emptyFilePath", args: args{listFilePath: "", list: defaultList}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run("errorHandling/"+tt.name, func(t *testing.T) {
			if err := saveTasksFile(tt.args.listFilePath, tt.args.list); (err != nil) != tt.wantErr {
				t.Errorf("saveTasksFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		t.Run("savedState/"+tt.name, func(t *testing.T) {
			if err := saveTasksFile(tt.args.listFilePath, tt.args.list); err != nil {
				return
			}

			bytes, err := ioutil.ReadFile(tt.args.listFilePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("saveTasksFile() read error = %v, wantErr %v", err, tt.wantErr)
			}

			savedList := list{}
			if err := yaml.Unmarshal(bytes, &savedList); (err != nil) != tt.wantErr {
				t.Errorf("saveTasksFile() unmarshal error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.args.list, savedList) {
				t.Errorf("saveTasksFile() drift error: %v", pretty.Diff(tt.args.list, savedList))
			}

		})
	}

}
