package config

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

func TestSearch(t *testing.T) {
	t.Run("save and read last search term from a file", func(t *testing.T) {
		conf := tempConfig(t)

		err := conf.WriteSearch("app", 2)
		assert.Nil(t, err)

		s := conf.ReadSearch()

		assert.Equal(t, 2, s.Index)
		assert.Equal(t, "app", s.Term)
	})
}
