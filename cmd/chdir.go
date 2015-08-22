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
		dir, err = filepath.Abs(args.CommandName())
		if err != nil {
			cli.Exitf(1, "err: %s\n", err)
		}
	}

	entries, err := conf.ReadEntries()
	if err != nil {
		cli.Exitf(1, "err: %s\n", err)
	}

	entry, found := entries.Find(dir)

	if found {
		entry.UpdateScore()
	} else {
		entries = append(entries, *scoring.NewEntry(dir))
	}

	if err := conf.WriteEntries(entries); err != nil {
		cli.Exitf(1, "err: %s\n", err)
	}
}

func init() {
	cli.RegisterCommand("chdir", "Update the scrore of directory during chdir.", updateCmd)
}
