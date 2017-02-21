package shell

import (
	"testing"
)

func TestGuessFish(t *testing.T) {
	if Guess("/usr/local/bin/fish") != Fish {
		t.Errorf("Expected /usr/local/bin/fish to match the fish shell")
	}
}

func TestFishCompiles(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected Fish to be a valid template: \n%s", r)
		}
	}()

	Fish.MustCompile("j")
}

func TestGuessZsh(t *testing.T) {
	if Guess("/bin/zsh") != Zsh {
		t.Errorf("Expected /bin/zsh to match the zsh shell")
	}
}

func TestZshCompiles(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected Zsh to be a valid template: \n%s", r)
		}
	}()

	Zsh.MustCompile("j")
}

func TestGuessBash(t *testing.T) {
	if Guess("/bin/bash") != Bash {
		t.Errorf("Expected /bin/bash to match the bash shell")
	}

	if Guess("/bin/sh") != Bash {
		// Its the most common one so fullback to it.
		t.Errorf("Expected unknown shells to match the bash shell")
	}
}

func TestBashCompiles(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected Bash to be a valid template: \n%s", r)
		}
	}()

	Bash.MustCompile("j")
}
