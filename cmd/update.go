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
		cli.Errf("err: %s\n", err)
		os.Exit(1)
	} else {
		dir, err = filepath.Abs(args.CommandName())
		if err != nil {
			cli.Errf("err: %s\n", err)
			os.Exit(1)
		}
	}

	entries, err := conf.ReadEntries()
	if err != nil {
		cli.Errf("err: %s\n", err)
		os.Exit(1)
	}

	println("Finding")
	entry, found := entries.Find(dir)

	if found {
		entry.UpdateScore()
	} else {
		entries = append(entries, *scoring.NewEntry(dir))
	}

	if err := conf.WriteEntries(entries); err != nil {
		cli.Errf("err: %s\n", err)
	}
}

func init() {
	cli.RegisterCommand("update", "Updates the score of a directory.", updateCmd)
}
