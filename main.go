package main

import (
	"fmt"
	"os"

	"github.com/gsamokovarov/jump/cli"
	_ "github.com/gsamokovarov/jump/cmd"
)

func main() {
	args := cli.ParseArgs(os.Args)

	if err := cli.DispatchCommand(args, "--help"); err != nil {
		panic(fmt.Sprintf("bug: %s", err.Error()))
	}
}
