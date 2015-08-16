package cmd

import (
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/shell"
)

func shellCmd(_ cli.Args, conf *config.Config) {
	sh := shell.Guess()

	cli.Outf("%s", sh)
}

func init() {
	cli.RegisterCommand("shell", "Display a shell integration script.", shellCmd)
}
