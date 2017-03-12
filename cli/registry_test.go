package cli

import (
	"testing"
)

func TestCommandRegistryCommands(t *testing.T) {
	registry := commandRegistry{
		"foo": Command{Name: "foo"},
		"bar": Command{Name: "bar"},
	}

	for i, cmd := range registry.Commands() {
		if i == 0 && cmd.Name != "bar" {
			t.Error("Expected command at 0 to be bar, got %v", cmd)
		}

		if i == 1 && cmd.Name != "foo" {
			t.Error("Expected command at 1 to be foo, got %v", cmd)
		}
	}
}

func TestCommandRegistryOptions(t *testing.T) {
	registry := commandRegistry{
		"foo":    Command{Name: "foo"},
		"--halp": Command{Name: "--halp"},
	}

	options := registry.Options()

	for i, cmd := range options {
		if i == 0 && cmd.Name != "--halp" {
			t.Error("Expected command at 0 to be --halp, got %v", cmd)
		}
	}

	if len(options) != 1 {
		t.Error("Expected only 1 option, got %v", options)
	}
}
