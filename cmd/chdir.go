package cmd

import (
	"os"
	"path/filepath"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

func chdirCmd(args cli.Args, conf config.Config) error {
	dir, err := os.Getwd()
	if len(args) == 0 && err != nil {
		return err
	} else {
		if dir, err = filepath.Abs(args.CommandName()); err != nil {
			return err
		}
	}

	entries, err := conf.ReadEntries()
	if err != nil {
		return err
	}

	if entry, found := entries.Find(dir); found {
		entry.UpdateScore()
	} else {
		*entries = append(*entries, *scoring.NewEntry(dir))
	}

	if err := conf.WriteEntries(entries); err != nil {
		return err
	}

	return nil
}

func init() {
	cli.RegisterCommand("chdir", "Update the score of directory during chdir.", chdirCmd)
}
