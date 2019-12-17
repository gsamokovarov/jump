package cmd

import (
	"os"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
)

var pinUsage = `Usage: jump pin term

No term specified. Please specify a term that will be permanently
attached to the current directory. If the term contains spaces, they
will be normalized to OS separators.
`

func pinCmd(args cli.Args, conf config.Config) error {
	term := termFromArgs(args, conf)
	if term == "" {
		cli.Exitf(1, pinUsage)
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	return conf.WritePin(term, dir)
}

func init() {
	cli.RegisterCommand("pin", "Pin a directory to a search term.", pinCmd)
}
