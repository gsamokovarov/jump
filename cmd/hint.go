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
	hints := hintSelect(fuzzyEntries.Entries, term)

	for _, entry := range hints {
		cli.Outf("%s\n", entry.Path)
	}

	return nil
}

func hintSelect(entries scoring.Entries, term string) scoring.Entries {
	termLength := len(term)

	switch {
	case termLength == 0:
		return hintSliceEntries(entries, 5)
	case termLength < 7:
		return hintSliceEntries(entries, 1)
	case termLength >= 7 && termLength < 10:
		return hintSliceEntries(entries, 3)
	default:
		return hintSliceEntries(entries, 1)
	}
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
