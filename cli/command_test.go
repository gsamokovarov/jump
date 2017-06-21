package cli

import (
	"testing"

	"github.com/gsamokovarov/jump/config"
)

func TestDispatchCommand(t *testing.T) {
	RegisterCommand("test", "A testing command.", func(Args, config.Config) error { return nil })

	args := Args([]string{"test"})
	if _, err := DispatchCommand(args, "default"); err == nil {
		t.Errorf("Expected an error on missing registered default command")
	}

	RegisterCommand("default", "A testing command.", func(Args, config.Config) error { return nil })

	if command, _ := DispatchCommand(args, "default"); command.Name != "test" {
		t.Errorf("Expected test command to be dispatched and executed")
	}
}

func TestCommandIsOption(t *testing.T) {
	command := &Command{"--test", "Testing command.", func(Args, config.Config) error { return nil }}

	if !command.IsOption() {
		t.Errorf("Expected --test command to be an option")
	}
}
