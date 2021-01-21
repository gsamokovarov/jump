package cmd

import (
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
)

const version = "0.40.0"

func versionCmd(cli.Args, config.Config) error {
	cli.Outf("%s\n", version)

	return nil
}

func init() {
	cli.RegisterCommand("--version", "Show version.", versionCmd)
}
