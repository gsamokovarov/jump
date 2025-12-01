package shell

import (
	"bytes"
	_ "embed"
	"io"
	"path/filepath"
	"text/template"
)

//go:embed integration.zsh
var zshScript string

//go:embed integration.fish
var fishScript string

//go:embed integration.nu
var nushellScript string

//go:embed integration.ps1
var pwshScript string

//go:embed integration.bash
var bashScript string

// Context holds the context for shell template compilation.
type Context struct {
	Bind string
}

// Shell is a small wrapper around string that can pretty print the shell
// integration scripts.
type Shell string

// MustCompile compiles the shell template script to ready-to-execute (read
// copy and paste) shell integration.
func (s Shell) MustCompile(context Context) string {
	outputBuffer := &bytes.Buffer{}

	tmpl := template.Must(template.New("shell").Parse(string(s)))
	tmpl.Execute(outputBuffer, context)

	outputBytes, err := io.ReadAll(outputBuffer)
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
	case "pwsh", "powershell":
		return Pwsh
	case "nu", "nushell":
		return Nushell
	default:
		return Bash
	}
}

// Bash is the bash shell integration.
var Bash = Shell(bashScript)

// Zsh is the zsh shell integration.
var Zsh = Shell(zshScript)

// Fish is the fish shell integration.
var Fish = Shell(fishScript)

// Nushell is the nushell shell integration.
var Nushell = Shell(nushellScript)

// Pwsh is the PowerShell shell integration.
var Pwsh = Shell(pwshScript)
