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

func Test_hintCmd(t *testing.T) {
	t.Run("short input", func(t *testing.T) {
		p1 := p.Join(td, "web-console")
		p2 := p.Join(td, "/client/website")

		conf := &config.InMemory{
			Entries: s.Entries{
				entry(p2, &s.Score{Weight: 90, Age: s.Now}),
				entry(p1, &s.Score{Weight: 100, Age: s.Now}),
			},
		}

		output := capture(&os.Stdout, func() {
			assert.Nil(t, hintCmd(cli.Args{"webcons"}, conf))
		})

		lines := strings.Fields(output)
		assert.Len(t, 2, lines)

		assert.Equal(t, p1, lines[0])
		assert.Equal(t, p2, lines[1])
	})

	t.Run("no entries", func(t *testing.T) {
		conf := &config.InMemory{}

		output := capture(&os.Stdout, func() {
			assert.Nil(t, hintCmd(cli.Args{"webcons"}, conf))
		})

		assert.Equal(t, "", output)
	})
}
