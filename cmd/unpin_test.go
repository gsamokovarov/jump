package cmd

import (
	"os"
	p "path"
	"testing"

	"github.com/gsamokovarov/assert"
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	s "github.com/gsamokovarov/jump/scoring"
)

func Test_unpinCmd(t *testing.T) {
	p1 := p.Join(td, "web")
	p2 := p.Join(td, "web-console")

	conf := &config.InMemory{
		Entries: s.Entries{
			entry(p2, &s.Score{Weight: 1, Age: s.Now}),
			entry(p1, &s.Score{Weight: 100, Age: s.Now}),
		},
	}

	inside(p2, func() {
		assert.Nil(t, pinCmd(cli.Args{"t b"}, conf))
	})

	assert.Nil(t, unpinCmd(cli.Args{"t/b"}, conf))

	output := capture(&os.Stdout, func() {
		assert.Nil(t, cdCmd(cli.Args{"t b"}, conf))
	})

	assert.Equal(t, p1+"\n", output)
}
