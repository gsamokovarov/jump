package cmd

import (
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

func hintCmd(args cli.Args, conf config.Config) error {
	term := termFromArgs(args, conf)

	entries, err := conf.ReadEntries()
	if err != nil {
		return err
	}

	fuzzyEntries := scoring.NewFuzzyEntries(entries, term)
	hints := hintSliceEntries(fuzzyEntries.Entries, 5)

	for _, entry := range hints {
		cli.Outf("%s\n", entry.Path)
	}

	return nil
}

func hintSliceEntries(entries scoring.Entries, limit int) scoring.Entries {
	if limit < len(entries) {
		return entries[0:limit]
	}

	return entries
}

func init() {
	cli.RegisterCommand("hint", "Hints relevant paths for jumping.", hintCmd)
}
