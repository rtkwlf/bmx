//+build !darwin

package main

import (
	"log"

	"github.com/rtkwlf/bmx/config"
	"github.com/rtkwlf/bmx/console"
)

func selectConsoleReader(userConfig config.UserConfig, checkTty bool) console.ConsoleReader {
	if checkTty {
		if console.IsTtyAvailable() {
			return console.NewConsoleReader(true)
		}
		log.Fatal("Cannot create tty connection for writing output")
	}
	return console.NewConsoleReader(false)
}
