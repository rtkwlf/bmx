package console

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/andybrewer/mack"
)

type AppleScriptConsole struct {
}

func NewAppleScriptReader() AppleScriptConsole {
	return AppleScriptConsole{}
}

func (r AppleScriptConsole) Print(prompt string) error {
	return mack.Notify(prompt)
}
func (r AppleScriptConsole) Println(prompt string) error {
	return mack.Notify(prompt)
}

func (r AppleScriptConsole) ReadLine(prompt string) (string, error) {
	dialog := mack.DialogOptions{
		Text:   prompt,
		Title:  prompt,
		Answer: " ",
	}
	response, err := mack.DialogBox(dialog)
	if err != nil {
		return "", err
	}

	if response.GaveUp {
		return "", fmt.Errorf("No option was selected")
	}
	return response.Text, nil
}

func (r AppleScriptConsole) ReadInt(prompt string) (int, error) {
	dialog := mack.DialogOptions{
		Text:   prompt,
		Title:  prompt,
		Answer: " ",
	}
	response, err := mack.DialogBox(dialog)
	if err != nil {
		return -1, err
	}

	if response.GaveUp {
		return -1, fmt.Errorf("No option was selected")
	}

	i, err := strconv.Atoi(response.Text)
	if err != nil {
		return -1, err
	}

	return i, nil
}

func (r AppleScriptConsole) ReadPassword(prompt string) (string, error) {
	dialog := mack.DialogOptions{
		Text:         prompt,
		Title:        prompt,
		HiddenAnswer: true,
		Answer:       "123",
	}
	response, err := mack.DialogBox(dialog)
	if err != nil {
		return "", err
	}

	if response.GaveUp {
		return "", fmt.Errorf("No option was selected")
	}
	return response.Text, nil
}

// Option displays a desktop dialog box prompting for a selection.
//  consolerw.Option("Message text", "Action", []string{"hello", "world"})  // Display a dialog box with hello and world as options.
//  selection, err := consolerw.Option("My dialog")              			// Capture the selected index for the dialog box
//
// Parameters:
//
//  message string     // Required - The message explaining the choices
//  prompt string      // Required - The prompt for input
//  options string     // Required - The options given in the dialog box
func (r AppleScriptConsole) Option(message string, prompt string, options []string) (int, error) {
	if len(options) == 0 {
		return -1, fmt.Errorf("No options available for selection")
	}

	if len(options) == 1 {
		return 0, nil
	}

	listOptions := mack.ListOptions{
		Items:   options,
		Title:   message,
		Message: prompt,
	}
	response, didCancel, err := mack.ListWithOpts(listOptions)
	if err != nil {
		return -1, err
	}

	if didCancel {
		return -1, fmt.Errorf("No option was selected")
	}

	for idx, app := range options {
		if strings.EqualFold(response[0], app) {
			return idx, nil
		}
	}
	return -1, fmt.Errorf("No option was selected")
}
