package cmd

import (
	"github.com/gsamokovarov/jump/cli"
)

func Example_versionCmd() {
	versionCmd(cli.Args{}, nil)

	// Output:
	// 0.11.0
}
