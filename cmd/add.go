package cmd

import (
	"encoding/json"
	"fercslv/task-cli/structs"

	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"time"
)

func init() {
	rootCmd.AddCommand(addCommand)
}

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a new task with a name",
	Long:  `Add a new task to the list. The default status is 'todo'`,
	Run: func(cmd *cobra.Command, args []string) {
		taskName := args[0]

		tasks, _ := ioutil.ReadFile("tasks.json")
		var data []structs.Task
		json.Unmarshal(tasks, &data)
		lastItem := len(data) - 1
		id := 1
		if lastItem > -1 {
			id = data[lastItem].Id + 1
		}

		task := structs.Task{
			Id:        id,
			Name:      taskName,
			Status:    "todo",
			CreatedAt: time.Now().UnixMicro(),
			UpdatedAt: time.Now().UnixMicro(),
		}

		data = append(data, task)
		jsonData, _ := json.Marshal(data)
		ioutil.WriteFile("tasks.json", jsonData, os.ModePerm)
		fmt.Printf("Task added successfully (%v)\n", task.Id)
	},
}
