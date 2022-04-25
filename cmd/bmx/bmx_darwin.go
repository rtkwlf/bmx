//go:build darwin
// +build darwin

package main

import (
	"github.com/rtkwlf/bmx/config"
	"github.com/rtkwlf/bmx/console"
)

func selectConsoleReader(userConfig config.UserConfig, limited bool) console.ConsoleReader {
	if userConfig.Input == config.AlwaysUseAppleScript {
		return console.NewAppleScriptReader()
	}

	if !limited {
		return console.NewConsoleReader(false)
	}

	if userConfig.Input == config.UseAppleScriptLimited {
		return console.NewAppleScriptReader()
	}

	if console.IsTtyAvailable() {
		return console.NewConsoleReader(true)
	}
	return console.NewAppleScriptReader()
}
