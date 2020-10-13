// +build windows

package console

import (
	"bufio"
	"os"
)

func NewScanner() *bufio.Scanner {
	return bufio.NewScanner(os.Stdin)
}

func NewPromptWriter() *bufio.Writer {
	return bufio.NewWriter(os.Stderr)
}
