package main

import (
	"fmt"
	"os"

	"github.com/gsamokovarov/jump/cli"
	_ "github.com/gsamokovarov/jump/cmd"
	"github.com/gsamokovarov/jump/config"
)

func main() {
	args := cli.ParseArgs(os.Args)

	if err := config.SetupDefault(os.Getenv("JUMP_HOME")); err != nil {
		panic(fmt.Sprintf("bug: %s", err.Error()))
	}

	if err := cli.DispatchCommand(args, "--help"); err != nil {
		panic(fmt.Sprintf("bug: %s", err.Error()))
	}
}
