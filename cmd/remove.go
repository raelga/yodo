// Package cmd > remove
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
package cmd

import (
	"fmt"
	"strconv"

	"github.com/raelga/yodo/util"

	"github.com/spf13/cobra"
)

// removeCmd represents the add command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a task from the list",
	Long:  `Removes a task from the data`,
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			panic(fmt.Sprintf("Unable to parse task %s: %s", args[0], err))
		}
		fmt.Printf(util.RemoveTask(taskID))
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

}
