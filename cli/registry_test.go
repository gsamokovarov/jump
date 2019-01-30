package cli

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

func TestCommandRegistryCommands(t *testing.T) {
	registry := commandRegistry{
		"foo": Command{Name: "foo"},
		"bar": Command{Name: "bar"},
	}

	commands := registry.Commands()
	assert.Len(t, 2, commands)

	assert.Equal(t, "bar", commands[0].Name)
	assert.Equal(t, "foo", commands[1].Name)
}

func TestCommandRegistryOptions(t *testing.T) {
	registry := commandRegistry{
		"foo":      Command{Name: "foo"},
		"--switch": Command{Name: "--switch"},
	}

	options := registry.Options()
	assert.Len(t, 1, options)

	assert.Equal(t, "--switch", options[0].Name)
}
