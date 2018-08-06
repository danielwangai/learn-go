package cmd

import (
	"fmt"
	"learn-go/kanban/db"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "To add a new task.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong while creating the task.")
			os.Exit(1)
		}
		fmt.Println("Task created successfully.")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
