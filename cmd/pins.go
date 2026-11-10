package cmd

import (
	"maps"
	"slices"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
)

func pinsCmd(args cli.Args, conf config.Config) error {
	pins, err := conf.ReadPins()
	if err != nil {
		return err
	}

	for _, term := range slices.Sorted(maps.Keys(pins)) {
		cli.Outf("%s\t%s\n", term, pins[term])
	}

	return nil
}

func init() {
	cli.RegisterCommand("pins", "Lists all the pinned search terms.", pinsCmd)
}
