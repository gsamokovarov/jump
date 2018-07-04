package cmd

import (
	"github.com/gsamokovarov/jump/cli"
)

func Example_helpCmd() {
	_ = helpCmd(cli.Args{}, nil)

	// Output:
	// Usage: jump [COMMAND ...]
	//
	// Jump to a fuzzy-matched directory passed as an argument.
	//
	// Commands:
	//   cd           Fuzzy match a directory to jump to.
	//   chdir        Update the score of directory during chdir.
	//   clean        Cleans the database of inexisting entries.
	//   forget       Removes the current directory from the database.
	//   hint         Hints relevant paths for jumping.
	//   pin          Pin a directory to a search term.
	//   shell        Display a shell integration script.
	//   top          List the directories as they are scored.
	//   unpin        Unpin a search term.
	//
	// Options:
	//   --help       Show this screen.
	//   --version    Show version.
}
