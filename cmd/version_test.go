package cmd

import (
	"github.com/gsamokovarov/jump/cli"
)

func ExampleVersionCmd() {
	versionCmd(cli.Args{}, nil)

	// Output:
	// 0.11.0
}
