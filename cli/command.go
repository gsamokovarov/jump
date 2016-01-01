package cli

import (
	"errors"
	"strings"

	"github.com/gsamokovarov/jump/config"
)

// CommandFn represents a command handler function.
type CommandFn func(Args, *config.Config)

// Command represents a command line action.
type Command struct {
	Name   string
	Desc   string
	Action CommandFn
}

// IsOption returns whether the current command starts with --.
//
// We consider commands starting with -- options, so we can display them as such.
func (c *Command) IsOption() bool {
	return strings.HasPrefix(c.Name, "--")
}

// RegisterCommand registers a command in the global command registry.
// ParseArguments looks into it to decide which command to dispatch.
func RegisterCommand(name, desc string, action CommandFn) {
	Registry[name] = Command{name, desc, action}
}

// ErrNoDefaultCommand is used when the default default command isn't
// registered.
var ErrNoDefaultCommand = errors.New("default command is not registered")

// DispatchCommand dispatches the control to an registered command, if
// possible.
//
// A command name is guessed out of the arguments. If the guessed name isn't
// registered, the dispatch will fall-back to the default command specified. It
// is expected that it is always registered. It is an error if its not.
func DispatchCommand(args Args, defaultCommand string) (*Command, error) {
	command, ok := Registry[defaultCommand]
	if !ok {
		return nil, ErrNoDefaultCommand
	}

	if command, ok := Registry[args.First()]; ok {
		return &command, nil
	}

	return &command, nil
}
