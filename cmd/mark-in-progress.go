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
	"time"
)

func init() {
	rootCmd.AddCommand(markInProgress)
}

var markInProgress = &cobra.Command{
	Use:     "mark-in-progress",
	Short:   "Sets a task as in progress",
	Long:    `Sets a task as in progress with the argument id provided.`,
	Example: "task-cli mark-in-progress [id]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Invalid arguments. It's required the task id and the new name to update the task.")
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

		data[idx].Status = "in-progress"
		data[idx].UpdatedAt = time.Now().UnixMicro()

		jsonData, _ := json.Marshal(data)
		ioutil.WriteFile("tasks.json", jsonData, os.ModePerm)
		fmt.Printf("Task updated successfully (%v)\n", data[idx].Id)
	},
}
