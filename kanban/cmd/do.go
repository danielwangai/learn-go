package cmd

import (
	"fmt"
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
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
