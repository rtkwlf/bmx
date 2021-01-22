//+build !darwin

package main

import (
	"log"

	"github.com/rtkwlf/bmx/config"
	"github.com/rtkwlf/bmx/console"
)

func selectConsoleReader(userConfig config.UserConfig, limited bool) console.ConsoleReader {
	if !limited {
		return console.NewConsoleReader(false)
	}

	if !console.IsTtyAvailable() {
		log.Fatal("Cannot create tty connection for writing output")
	}

	return console.NewConsoleReader(true)
}
