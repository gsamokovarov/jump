package cli

import (
	"errors"
	"strings"
)

// Every registered command gets saved in this global commands registry.
var Commands = map[string]Command{}

type CommandFn func([]string)

// Represents a command line action.
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

// Register a command in the global command registry. ParseArguments looks into
// it to decide which command to dispatch.
func RegisterCommand(name, desc string, action func([]string)) {
	Commands[name] = Command{name, desc, action}
}

// Used when the default default command isn't registered.
var ErrNoDefaultCommand = errors.New("default command is not registered")

// Dispatches the control to an registered command, if possible.
//
// A command name is guessed out of the arguments. If the guessed name isn't
// registered, the dispatch will fall-back to the default command specified. It
// is expected that it is always registered. It is an error if its not.
func DispatchCommand(args Args, defaultCommand string) error {
	if _, ok := Commands[defaultCommand]; !ok {
		return ErrNoDefaultCommand
	}

	if command, ok := Commands[args.CommandName()]; ok {
		command.Action(args.Rest())
	} else {
		Commands[defaultCommand].Action(args.Rest())
	}

	return nil
}
