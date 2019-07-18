package cmd

import (
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

func chdirCmd(args cli.Args, conf config.Config) error {
	dir, err := cwdFromArgs(args)
	if err != nil {
		return err
	}

	entries, err := conf.ReadEntries()
	if err != nil {
		return err
	}

	if entry, found := entries.Find(dir); found {
		entry.UpdateScore(chdirWeight(conf))
	} else {
		entries = append(entries, scoring.NewEntry(dir))
	}

	return conf.WriteEntries(entries)
}

const chdirBadMatchFactor = 4

func chdirWeight(conf config.Config) int64 {
	try := conf.ReadSearch().Index

	// Don't boost the jump if no one called `j` without arguments before and
	// don't over-boost entries if we're spamming the `j` key.
	if try == 0 || try >= chdirBadMatchFactor {
		return 1
	}

	return int64(try+1) * chdirBadMatchFactor
}

func init() {
	cli.RegisterCommand("chdir", "Update the score of directory during chdir.", chdirCmd)
}
