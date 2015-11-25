package cmd

import (
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

	if filepath.IsAbs(term) {
		// If an auto-completion triggered a full path, just go there.
		cli.Outf("%s\n", term)
	} else {
		index, search := 0, conf.ReadSearch()

		// If we happen to match the last term, e.g. j is called with no
		// arguments then jump to the previous search.
		if len(term) == 0 {
			term, index = search.Term, search.Index+1
		}

		fuzzyEntries := scoring.NewFuzzyEntries(entries, term)
		if entry, empty := fuzzyEntries.Select(index); !empty {
			cli.Outf("%s\n", entry.Path)
			conf.WriteSearch(term, index)
		}
	}
}

func init() {
	cli.RegisterCommand("cd", "Fuzzy match a directory to jump to.", cdCmd)
}
