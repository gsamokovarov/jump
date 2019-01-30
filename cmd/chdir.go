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
		entry.UpdateScore()
	} else {
		entries = append(entries, scoring.NewEntry(dir))
	}

	return conf.WriteEntries(entries)
}

func init() {
	cli.RegisterCommand("chdir", "Update the score of directory during chdir.", chdirCmd)
}
