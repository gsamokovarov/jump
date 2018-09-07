package cmd

import (
	"os"
	p "path"
	"strings"
	"testing"

	"github.com/gsamokovarov/assert"
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	s "github.com/gsamokovarov/jump/scoring"
)

func Test_cdCmd(t *testing.T) {
	conf := &config.Testing{
		Entries: s.Entries{
			&s.Entry{p.Join(td, "web-console"), &s.Score{Weight: 100, Age: s.Now}},
			&s.Entry{p.Join(td, "/client/website"), &s.Score{Weight: 90, Age: s.Now}},
		},
	}

	output := capture(&os.Stdout, func() {
		assert.Nil(t, cdCmd(cli.Args{"wc"}, conf))
	})

	assert.True(t, strings.Contains(output, p.Join(td, "web-console")))
}

func Test_cdCmd_noEntries(t *testing.T) {
	conf := &config.Testing{}

	output := capture(&os.Stderr, func() {
		assert.Nil(t, cdCmd(cli.Args{"wc"}, conf))
	})

	assert.Equal(t, noEntriesMessage, output)
}

func Test_cdCmd_multipleArgumentsAsSeparators(t *testing.T) {
	conf := &config.Testing{
		Entries: s.Entries{
			&s.Entry{p.Join(td, "web-console"), &s.Score{Weight: 100, Age: s.Now}},
			&s.Entry{p.Join(td, "/client/website"), &s.Score{Weight: 90, Age: s.Now}},
		},
	}

	output := capture(&os.Stdout, func() {
		assert.Nil(t, cdCmd(cli.Args{"cl", "web"}, conf))
	})

	assert.True(t, strings.Contains(output, p.Join(td, "/client/website")))
}

func Test_cdCmd_absolutePath(t *testing.T) {
	conf := &config.Testing{
		Entries: s.Entries{
			&s.Entry{p.Join(td, "web-console"), &s.Score{Weight: 100, Age: s.Now}},
			&s.Entry{p.Join(td, "/client/website"), &s.Score{Weight: 90, Age: s.Now}},
		},
	}

	output := capture(&os.Stdout, func() {
		assert.Nil(t, cdCmd(cli.Args{"/absolute/path"}, conf))
	})

	assert.Equal(t, "/absolute/path\n", output)
}

func Test_cdCmd_exactMatch(t *testing.T) {
	p1 := p.Join(td, "web-console")
	p2 := p.Join(td, "/client/website")
	p3 := p.Join(td, "web")

	conf := &config.Testing{
		Entries: s.Entries{
			&s.Entry{p1, &s.Score{Weight: 100, Age: s.Now}},
			&s.Entry{p2, &s.Score{Weight: 90, Age: s.Now}},
			&s.Entry{p3, &s.Score{Weight: 80, Age: s.Now}},
		},
	}

	output := capture(&os.Stdout, func() {
		assert.Nil(t, cdCmd(cli.Args{"web"}, conf))
	})

	// If someone typed a dir exactly, jump straight to it. Not good for short
	// names like this test here, but pretty useful for most of the catch-all
	// directories.
	assert.Equal(t, p3+"\n", output)
}
