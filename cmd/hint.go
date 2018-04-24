package cmd

import (
	"strings"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

func hintCmd(args cli.Args, conf config.Config) error {
	term := strings.Join(args.Raw(), osSeparator)

	entries, err := conf.ReadEntries()
	if err != nil {
		return err
	}

	fuzzyEntries := scoring.NewFuzzyEntries(entries, term)
	for _, hint := range fuzzyEntries.Entries {
		cli.Outf("%s\n", hint.Path)
		break
	}

	return nil
}

func init() {
	cli.RegisterCommand("hint", "Hints relevant paths for jumping.", hintCmd)
}
