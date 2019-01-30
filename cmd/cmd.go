package cmd

import (
	"os"
	"path/filepath"

	"github.com/gsamokovarov/jump/cli"
)

func cwdFromArgs(args cli.Args) (string, error) {
	if len(args) == 0 {
		return os.Getwd()
	}

	return filepath.Abs(args.CommandName())
}
