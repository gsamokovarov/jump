package config

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

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
