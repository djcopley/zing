package editor

import (
	"os"
	"os/exec"
	"strings"
)

// getUserEditor retrieves the user's configured editor as a slice of strings. Defaults to "nano" if no editor is set.
func getUserEditor() []string {
	editor := os.Getenv("VISUAL")
	if editor == "" {
		editor = os.Getenv("EDITOR")
	}
	if editor == "" {
		editor = "nano"
	}
	return strings.Fields(editor)
}

// Open opens the given file path in the user's configured editor.
// If no editor is configured, it defaults to "nano".
func Open(path string) error {
	editor := getUserEditor()
	editor = append(editor, path)

	cmd := exec.Command(editor[0], editor[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// ComposeMessage opens the user's configured editor and returns a byte slice containing the message. If no editor is
// configured, ComposeMessage defaults to "nano".
func ComposeMessage() (string, error) {
	tmpDir := os.TempDir()
	tmp, err := os.CreateTemp(tmpDir, "zing_*")
	if err != nil {
		return "'", err
	}
	err = tmp.Close()
	if err != nil {
		return "", err
	}
	defer os.Remove(tmp.Name())

	editor := getUserEditor()
	editor = append(editor, tmp.Name())

	cmd := exec.Command(editor[0], editor[1:]...)
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
