package cmd

import (
	"os"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

func cleanCmd(args cli.Args, conf config.Config) error {
	entries, err := conf.ReadEntries()
	if err != nil {
		return err
	}

	cleanEntries := make(scoring.Entries, len(entries))
	copy(cleanEntries, entries)

	for _, entry := range entries {
		// Remove the entries that no longer exists.
		if _, err := os.Stat(entry.Path); os.IsNotExist(err) {
			cli.Outf("Cleaning %s\n", entry.Path)
			cleanEntries.Remove(entry.Path)
		}
	}

	return conf.WriteEntries(cleanEntries)
}

func init() {
	cli.RegisterCommand("clean", "Cleans the database of non-existent entries.", cleanCmd)
}
