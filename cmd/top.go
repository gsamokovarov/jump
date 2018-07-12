package cmd

import (
	"sort"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
)

func topCmd(_ cli.Args, conf config.Config) error {
	entries, err := conf.ReadEntries()
	if err != nil {
		return err
	}

	// We usually keep them reversely sort to optimize the fuzzy search.
	sort.Sort(sort.Reverse(entries))

	for _, entry := range entries {
		cli.Outf("%s\n", entry.Path)
	}

	return nil
}

func init() {
	cli.RegisterCommand("top", "Lists the directories as they are scored.", topCmd)
}
