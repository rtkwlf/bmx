package console

type AppleScriptConsole struct {
}

func NewAppleScriptReader() *AppleScriptConsole {
	console := &AppleScriptConsole{}
	return console
}

func (r *DefaultConsoleReader) Print(prompt string) error {
	return nil
}
func (r *DefaultConsoleReader) Println(prompt string) error {
	return nil
}

func (r *DefaultConsoleReader) ReadLine(prompt string) (string, error) {
	return "", nil
}

func (r *DefaultConsoleReader) ReadInt(prompt string) (int, error) {
	return 0, nil
}

func (r *DefaultConsoleReader) ReadPassword(prompt string) (string, error) {
	return "", nil
}

func (r *DefaultConsoleReader) Option(message string, prompt string, options []string) (int, error) {
	return 0, nil
}
