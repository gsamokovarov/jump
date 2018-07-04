package cmd

import (
	"strings"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
)

var unpinUsageMsg = `jump unpin term

No term specified. Please specify a term to be removed from the pinned
database.
`

func unpinCmd(args cli.Args, conf config.Config) error {
	term := strings.Join(args.Raw(), osSeparator)
	if term == "" {
		cli.Exitf(1, unpinUsageMsg)
	}

	return conf.RemovePin(term)
}

func init() {
	cli.RegisterCommand("unpin", "Unpin a search term.", unpinCmd)
}
