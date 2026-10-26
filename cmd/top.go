package cmd

import (
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

func topCmd(args cli.Args, conf config.Config) error {
	entries, err := conf.ReadEntries()
	if err != nil {
		return err
	}

	showScore := args.Has("--score")
	args = args.Without("--score")

	// Without a term, the entries are ranked by their score alone.
	term := termFromArgs(args, conf)
	for _, entry := range entries.Ranked(term) {
		topCmdShowEntry(entry, showScore)
	}

	return nil
}

func topCmdShowEntry(entry *scoring.Entry, showScore bool) {
	if showScore {
		cli.Outf("%s %.2f\n", entry.Path, entry.CalculateScore())
	} else {
		cli.Outf("%s\n", entry.Path)
	}
}

func init() {
	cli.RegisterCommand("top", "Lists the directories as they are scored.", topCmd)
}
