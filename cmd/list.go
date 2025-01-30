package cmd

import (
	"encoding/json"
	"fercslv/task-cli/structs"

	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"time"
)

func init() {
	rootCmd.AddCommand(listCommand)
}

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long: `List all tasks added. You can also filter the list by status 
			(todo, in-progress, done) adding the desired status after the command.`,
	Run: func(cmd *cobra.Command, args []string) {
		status := ""
		if len(args) > 0 {
			status = args[0]
		}

		if status != "todo" && status != "in-progress" && status != "done" && status != "" {
			fmt.Println("Invalid status for filtering.")
			return
		}

		tasks, _ := ioutil.ReadFile("tasks.json")
		var data []structs.Task
		err := json.Unmarshal(tasks, &data)
		if err != nil {
			fmt.Println(err)
			return
		}

		newData := make([]table.Row, 0)
		for _, task := range data {
			if status == "" || (status != "" && task.Status == status) {
				newData = append(newData, table.Row{task.Id, task.Name, task.Status, time.UnixMicro(task.CreatedAt).Format(time.DateTime), time.UnixMicro(task.UpdatedAt).Format(time.DateTime)})
			}
		}

		if len(newData) == 0 {
			if status != "" {
				fmt.Println("No tasks found.")
			} else {
				fmt.Printf("No tasks found with status %s\n", status)
			}
			return
		}

		if status != "" {
			fmt.Printf("Found %v tasks with status %s\n", len(newData), status)
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Id", "Task Name", "Status", "Created At", "Updated At"})
		t.AppendRows(newData)
		t.SetStyle(table.StyleLight)
		t.Render()
	},
}
