package cmd

import (
	"os"
	"path/filepath"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

func updateCmd(args cli.Args, conf *config.Config) {
	dir, err := os.Getwd()
	if len(args) == 0 && err != nil {
		cli.Exitf(1, "err: %s\n", err)
	} else {
		if dir, err = filepath.Abs(args.CommandName()); err != nil {
			cli.Exitf(1, "err: %s\n", err)
		}
	}

	entries, err := conf.ReadEntries()
	if err != nil {
		cli.Exitf(1, "err: %s\n", err)
	}

	if entry, found := entries.Find(dir); found {
		entry.UpdateScore()
	} else {
		*entries = append(*entries, *scoring.NewEntry(dir))
	}

	if err := conf.WriteEntries(entries); err != nil {
		cli.Exitf(1, "err: %s\n", err)
	}
}

func init() {
	cli.RegisterCommand("chdir", "Update the score of directory during chdir.", updateCmd)
}
