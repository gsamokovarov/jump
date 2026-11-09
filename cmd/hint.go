package cmd

import (
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
)

const hintLimit = 5

func hintCmd(args cli.Args, conf config.Config) error {
	term := termFromArgs(args, conf)

	entries, err := conf.ReadEntries()
	if err != nil {
		return err
	}

	ranked := entries.Ranked(term)
	for _, entry := range ranked[:min(len(ranked), hintLimit)] {
		cli.Outf("%s\n", entry.Path)
	}

	return nil
}

func init() {
	cli.RegisterCommand("hint", "Hints relevant paths for jumping.", hintCmd)
}
