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
			entry(p.Join(td, "web-console"), &s.Score{Weight: 100, Age: s.Now}),
			entry(p.Join(td, "/client/website"), &s.Score{Weight: 90, Age: s.Now}),
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
			entry(p.Join(td, "web-console"), &s.Score{Weight: 100, Age: s.Now}),
			entry(p.Join(td, "/client/website"), &s.Score{Weight: 90, Age: s.Now}),
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
			entry(p.Join(td, "web-console"), &s.Score{Weight: 100, Age: s.Now}),
			entry(p.Join(td, "/client/website"), &s.Score{Weight: 90, Age: s.Now}),
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
			entry(p3, &s.Score{Weight: 80, Age: s.Now}),
			entry(p2, &s.Score{Weight: 90, Age: s.Now}),
			entry(p1, &s.Score{Weight: 100, Age: s.Now}),
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

func Test_cdCmd_implicitConfigSpace(t *testing.T) {
	p1 := p.Join(td, "web-console")
	p2 := p.Join(td, "/client/webcon")
	p3 := p.Join(td, "web")

	conf := &config.InMemory{
		Entries: s.Entries{
			entry(p3, &s.Score{Weight: 80, Age: s.Now}),
			entry(p2, &s.Score{Weight: 90, Age: s.Now}),
			entry(p1, &s.Score{Weight: 100, Age: s.Now}),
		},
	}

	output := capture(&os.Stdout, func() {
		assert.Nil(t, cdCmd(cli.Args{"cli", "web"}, conf))
	})

	assert.Equal(t, p2+"\n", output)
}

func Test_cdCmd_explicitConfigSpace(t *testing.T) {
	p1 := p.Join(td, "web-console")
	p2 := p.Join(td, "/client/webcon")
	p3 := p.Join(td, "web")

	conf := &config.InMemory{
		Entries: s.Entries{
			entry(p3, &s.Score{Weight: 80, Age: s.Now}),
			entry(p2, &s.Score{Weight: 90, Age: s.Now}),
			entry(p1, &s.Score{Weight: 100, Age: s.Now}),
		},
		Settings: config.Settings{Space: config.SpaceSlash},
	}

	output := capture(&os.Stdout, func() {
		assert.Nil(t, cdCmd(cli.Args{"cli", "web"}, conf))
	})

	assert.Equal(t, p2+"\n", output)
}

func Test_cdCmd_exactMatch_enoughLength(t *testing.T) {
	p1 := p.Join(td, "web-console")
	p2 := p.Join(td, "/client/webcon")
	p3 := p.Join(td, "web")

	conf := &config.InMemory{
		Entries: s.Entries{
			entry(p3, &s.Score{Weight: 80, Age: s.Now}),
			entry(p2, &s.Score{Weight: 90, Age: s.Now}),
			entry(p1, &s.Score{Weight: 100, Age: s.Now}),
		},
	}

	// If someone typed a dir exactly, jump straight to it.
	output := capture(&os.Stdout, func() {
		assert.Nil(t, cdCmd(cli.Args{"webcon"}, conf))
	})

	assert.Equal(t, p2+"\n", output)
}

func Test_cdCmd_baseDir(t *testing.T) {
	baseDir := p.Join(td, "client")

	conf := &config.InMemory{
		Entries: s.Entries{
			entry(p.Join(baseDir, "website"), &s.Score{Weight: 100, Age: s.Now}),
			entry(p.Join(baseDir, "webtools"), &s.Score{Weight: 90, Age: s.Now}),
			entry(p.Join(td, "web-console"), &s.Score{Weight: 200, Age: s.Now}), // Higher score but not under baseDir
		},
	}

	t.Run("finds entry under base directory", func(t *testing.T) {
		output := capture(&os.Stdout, func() {
			assert.Nil(t, cdCmd(cli.Args{baseDir, "website"}, conf))
		})

		assert.Equal(t, p.Join(baseDir, "website")+"\n", output)
	})

	t.Run("fuzzy matches under base directory", func(t *testing.T) {
		output := capture(&os.Stdout, func() {
			assert.Nil(t, cdCmd(cli.Args{baseDir, "web"}, conf))
		})

		assert.Equal(t, p.Join(baseDir, "website")+"\n", output)
	})

	t.Run("returns base directory when no matches found", func(t *testing.T) {
		conf := &config.InMemory{
			Entries: s.Entries{},
		}

		output := capture(&os.Stdout, func() {
			assert.Nil(t, cdCmd(cli.Args{baseDir, "foo"}, conf))
		})

		assert.Equal(t, baseDir+"\n", output)
	})

	t.Run("ignores entries outside base directory", func(t *testing.T) {
		// Even though web-console has higher score, it should not be returned
		// because it's not under the base directory
		output := capture(&os.Stdout, func() {
			assert.Nil(t, cdCmd(cli.Args{baseDir, "web"}, conf))
		})

		assert.NotEqual(t, p.Join(td, "web-console")+"\n", output)
		assert.Equal(t, p.Join(baseDir, "website")+"\n", output)
	})

	t.Run("direct children", func(t *testing.T) {
		baseDir := td
		childDir := "web"

		conf := &config.InMemory{
			Entries: s.Entries{},
		}

		output := capture(&os.Stdout, func() {
			assert.Nil(t, cdCmd(cli.Args{baseDir, childDir}, conf))
		})

		// Should return the direct child path since td/web exists
		assert.Equal(t, p.Join(baseDir, childDir)+"\n", output)
	})

	t.Run("acts like cd", func(t *testing.T) {
		baseDir := p.Join(td, "client")
		childDir := "web"

		conf := &config.InMemory{
			Entries: s.Entries{},
		}

		inside(td, func() {
			output := capture(&os.Stdout, func() {
				assert.Nil(t, cdCmd(cli.Args{baseDir, childDir}, conf))
			})

			expectedPath := p.Join(td, childDir)
			assert.Equal(t, expectedPath+"\n", output)
		})
	})
}
