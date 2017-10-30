package cli

import (
	"testing"

	"github.com/gsamokovarov/assert"
	"github.com/gsamokovarov/jump/config"
)

func TestDispatchCommand(t *testing.T) {
	RegisterCommand("test", "A testing command.", commandFn)

	args := Args([]string{"test"})

	_, err := DispatchCommand(args, "default")
	assert.NotNil(t, err)

	RegisterCommand("default", "A testing command.", commandFn)

	command, err := DispatchCommand(args, "default")
	assert.Nil(t, err)

	assert.Equal(t, "test", command.Name)
}

func TestCommandIsOption(t *testing.T) {
	command := &Command{"--test", "Testing command.", commandFn}

	assert.True(t, command.IsOption())
}

func commandFn(Args, config.Config) error { return nil }
