package cmd

import (
	"os"
	"strings"
	"testing"

	"github.com/gsamokovarov/assert"
	"github.com/gsamokovarov/jump/cli"
)

func Test_cleanCmd(t *testing.T) {
	conf := &testConfig{}

	assert.Nil(t, chdirCmd(cli.Args{"/inexistent/dir/dh891n2kisdha"}, conf))

	entries, err := conf.ReadEntries()
	assert.Nil(t, err)

	assert.Len(t, 1, entries)

	output := capture(&os.Stdout, func() {
		assert.Nil(t, cleanCmd(cli.Args{}, conf))
	})

	assert.True(t, strings.Contains(output, "Cleaning"))
}
