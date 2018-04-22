package cmd

import (
	"os"
	p "path"
	"strings"
	"testing"

	"github.com/gsamokovarov/assert"
	"github.com/gsamokovarov/jump/cli"
	s "github.com/gsamokovarov/jump/scoring"
)

func Test_hintCmd_short(t *testing.T) {
	p1 := p.Join(td, "web-console")
	p2 := p.Join(td, "/client/website")

	conf := &testConfig{
		Entries: s.Entries{
			&s.Entry{p2, &s.Score{Weight: 90, Age: s.Now}},
			&s.Entry{p1, &s.Score{Weight: 100, Age: s.Now}},
		},
	}

	output := capture(&os.Stdout, func() {
		assert.Nil(t, hintCmd(cli.Args{}, conf))
	})

	lines := strings.Fields(output)
	assert.Len(t, 2, lines)

	assert.Equal(t, p1, lines[0])
	assert.Equal(t, p2, lines[1])
}

func Test_hintCmd_long(t *testing.T) {
	p1 := p.Join(td, "web-console")
	p2 := p.Join(td, "/client/website")
	p3 := p.Join(td, "web")

	conf := &testConfig{
		Entries: s.Entries{
			&s.Entry{p3, &s.Score{Weight: 80, Age: s.Now}},
			&s.Entry{p2, &s.Score{Weight: 90, Age: s.Now}},
			&s.Entry{p1, &s.Score{Weight: 100, Age: s.Now}},
		},
	}

	output := capture(&os.Stdout, func() {
		assert.Nil(t, hintCmd(cli.Args{"wc"}, conf))
	})

	lines := strings.Fields(output)
	assert.Len(t, 1, lines)

	assert.Equal(t, p1, lines[0])

	output = capture(&os.Stdout, func() {
		assert.Nil(t, hintCmd(cli.Args{"webonos"}, conf))
	})

	// If you write more than 6 chars, maybe you need more options.
	lines = strings.Fields(output)
	assert.Len(t, 3, lines)

	output = capture(&os.Stdout, func() {
		assert.Nil(t, hintCmd(cli.Args{"client/webs"}, conf))
	})

	// If you wrote more than 9 chars, well, we tried.
	lines = strings.Fields(output)
	assert.Len(t, 1, lines)

	assert.Equal(t, p2, lines[0])
}
