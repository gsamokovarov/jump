package shell

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

var (
	Fish = Shell{"fish", "jump.fish"}
	Bash = Shell{"bash", "jump.bash"}
	Zsh  = Shell{"zsh", "jump.zsh"}
)

type Shell struct {
	Name   string
	Script string
}

// Integration return the integration script contents for that specific shell.
func (s Shell) Integration() (string, error) {
	scriptPath := filepath.Join(dir(), s.Script)

	bytes, err := ioutil.ReadFile(scriptPath)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// Guess tries to guess the current shell session.
func Guess() Shell {
	shell := filepath.Base(os.Getenv("SHELL"))

	switch shell {
	case "fish":
		return Fish
	case "zsh":
		return Zsh
	default:
		return Bash
	}
}

func dir() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Dir(file)
}
