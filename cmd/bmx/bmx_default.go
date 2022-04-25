//go:build !darwin
// +build !darwin

package main

import (
	"github.com/rtkwlf/bmx/config"
	"github.com/rtkwlf/bmx/console"
)

func selectConsoleReader(userConfig config.UserConfig, limited bool) console.ConsoleReader {
	return console.NewConsoleReader(limited)
}
