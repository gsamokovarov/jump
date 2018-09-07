package cmd

import (
	"os"
	"testing"

	"github.com/gsamokovarov/assert"
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
)

func Test_importCmd_autojump(t *testing.T) {
	oldHOME := os.Getenv("HOME")
	defer os.Setenv("HOME", oldHOME)

	os.Setenv("HOME", td)

	conf := &config.Testing{}

	output := capture(&os.Stdout, func() {
		assert.Nil(t, importCmd(cli.Args{"autojump"}, conf))
	})

	assert.Equal(t, `Importing /Users/genadi/Development/jump
Importing /Users/genadi/Development/mock_last_status
Importing /Users/genadi/Development
Importing /Users/genadi/.go/src/github.com/gsamokovarov/jump
Importing /usr/local/Cellar/autojump
Importing /Users/genadi/Development/gloat
`, output)

	assert.Len(t, 6, conf.Entries)
}

func Test_importCmd_z(t *testing.T) {
	oldHOME := os.Getenv("HOME")
	defer os.Setenv("HOME", oldHOME)

	os.Setenv("HOME", td)

	conf := &config.Testing{}

	output := capture(&os.Stdout, func() {
		assert.Nil(t, importCmd(cli.Args{"z"}, conf))
	})

	assert.Equal(t, `Importing /Users/genadi/Development/hack
Importing /Users/genadi/Development/masse
Importing /Users/genadi/Development
Importing /Users/genadi/.go/src/github.com/gsamokovarov/jump
`, output)

	assert.Len(t, 4, conf.Entries)
}

func Test_importCmd_itALL(t *testing.T) {
	oldHOME := os.Getenv("HOME")
	defer os.Setenv("HOME", oldHOME)

	os.Setenv("HOME", td)

	conf := &config.Testing{}

	output := capture(&os.Stdout, func() {
		assert.Nil(t, importCmd(cli.Args{""}, conf))
	})

	assert.Equal(t, `Importing /Users/genadi/Development/hack
Importing /Users/genadi/Development/masse
Importing /Users/genadi/Development
Importing /Users/genadi/.go/src/github.com/gsamokovarov/jump
Importing /Users/genadi/Development/jump
Importing /Users/genadi/Development/mock_last_status
Importing /usr/local/Cellar/autojump
Importing /Users/genadi/Development/gloat
`, output)

	assert.Len(t, 8, conf.Entries)
}
