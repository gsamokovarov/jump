package shell

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"text/template"
)

// Shell is a small wrapper around string that can pretty print the shell
// integration scripts.
type Shell string

// MustCompile compiles the shell template script to ready to execute (read
// copy and paste) shell integration.
func (s Shell) MustCompile(shortcut string) string {
	var context struct{ Bind string }
	context.Bind = shortcut

	outputBuffer := &bytes.Buffer{}

	tmpl := template.Must(template.New("shell").Parse(string(s)))
	tmpl.Execute(outputBuffer, context)

	outputBytes, err := ioutil.ReadAll(outputBuffer)
	if err != nil {
		panic(err)
	}

	return string(outputBytes)
}

// Guess tries to guess the current shell session.
func Guess(hint string) Shell {
	shell := filepath.Base(hint)

	switch shell {
	case "fish":
		return Fish
	case "zsh":
		return Zsh
	case "pwsh":
		return Pwsh
	default:
		return Bash
	}
}
