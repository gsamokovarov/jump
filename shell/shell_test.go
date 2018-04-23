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

func TestGuessBash(t *testing.T) {
	assert.Equal(t, Bash, Guess("/bin/bash"))
	assert.Equal(t, Bash, Guess("/bin/sh"))
}

func TestBashCompiles(t *testing.T) {
	Bash.MustCompile("j")
}
