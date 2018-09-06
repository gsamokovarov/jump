package cmd

import (
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/importer"
	"github.com/gsamokovarov/jump/scoring"
)

func importCmd(args cli.Args, conf config.Config) error {
	imp := importer.Autojump(conf)

	return imp.Import(func(entry *scoring.Entry) {
		cli.Outf("Importing %s\n", entry.Path)
	})
}

func init() {
	cli.RegisterCommand("import", "Import autojump or z scores.", importCmd)
}
