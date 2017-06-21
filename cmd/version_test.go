package cmd

import (
	"github.com/gsamokovarov/jump/cli"
)

func Example_versionCmd() {
	_ = versionCmd(cli.Args{}, nil)

	// Output:
	// 0.13.0
}
