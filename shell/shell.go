package shell

import (
	"os"
	"path/filepath"
)

type Shell string

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
