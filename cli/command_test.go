package cli

import "testing"

func TestDispatchCommand(t *testing.T) {
	executed := false
	RegisterCommand("test", "A testing command.", func([]string) { executed = true })

	args := Args([]string{"test"})
	if err := DispatchCommand(args, "default"); err == nil && executed == true {
		t.Errorf("Expected an error on missing registered default command")
	}

	RegisterCommand("default", "A testing command.", func([]string) {})

	if _ = DispatchCommand(args, "default"); executed == false {
		t.Errorf("Expected test command to be dispatched and executed")
	}
}

func TestCommand(t *testing.T) {
	command := &Command{"--test", "Testing command.", func([]string) {}}

	if !command.IsOption() {
		t.Errorf("Expected --test command to be an option")
	}
}
