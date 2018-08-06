package cmd

import (
	"fmt"
	"learn-go/kanban/db"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Complete a task.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg) // convert string to integer
			if err != nil {
				// filter out non integer arguments
				fmt.Println("Failed to parse argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			// there's an error
			fmt.Println("Something went wrong: ", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number ", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as complete. Error \"%s\"", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as complete.", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
