package cli

import "strings"

// Args represents an ordered collection of command line argument.
type Args []string

// ParseArgs parses an OS-like command line input arguments.
//
// If you need to create an arguments object out of a plain []string use
// Args{args}.
func ParseArgs(args []string) Args {
	return Args(args[1:])
}

// First extracts the first argument, no matter if command or option.
func (a Args) First() string {
	if len(a) > 0 {
		return a[0]
	}

	return ""
}

// CommandName extracts a command name out of all the arguments.
//
// A command cannot start with --. We think of those arguments as options.
func (a Args) CommandName() string {
	for i := 0; i < len(a); i++ {
		if !strings.HasPrefix(a[i], "--") {
			return a[i]
		}
	}

	return ""
}

// Has tells whether the arguments contains a specific argument.
//
// No distinction is made between arguments or options. Everything is matched
// as is.
func (a Args) Has(option string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == option {
			return true
		}
	}

	return false
}

// Get checks if an --option has a value given.
//
// If the option is not given, the default value is returned. There is
// currently no concept of optional arguments, as we haven't got the need.
func (a Args) Get(option, defaultValue string) string {
	for i := 0; i < len(a); i++ {
		name, value := nameValue(a[i])

		if name == option {
			if value == "" {
				if i+1 < len(a) {
					return a[i+1]
				}
				return defaultValue
			}

			return value
		}
	}

	return defaultValue
}

// Rest extracts the arguments after the command name.
func (a Args) Rest() Args {
	if len(a) >= 1 {
		return a[1:]
	}

	return a
}

func nameValue(arg string) (name string, value string) {
	parts := strings.SplitN(arg, "=", 2)

	if len(parts) == 1 {
		return parts[0], ""
	}

	return parts[0], parts[1]
}
