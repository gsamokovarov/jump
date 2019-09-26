package cmd

import (
	"fmt"
	"os"
	p "path"
	"strings"
	"testing"

	"github.com/gsamokovarov/assert"
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	s "github.com/gsamokovarov/jump/scoring"
)

func Test_topCmd(t *testing.T) {
	var f = fmt.Sprintf

	t.Run("displays the most-visited score without arguments", func(t *testing.T) {
		wc := p.Join(td, "web-console")
		web := p.Join(td, "/client/website")

		conf := &config.Testing{
			Entries: s.Entries{
				&s.Entry{wc, &s.Score{Weight: 100, Age: s.Now}},
				&s.Entry{web, &s.Score{Weight: 90, Age: s.Now}},
			},
		}

		output := capture(&os.Stdout, func() {
			assert.Nil(t, topCmd(cli.Args{}, conf))
		})

		lines := strings.Split(output, "\n")
		assert.Len(t, 3, lines)

		assert.Equal(t, f("%s %.2f", wc, conf.Entries[0].CalculateScore()), lines[0])
		assert.Equal(t, f("%s %.2f", web, conf.Entries[1].CalculateScore()), lines[1])
		assert.Equal(t, "", lines[2])
	})

	t.Run("displays the score of a term with arguments", func(t *testing.T) {
		wc := p.Join(td, "web-console")
		neu := p.Join(td, "neuvents")

		conf := &config.Testing{
			Entries: s.Entries{
				&s.Entry{wc, &s.Score{Weight: 100, Age: s.Now}},
				&s.Entry{neu, &s.Score{Weight: 90, Age: s.Now}},
			},
		}

		output := capture(&os.Stdout, func() {
			assert.Nil(t, topCmd(cli.Args{"neuv"}, conf))
		})

		lines := strings.Split(output, "\n")
		assert.Len(t, 3, lines)

		assert.Equal(t, f("%s %.2f", neu, conf.Entries[0].CalculateScore()), lines[0])
		assert.Equal(t, f("%s %.2f", wc, conf.Entries[1].CalculateScore()), lines[1])
		assert.Equal(t, "", lines[2])
	})
}
