package main

import (
	"os"

	"github.com/gsamokovarov/jump/cli"
	_ "github.com/gsamokovarov/jump/cmd"
	"github.com/gsamokovarov/jump/config"
)

func main() {
	args := cli.ParseArgs(os.Args)

	config, err := config.SetupDefault(os.Getenv("JUMP_HOME"))
	if err != nil {
		cli.Exitf(1, "bug: %v\n", err)
	}

	command, err := cli.DispatchCommand(args, "--help")
	if err != nil {
		cli.Exitf(1, "bug: %v\n", err)
	}

	if err := command.Action(args.Rest(), config); err != nil {
		cli.Exitf(1, "err: %v\n", err)
	}
}
