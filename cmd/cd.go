package cmd

import (
	"path/filepath"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

func cdCmd(args cli.Args, conf *config.Config) {
	if len(args) == 0 {
		return
	}

	commandName := args.CommandName()
	entries, err := conf.ReadEntries()
	if err != nil {
		cli.Exitf(1, "%s\n", err)
	}

	if filepath.IsAbs(commandName) {
		// If an auto-completion triggered a full path, just go there.
		cli.Outf("%s\n", commandName)
	} else {
		fuzzyEntries := scoring.NewFuzzyEntries(entries, args.CommandName())
		if entry, empty := fuzzyEntries.Select(); !empty {
			cli.Outf("%s\n", entry.Path)
		}
	}
}

func init() {
	cli.RegisterCommand("cd", "Fuzzy match a directory to jump to.", cdCmd)
}
