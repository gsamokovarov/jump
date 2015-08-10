package cmd

import (
	"os"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

func updateCmd(_ cli.Args, conf *config.Config) {
	cwd, err := os.Getwd()
	if err != nil {
		cli.Errf("err: %s\n", err)
		os.Exit(1)
	}

	entries := conf.ReadEntries()
	entry, found := entries.Find(func(i int) bool {
		return entries[i].Path == cwd
	})

	if found {
		entry.UpdateScore()
	} else {
		entries = append(entries, *scoring.NewEntry(cwd))
	}

	if err := conf.WriteEntries(entries); err != nil {
		cli.Errf("err: %s\n", err)
	}
}

func init() {
	cli.RegisterCommand("update", "Updates the score of the working directory.", updateCmd)
}
