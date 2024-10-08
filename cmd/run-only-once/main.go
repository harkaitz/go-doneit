package main

import (
	"fmt"
	"os"
	"os/exec"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/logger"
	"github.com/harkaitz/go-doneit"
)

var HELP string =
`Usage: run-only-once ID COMMAND

Execute command only once, the database is in "~/.run-only-once.db".

Copyright (C) 2024 Harkaitz Agirre, all rights reserved`

func main() {

	var gdb		*gorm.DB
	var home	string
	var err		error

	defer func() {
		if err != nil {
			fmt.Fprintf(os.Stderr, "run-only-once: error:\n%v\n", err.Error())
			os.Exit(1)
		}
	}()

	if len(os.Args) < 3 {
		fmt.Println(HELP)
		return
	}

	home, err = os.UserHomeDir()
	if err != nil { return }

	gdb, err = gorm.Open(sqlite.Open(home + "/.run-only-once.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil { return }

	err = doneit.InitDatabase(gdb)
	if err != nil { return }

	err = doneit.OnlyOnce(gdb, func() (err error) {
		var cmd		*exec.Cmd
		cmd = exec.Command(os.Args[2], os.Args[3:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		err = cmd.Run()
		if err != nil { return }
		return
	}, "%s", os.Args[1])
	if err != nil { return }

	return
}
