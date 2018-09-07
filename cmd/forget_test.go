package cmd

import (
	"os"
	p "path"
	"strings"
	"testing"

	"github.com/gsamokovarov/assert"
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

func Test_forgetCmd(t *testing.T) {
	p := p.Join(td, "web-console")

	conf := &config.Testing{
		Entries: scoring.Entries{scoring.NewEntry(p)},
	}

	output := capture(&os.Stdout, func() {
		assert.Nil(t, forgetCmd(cli.Args{p}, conf))
		assert.Nil(t, cleanCmd(cli.Args{}, conf))
	})

	assert.True(t, strings.Contains(output, "Cleaning"))

	entries, err := conf.ReadEntries()
	assert.Nil(t, err)

	assert.Len(t, 0, entries)
}
