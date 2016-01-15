package cmd

import (
	"os"
	"path/filepath"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

func cdCmd(args cli.Args, conf *config.Config) {
	term := args.CommandName()
	entries, err := conf.ReadEntries()

	if err != nil {
		cli.Exitf(1, "%s\n", err)
	}

	// If an auto-completion triggered a full path, just go there.
	if filepath.IsAbs(term) {
		cli.Exitf(0, "%s\n", term)
	}

	index, search := 0, conf.ReadSearch()

	// If we happen to match the last term, e.g. j is called with no
	// arguments then jump to the previous search.
	if len(term) == 0 {
		term, index = search.Term, search.Index+1
	}

	for {
		fuzzyEntries := scoring.NewFuzzyEntries(entries, term)

		if entry, empty := fuzzyEntries.Select(index); !empty {
			// Remove the entries that no longer exists.
			if _, err := os.Stat(entry.Path); os.IsNotExist(err) {
				entries.Remove(entry.Path)
				conf.WriteEntries(entries)

				continue
			}

			cli.Outf("%s\n", entry.Path)
			conf.WriteSearch(term, index)
		}

		break
	}
}

func init() {
	cli.RegisterCommand("cd", "Fuzzy match a directory to jump to.", cdCmd)
}
