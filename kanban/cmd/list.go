package cmd

import (
	"fmt"
	"learn-go/kanban/db"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong: ", err.Error())
			os.Exit(1) // exit with error code 1 - something went wrong.
		}
		if len(tasks) == 0 {
			fmt.Println("You have no pending tasks yet.")
			return
		}
		// list all tasks
		for index, task := range tasks {
			fmt.Printf("%d. %s\n", index+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
