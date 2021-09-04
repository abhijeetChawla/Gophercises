package main

import (
	"fmt"
	"os"
	"path"
	"tasks/cmd"
	"tasks/db"

	homeDir "github.com/mitchellh/go-homedir"
)

func main() {
	dir, err := homeDir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	db.Init(path.Join(dir, "tasks.db"))
	cmd.Execute()
}
