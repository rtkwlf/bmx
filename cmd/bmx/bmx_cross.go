//+build !darwin

package main

import (
	"log"

	"github.com/rtkwlf/bmx/config"
	"github.com/rtkwlf/bmx/console"
)

func getInputReader(userConfig config.UserConfig, checkTty bool) console.ConsoleReader {
	if checkTty {
		if console.IsTtyAvailable() {
			return *console.NewConsoleReader(true)
		}
		log.Fatal("Cannot instantiate a mechanism for input")
	}
	return *console.NewConsoleReader(false)
}
