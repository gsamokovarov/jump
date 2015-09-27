package cmd

import (
	"sort"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

func hintCmd(args cli.Args, conf *config.Config) {
	var hints scoring.Entries

	entries, err := conf.ReadEntries()
	if err != nil {
		cli.Exitf(1, "%s\n", err)
	}

	if len(args) == 0 {
		// We usually keep them reversely sort to optimize the fuzzy search.
		sort.Sort(sort.Reverse(entries))

		hints = entries
	} else {
		fuzzyEntries := scoring.NewFuzzyEntries(entries, args.CommandName())
		fuzzyEntries.Sort()

		hints = fuzzyEntries.Entries
	}

	for _, entry := range hints {
		cli.Outf("%s\n", entry.Path)
	}
}

func init() {
	cli.RegisterCommand("hint", "Hints relevant paths for jumping.", hintCmd)
}
