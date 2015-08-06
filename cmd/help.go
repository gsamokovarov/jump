package cmd

import (
	"os"
	"text/template"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
)

const helpUsage = `Usage: jump [COMMAND ...]

Jump to a fuzzy-matched directory passed as an argument.

Commands:{{range .}}
  {{.Name}} {{.Desc}}{{end}}
`

func helpCmd(cli.Args, *config.Config) {
	tmpl := template.Must(template.New("--help").Parse(helpUsage))
	tmpl.Execute(os.Stderr, cli.Commands)
}

func init() {
	cli.RegisterCommand("--help", "Show this screen.", helpCmd)
}
