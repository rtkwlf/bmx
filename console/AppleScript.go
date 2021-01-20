package console

import "github.com/andybrewer/mack"

type AppleScriptConsole struct {
}

func NewAppleScriptReader() *AppleScriptConsole {
	console := &AppleScriptConsole{}
	return console
}

func (r *DefaultConsoleReader) Print(prompt string) error {
	mack.Say("Starting process")
	return nil
}
func (r *DefaultConsoleReader) Println(prompt string) error {
	mack.Say("Starting process")
	return nil
}

func (r *DefaultConsoleReader) ReadLine(prompt string) (string, error) {
	mack.Say("Starting process")
	return "", nil
}

func (r *DefaultConsoleReader) ReadInt(prompt string) (int, error) {
	mack.Say("Starting process")
	return 0, nil
}

func (r *DefaultConsoleReader) ReadPassword(prompt string) (string, error) {
	mack.Say("Starting process")
	return "", nil
}

func (r *DefaultConsoleReader) Option(message string, prompt string, options []string) (int, error) {
	mack.Say("Starting process")
	return 0, nil
}
