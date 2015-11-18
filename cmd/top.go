package cmd

import (
	"sort"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
)

func topCmd(args cli.Args, conf *config.Config) {
	entries, err := conf.ReadEntries()
	if err != nil {
		cli.Exitf(1, "%s\n", err)
	}

	// We usually keep them reversely sort to optimize the fuzzy search.
	sort.Sort(sort.Reverse(entries))

	for _, entry := range entries {
		cli.Outf("%s\n", entry.Path)
	}
}

func init() {
	cli.RegisterCommand("top", "List the directories as they are scored.", topCmd)
}
