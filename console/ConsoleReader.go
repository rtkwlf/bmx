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

// ConsoleReader is an interface for receiving user input.
type ConsoleReader interface {
	ReadLine(prompt string) (string, error)
	ReadPassword(prompt string) (string, error)
	ReadInt(prompt string) (int, error)
	Option(header string, prompt string, options []string) (int, error)
	Print(message string) error
	Println(message string) error
}

// DefaultConsoleReader is a console interface for emitting and receiving output.
type DefaultConsoleReader struct {
	Tty bool
}

// IsTtyAvailable attempts to open a tty connection, returning true if successful.
//  console.IsTtyAvailable()  			// Check if tty can be opened
//  result := console.IsTtyAvailable()  // Capture if tty is available
func IsTtyAvailable() bool {
	tty, err := openTty()
	if err != nil {
		return false
	}
	defer tty.Close()
	return true
}

// NewConsoleReader creates a that reads from the console.
//  reader := consolerw.NewConsoleReader(false)        // Write output to stderr
//  reader := consolerw.NewConsoleReader(true)        // Write output to tty device
//
// Parameters:
//
//  tty bool      // Required - True if output should be written to tty; false otherwise.
func NewConsoleReader(tty bool) DefaultConsoleReader {
	console := DefaultConsoleReader{
		Tty: tty,
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

// Print writes the message to the output channel.
//  err := consolerw.Print("List of Items")        // Display text to the console.
//
// Parameters:
//
//  message string      // Required - The message to be written to console.
func (r DefaultConsoleReader) Print(message string) error {
	if r.Tty {
		tty, err := openTty()
		if err != nil {
			log.Fatalf("Cannot open tty port: %v\n", err)
		}
		defer tty.Close()
		fmt.Fprint(tty, message)
	} else {
		fmt.Fprint(os.Stderr, message)
	}
	return nil
}

// Println writes the message to the output channel with a newline.
//  err := consolerw.Println("List of Items")        // Display text to the console.
//
// Parameters:
//
//  message string      // Required - The message to be written to console.
func (r DefaultConsoleReader) Println(message string) error {
	if r.Tty {
		tty, err := openTty()
		if err != nil {
			log.Fatalf("Cannot open tty port: %v\n", err)
		}
		defer tty.Close()
		fmt.Fprintln(tty, message)
	} else {
		fmt.Fprintln(os.Stderr, message)
	}
	return nil
}

// ReadLine prompts the console for input with a prompt.
//  text, err := consolerw.ReadLine("Selection:")        // Prompt for selection
//
// Parameters:
//
//  prompt string      // Required - The prompt for input.
func (r DefaultConsoleReader) ReadLine(prompt string) (string, error) {
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

// ReadInt prompts the console for an integer input with a prompt.
//  index, err := consolerw.ReadInt("Selection:")        // Prompt for selection
//
// Parameters:
//
//  prompt string      // Required - The prompt for integer.
func (r DefaultConsoleReader) ReadInt(prompt string) (int, error) {
	var s string
	var err error
	if s, err = r.ReadLine(prompt); err != nil {
		return -1, err
	}

	if s == "" {
		return -1, fmt.Errorf("Input is an empty string, and not valid")
	}

	var i int
	if i, err = strconv.Atoi(s); err != nil {
		return -1, err
	}

	return i, nil
}

// ReadPassword prompts the console for a password input with a prompt.
//  passw, err := consolerw.ReadPassword("Password:")        // Prompt for password
//
// Parameters:
//
//  prompt string      // Required - The prompt for input.
func (r DefaultConsoleReader) ReadPassword(prompt string) (string, error) {
	r.Print(prompt)
	pass, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	r.Println("")
	return string(pass[:]), nil
}

// Option displays a list of options, prompting the console for a selection.
//  consolerw.Option("Message text", "Action", []string{"hello", "world"})  // Display a list of hello and world, zero indexed for selection.
//
// Parameters:
//
//  header string     // Required - The header of the list selection
//  prompt string      // Required - The prompt for input
//  options string     // Required - The options given in the list
func (r DefaultConsoleReader) Option(header string, prompt string, options []string) (int, error) {
	if len(options) == 0 {
		return -1, fmt.Errorf("No options available for selection")
	}

	if len(options) == 1 {
		return 0, nil
	}

	r.Println(header)
	for idx, option := range options {
		r.Println(fmt.Sprintf("[%d] %s", idx, option))
	}
	selection, err := r.ReadInt(prompt)
	if err != nil {
		return -1, err
	}
	return selection, nil
}
