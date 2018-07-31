package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "kanban",
	Short: "Task is a CLI task manager",
}
