package cmd

import (
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
)

var unpinUsage = `Usage: jump unpin term

No term specified. Please specify a term to be removed from the pinned
database.
`

func unpinCmd(args cli.Args, conf config.Config) error {
	term := termFromArgs(args, conf)
	if term == "" {
		cli.Exitf(1, unpinUsage)
	}

	return conf.RemovePin(term)
}

func init() {
	cli.RegisterCommand("unpin", "Unpin a search term.", unpinCmd)
}
