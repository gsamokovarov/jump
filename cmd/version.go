package cmd

import "github.com/gsamokovarov/jump/cli"

const version = "0.0.1"

func versionCmd([]string) {
	cli.Errf("%s\n", version)
}

func init() {
	cli.RegisterCommand("--version", "Show version.", versionCmd)
}
