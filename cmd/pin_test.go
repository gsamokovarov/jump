package cmd

import (
	"os"
	p "path"
	"testing"

	"github.com/gsamokovarov/assert"
	"github.com/gsamokovarov/jump/cli"
	s "github.com/gsamokovarov/jump/scoring"
)

func Test_pinCmd(t *testing.T) {
	p1 := p.Join(td, "web")
	p2 := p.Join(td, "web-console")

	conf := &testConfig{
		Entries: s.Entries{
			&s.Entry{p1, &s.Score{Weight: 100, Age: s.Now}},
			&s.Entry{p2, &s.Score{Weight: 90, Age: s.Now}},
		},
	}

	assert.Nil(t, pinCmd(cli.Args{"sait", p2}, conf))

	output := capture(&os.Stdout, func() {
		cdCmd(cli.Args{"sait"}, conf)
	})

	assert.Equal(t, p2+"\n", output)
}
