// +build darwin linux

package console

import (
	"bufio"
	"log"
	"os"
)

func NewScanner() *bufio.Scanner {
	tty, err := os.Open("/dev/tty")
	if err != nil {
		log.Fatalf("can't open /dev/tty: %s", err)
	}
	return bufio.NewScanner(tty)
}

func NewPromptWriter() *bufio.Writer {
	return bufio.NewWriter(os.Stdin)
}
