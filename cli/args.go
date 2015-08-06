package cli

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

// Rest extracts the arguments after the command name.
func (a Args) Rest() Args {
	if len(a) >= 1 {
		return a[1:]
	}

	return a
}
