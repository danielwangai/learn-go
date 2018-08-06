package main

import (
	"fmt"
	"learn-go/kanban/cmd"
	"learn-go/kanban/db"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "task.db")
	errorHandler(db.Init(dbPath))
	errorHandler(cmd.RootCmd.Execute())
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
