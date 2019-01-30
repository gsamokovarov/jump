package cmd

import (
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
)

func forgetCmd(args cli.Args, conf config.Config) error {
	dir, err := cwdFromArgs(args)
	if err != nil {
		return err
	}

	entries, err := conf.ReadEntries()
	if err != nil {
		return err
	}

	if entry, found := entries.Find(dir); found {
		cli.Outf("Cleaning %s\n", entry.Path)
		entries.Remove(entry.Path)

		return conf.WriteEntries(entries)
	}

	return nil
}

func init() {
	cli.RegisterCommand("forget", "Removes the current directory from the database.", forgetCmd)
}
