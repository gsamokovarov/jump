package cmd

import (
	"os"
	p "path"
	"testing"

	"github.com/gsamokovarov/assert"
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
)

func Test_chdirCmd(t *testing.T) {
	p1 := p.Join(td, "web-console")
	p2 := p.Join(td, "/client/website")

	conf := &config.InMemory{}

	entries, err := conf.ReadEntries()
	assert.Nil(t, err)

	assert.Len(t, 0, entries)

	// Test that a new entry is added to the list.
	chdirCmd(cli.Args{p1}, conf)

	entries, err = conf.ReadEntries()
	assert.Nil(t, err)

	assert.Len(t, 1, entries)

	// Test that a new entry is added to the list.
	assert.Nil(t, chdirCmd(cli.Args{p2}, conf))

	entries, err = conf.ReadEntries()
	assert.Nil(t, err)

	assert.Len(t, 2, entries)

	// Test that once an existing path is entered again, it's not duplicated in
	// the entries.
	assert.Nil(t, chdirCmd(cli.Args{p2}, conf))

	entries, err = conf.ReadEntries()
	assert.Nil(t, err)

	assert.Len(t, 2, entries)
}

func Test_chdirCmd_cwd(t *testing.T) {
	conf := &config.InMemory{}

	entries, err := conf.ReadEntries()
	assert.Nil(t, err)

	assert.Len(t, 0, entries)

	// Test that the current directory is added to the list.
	assert.Nil(t, chdirCmd(cli.Args{}, conf))

	entries, err = conf.ReadEntries()
	assert.Nil(t, err)

	assert.Len(t, 1, entries)

	cwd, err := os.Getwd()
	assert.Nil(t, err)

	assert.Equal(t, entries[0].Path, cwd)
}
