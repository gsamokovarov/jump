package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
)

const osSeparator = string(os.PathSeparator)

func cwdFromArgs(args cli.Args) (string, error) {
	if len(args) == 0 {
		return os.Getwd()
	}

	return filepath.Abs(args.CommandName())
}

func termFromArgs(args cli.Args, conf config.Config) string {
	settings := conf.ReadSettings()

	if settings.Space == config.SpaceSlash {
		return strings.Join(args.Raw(), osSeparator)
	}

	return strings.Join(args.Raw(), "")
}
