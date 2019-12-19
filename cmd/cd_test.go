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
	conf := &config.InMemory{
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
	conf := &config.InMemory{}

	output := capture(&os.Stderr, func() {
		assert.Nil(t, cdCmd(cli.Args{"wc"}, conf))
	})

	assert.Equal(t, noEntriesMessage, output)
}

func Test_cdCmd_multipleArgumentsAsSeparators(t *testing.T) {
	conf := &config.InMemory{
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
	conf := &config.InMemory{
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
	p2 := p.Join(td, "/client/webcon")
	p3 := p.Join(td, "web")

	conf := &config.InMemory{
		Entries: s.Entries{
			&s.Entry{p3, &s.Score{Weight: 80, Age: s.Now}},
			&s.Entry{p2, &s.Score{Weight: 90, Age: s.Now}},
			&s.Entry{p1, &s.Score{Weight: 100, Age: s.Now}},
		},
	}

	// You have to type at least 5 characters here to trigger the exact match
	// here to avoid jumps to popular `app`, `src`, `test`, `spec` or the likes
	// that are common to project structures.
	output := capture(&os.Stdout, func() {
		assert.Nil(t, cdCmd(cli.Args{"web"}, conf))
	})

	assert.NotEqual(t, p3+"\n", output)
}

func Test_cdCmd_exactMatch_enoughLength(t *testing.T) {
	p1 := p.Join(td, "web-console")
	p2 := p.Join(td, "/client/webcon")
	p3 := p.Join(td, "web")

	conf := &config.InMemory{
		Entries: s.Entries{
			&s.Entry{p3, &s.Score{Weight: 80, Age: s.Now}},
			&s.Entry{p2, &s.Score{Weight: 90, Age: s.Now}},
			&s.Entry{p1, &s.Score{Weight: 100, Age: s.Now}},
		},
	}

	// If someone typed a dir exactly, jump straight to it.
	output := capture(&os.Stdout, func() {
		assert.Nil(t, cdCmd(cli.Args{"webcon"}, conf))
	})

	assert.Equal(t, p2+"\n", output)
}
