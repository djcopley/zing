package editor

import (
	"os"
	"os/exec"
)

func getUserEditor() string {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}
	return editor
}

// ComposeMessage opens the user's configured editor and returns a byte slice containing the message. If no editor is
// configured, ComposeMessage defaults to "nano".
func ComposeMessage() (string, error) {
	tmp, err := os.CreateTemp("/tmp", "zing_*") // todo: can I customize the * tmp id?
	if err != nil {
		return "'", err
	}
	err = tmp.Close()
	if err != nil {
		return "", err
	}
	defer os.Remove(tmp.Name())

	cmd := exec.Command(getUserEditor(), tmp.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	resultBytes, err := os.ReadFile(tmp.Name())
	if err != nil {
		return "", err
	}
	result := string(resultBytes)

	return result, nil
}
