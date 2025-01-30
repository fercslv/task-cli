package cmd

import (
	"fercslv/task-cli/struts"
	"fmt"
	"github.com/spf13/cobra"
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

		task := struts.Task{
			Name: taskName,
		}
		fmt.Println("Adding task: " + taskName)
	},
}
