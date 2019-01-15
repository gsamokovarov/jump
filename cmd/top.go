package cmd

import (
	"sort"
	"strings"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

func topCmd(args cli.Args, conf config.Config) error {
	entries, err := conf.ReadEntries()
	if err != nil {
		return err
	}

	if len(args) == 0 {
		// We usually keep them reversely sort to optimize the fuzzy search.
		sort.Sort(sort.Reverse(entries))

		for _, entry := range entries {
			cli.Outf("%s\n", entry.Path)
		}

		return nil
	}

	term := strings.Join(args.Raw(), osSeparator)
	fuzzyEntries := scoring.NewFuzzyEntries(entries, term)

	for _, entry := range fuzzyEntries.Entries {
		cli.Outf("%s\n", entry.Path)
	}

	return nil
}

func init() {
	cli.RegisterCommand("top", "Lists the directories as they are scored.", topCmd)
}
