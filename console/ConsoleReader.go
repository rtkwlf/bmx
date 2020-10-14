package console

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

type ConsoleReader interface {
	ReadLine(prompt string) (string, error)
	ReadPassword(prompt string) (string, error)
	ReadInt(prompt string) (int, error)
	Print(prompt string) error
	Println(prompt string) error
}

type DefaultConsoleReader struct {
	Tty bool
}

func NewConsoleReader() *DefaultConsoleReader {
	console := &DefaultConsoleReader{
		Tty: false,
	}
	return console
}

func openTty() (*os.File, error) {
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR|syscall.O_NOCTTY, 0)
	if err != nil {
		return tty, err
	}
	return tty, nil
}

func (r *DefaultConsoleReader) Print(prompt string) error {
	if r.Tty {
		tty, err := openTty()
		if err != nil {
			log.Fatalf("Cannot open tty port: %v\n", err)
		}
		defer tty.Close()
		fmt.Fprint(tty, prompt)
	} else {
		fmt.Fprint(os.Stderr, prompt)
	}
	return nil
}
func (r *DefaultConsoleReader) Println(prompt string) error {
	if r.Tty {
		tty, err := openTty()
		if err != nil {
			log.Fatalf("Cannot open tty port: %v\n", err)
		}
		defer tty.Close()
		fmt.Fprintln(tty, prompt)
	} else {
		fmt.Fprintln(os.Stderr, prompt)
	}
	return nil
}

func (r *DefaultConsoleReader) ReadLine(prompt string) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	r.Print(prompt)

	var s string
	scanner.Scan()
	if scanner.Err() != nil {
		return "", scanner.Err()
	}
	s = scanner.Text()
	return s, nil
}

func (r *DefaultConsoleReader) ReadInt(prompt string) (int, error) {
	var s string
	var err error
	if s, err = r.ReadLine(prompt); err != nil {
		return -1, err
	}

	var i int
	if i, err = strconv.Atoi(s); err != nil {
		return -1, err
	}

	return i, nil
}

func (r *DefaultConsoleReader) ReadPassword(prompt string) (string, error) {
	r.Print(prompt)
	pass, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	return string(pass[:]), nil
}
