package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Print the version number of Task CLI",
	Long:    `As all software does have versions, here you will find the current version of Task CLI.`,
	Example: "task-cli version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Task CLI version v0.0")
	},
}
