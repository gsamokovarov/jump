package cmd

import (
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/shell"
)

func shellCmd(_ cli.Args, conf *config.Config) {
	sh := shell.Guess()

	integration, err := sh.Integration()
	if err != nil {
		cli.Exitf(1, "%s\n", err)
	}

	cli.Outf("%s", integration)
}

func init() {
	cli.RegisterCommand("shell", "Display a shell integration script.", shellCmd)
}
