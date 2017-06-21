package cmd

import (
	"os"
	"path/filepath"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
)

var pinUsageMsg = `jump pin term [directory]

  No term specified. See the signature of the pin call above.
`

func pinCmd(args cli.Args, conf config.Config) error {
	var err error

	term := args.CommandName()
	if term == "" {
		cli.Exitf(1, pinUsageMsg)
	}

	dir := args.Rest().CommandName()
	if dir == "" {
		if dir, err = os.Getwd(); err != nil {
			return err
		}
	}

	dir, err = filepath.Abs(dir)
	if err != nil {
		return err
	}

	return conf.WritePin(term, dir)
}

func init() {
	cli.RegisterCommand("pin", "Pin a directory to a search term.", pinCmd)
}
