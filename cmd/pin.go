package cmd

import (
	"os"
	"strings"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
)

var pinUsageMsg = `jump pin term

	No term specified. Please specify a term that will be permanently attached to
	the current directory. If the term contains spaces, they will be normalized
	to OS separators.
`

func pinCmd(args cli.Args, conf config.Config) error {
	term := strings.Join(args.Raw(), osSeparator)
	if term == "" {
		cli.Exitf(1, pinUsageMsg)
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
