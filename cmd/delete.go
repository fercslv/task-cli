package cmd

import (
	"encoding/json"
	"fercslv/task-cli/structs"

	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"slices"
	"strconv"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Deletes a task",
	Long:    `Deletes a task with the id provided as argument.`,
	Example: "task-cli delete [id]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Invalid arguments. It's required the task id to delete the task.")
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid id.")
			return
		}

		tasks, _ := ioutil.ReadFile("tasks.json")
		var data []structs.Task
		err = json.Unmarshal(tasks, &data)
		if err != nil {
			fmt.Println(err)
			return
		}

		idx := slices.IndexFunc(data, func(task structs.Task) bool { return task.Id == id })

		if idx == -1 {
			fmt.Println("Task not found.")
			return
		}

		data = slices.Delete(data, idx, idx+1)

		jsonData, _ := json.Marshal(data)
		ioutil.WriteFile("tasks.json", jsonData, os.ModePerm)
		fmt.Println("Task deleted successfully")
	},
}
