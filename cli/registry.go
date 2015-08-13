package cli

import "sort"

// Every registered command gets saved in this global commands registry.
var Registry = make(commandRegistry)

// A command registry can be queried for the available commands and options.
type CommandRegistry interface {
	Commands() []Command
	Options() []Command
}

// A specific implementation of the CommandsRegistry interface.
type commandRegistry map[string]Command

// Commands returns all of the registered commands, sorted alphabetically.
func (c commandRegistry) Commands() []Command {
	commands := []Command{}

	for _, cmdName := range c.sortedKeys() {
		cmd := c[cmdName]
		if !cmd.IsOption() {
			commands = append(commands, cmd)
		}
	}

	return commands
}

// Commands returns all of the registered options, sorted alphabetically.
func (c commandRegistry) Options() []Command {
	options := []Command{}

	for _, cmdName := range c.sortedKeys() {
		cmd := c[cmdName]
		if cmd.IsOption() {
			options = append(options, cmd)
		}
	}

	return options
}

func (c commandRegistry) sortedKeys() []string {
	var keys []string

	for key := range c {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return keys
}
