package config

import (
	"testing"

	"github.com/gsamokovarov/assert"
	"github.com/gsamokovarov/jump/scoring"
)

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
		assert.Equal(t, 0.6931471805599453, entries[0].Score.Calculate())
	})
}
