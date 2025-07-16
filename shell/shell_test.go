package shell

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

func TestGuessFish(t *testing.T) {
	assert.Equal(t, Fish, Guess("/usr/local/bin/fish"))
}

func TestFishCompiles(t *testing.T) {
	Fish.MustCompile("j")
}

func TestGuessZsh(t *testing.T) {
	assert.Equal(t, Zsh, Guess("/usr/zsh"))
}

func TestZshCompiles(t *testing.T) {
	Zsh.MustCompile("j")
}

func TestGuessPwsh(t *testing.T) {
	assert.Equal(t, Pwsh, Guess("/usr/bin/pwsh"))
	assert.Equal(t, Pwsh, Guess("~/.dotnet/tools/pwsh"))
}

func TestPwshCompiles(t *testing.T) {
	Pwsh.MustCompile("j")
}

func TestGuessBash(t *testing.T) {
	assert.Equal(t, Bash, Guess("/bin/bash"))
	assert.Equal(t, Bash, Guess("/bin/sh"))
}

func TestBashCompiles(t *testing.T) {
	Bash.MustCompile("j")
}

func TestGuessNushell(t *testing.T) {
	assert.Equal(t, Nushell, Guess("/usr/bin/nu"))
	assert.Equal(t, Nushell, Guess("/opt/homebrew/bin/nu"))
}

func TestNushellCompiles(t *testing.T) {
	Nushell.MustCompile("j")
}
