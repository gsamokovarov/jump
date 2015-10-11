package cli

import "math"

// Represents an ordered collection of command line argument.
type Args []string

// Parse an OS-like arguments.
//
// If you need to create an arguments object out of a plain []string use
// Args{args}.
func ParseArgs(args []string) Args {
	return Args(args[1:])
}

// CommandName extracts a command name out of all the arguments.
func (a Args) CommandName() string {
	if len(a) >= 1 {
		return a[0]
	}

	return ""
}

// Value gets the value for an --option.
//
// The value can be specified by `--count 3`. Currently, no `--count=3` is
// supported.
//
// A default value should be given and will be returned if no value is found.
func (a Args) Value(option, defaultValue string) string {
	stopIndex := int(math.Max(float64(len(a)-1), 0))

	for i := 0; i < stopIndex; i++ {
		if a[i] == option {
			return a[i+1]
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
