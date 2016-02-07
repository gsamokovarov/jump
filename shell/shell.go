package shell

import "path/filepath"

// Shell is a small wrapper around string that can pretty print the shell
// integration scripts.
type Shell string

// Guess tries to guess the current shell session.
func Guess(hint string) Shell {
	shell := filepath.Base(hint)

	switch shell {
	case "fish":
		return Fish
	case "zsh":
		return Zsh
	default:
		return Bash
	}
}
