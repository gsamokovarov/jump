package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gsamokovarov/assert"
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
