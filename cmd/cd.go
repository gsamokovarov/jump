package cmd

import (
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

func cdCmd(args cli.Args, conf *config.Config) {
	if len(args) == 0 {
		return
	}

	entries, err := conf.ReadEntries()
	if err != nil {
		cli.Exitf(1, "%s\n", err)
	}

	fuzzyEntries := scoring.NewFuzzyEntries(entries, args.CommandName())
	if entry, empty := fuzzyEntries.Select(); !empty {
		cli.Outf("%s\n", entry.Path)
	}
}

func init() {
	cli.RegisterCommand("cd", "Fuzzy match a directory to jump to.", cdCmd)
}
