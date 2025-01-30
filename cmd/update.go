package cmd

import (
	"encoding/json"
	"fercslv/task-cli/structs"
	"time"

	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"slices"
	"strconv"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Updates a task",
	Long:    `Updates a task with the arguments id and name in the order provided.`,
	Example: "task-cli update [id] [name]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Invalid arguments. It's required the task id and the new name to update the task.")
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid id.")
			return
		}

		name := args[1]
		if name == "" {
			fmt.Println("Invalid name.")
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

		data[idx].Name = name
		data[idx].UpdatedAt = time.Now().UnixMicro()

		jsonData, _ := json.Marshal(data)
		ioutil.WriteFile("tasks.json", jsonData, os.ModePerm)
		fmt.Printf("Task updated successfully (%v)\n", data[idx].Id)
	},
}
