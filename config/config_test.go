package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gsamokovarov/assert"
	"github.com/gsamokovarov/jump/scoring"
)

func tempConfig(t *testing.T) Config {
	conf, err := Setup(t.TempDir())
	if err != nil {
		t.Fatalf("Cannot setup temporary testing config: %v", err)
	}

	return conf
}

func TestReadFile(t *testing.T) {
	name := filepath.Join(t.TempDir(), "value.json")

	t.Run("missing file", func(t *testing.T) {
		value := struct{ Ok bool }{true}
		assert.Nil(t, readFile(name, &value))

		assert.True(t, value.Ok)
	})

	t.Run("empty file", func(t *testing.T) {
		assert.Nil(t, os.WriteFile(name, nil, 0644))

		value := struct{ Ok bool }{true}
		assert.Nil(t, readFile(name, &value))

		assert.True(t, value.Ok)
	})

	t.Run("bad JSON", func(t *testing.T) {
		assert.Nil(t, os.WriteFile(name, []byte(`{"Ok":true`), 0644))

		var value struct{ Ok bool }
		assert.NotNil(t, readFile(name, &value))
	})
}

func TestWriteFile(t *testing.T) {
	dir := t.TempDir()
	name := filepath.Join(dir, "value.json")

	t.Run("round trip", func(t *testing.T) {
		assert.Nil(t, writeFile(name, struct{ Ok bool }{true}))

		var value struct{ Ok bool }
		assert.Nil(t, readFile(name, &value))

		assert.True(t, value.Ok)
	})

	t.Run("failed write leaves no temporary file", func(t *testing.T) {
		assert.NotNil(t, writeFile(name, func() {}))

		files, err := os.ReadDir(dir)
		assert.Nil(t, err)

		assert.Len(t, 1, files)
	})
}

func TestEntries(t *testing.T) {
	t.Run("save and read directory score entries from a file", func(t *testing.T) {
		conf := tempConfig(t)

		initial, err := conf.ReadEntries()
		assert.Nil(t, err)

		initial = append(initial, scoring.NewEntry("/test/dir"))

		err = conf.WriteEntries(initial)
		assert.Nil(t, err)

		entries, err := conf.ReadEntries()
		assert.Nil(t, err)

		assert.Len(t, 1, entries)
		assert.Equal(t, "/test/dir", entries[0].Path)
		assert.Equal(t, 4, entries[0].Score.Calculate())
	})
}

func TestSearch(t *testing.T) {
	t.Run("read the last search from the environment", func(t *testing.T) {
		t.Setenv("JUMP_TERM", "app")
		t.Setenv("JUMP_TRY", "2")

		s := tempConfig(t).ReadSearch()

		assert.Equal(t, 2, s.Index)
		assert.Equal(t, "app", s.Term)
	})

	t.Run("no last search", func(t *testing.T) {
		t.Setenv("JUMP_TERM", "")
		t.Setenv("JUMP_TRY", "bogus")

		s := tempConfig(t).ReadSearch()

		assert.Equal(t, 0, s.Index)
		assert.Equal(t, "", s.Term)
	})
}

func TestPins(t *testing.T) {
	t.Run("save and read last search term from a file", func(t *testing.T) {
		conf := tempConfig(t)

		err := conf.WritePin("wc", "/test/web-console")
		assert.Nil(t, err)

		pin, found := conf.FindPin("wc")
		assert.True(t, found)

		assert.Equal(t, "/test/web-console", pin)
	})
}

func TestSettings(t *testing.T) {
	t.Run("save and read settings from a file", func(t *testing.T) {
		conf := tempConfig(t)

		err := conf.WriteSettings(Settings{
			Space:    SpaceIgnore,
			Preserve: true,
		})
		assert.Nil(t, err)

		s := conf.ReadSettings()

		assert.Equal(t, SpaceIgnore, s.Space)
		assert.Equal(t, true, s.Preserve)
	})
}
