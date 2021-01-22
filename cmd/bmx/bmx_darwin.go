//+build darwin

package main

import (
	"github.com/rtkwlf/bmx/config"
	"github.com/rtkwlf/bmx/console"
)

func selectConsoleReader(userConfig config.UserConfig, checkTty bool) console.ConsoleReader {
	if userConfig.AlwaysUseAppleScript {
		return *console.NewAppleScriptReader()
	}

	if checkTty {
		if console.IsTtyAvailable() {
			return *console.NewConsoleReader(true)
		}
		return *console.NewAppleScriptReader()
	}

	return *console.NewConsoleReader(false)
}
