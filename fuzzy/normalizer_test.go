package fuzzy

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

func TestNormalizeBase(t *testing.T) {
	t.Run("base name", func(t *testing.T) {
		n := NewNormalizer("soc")
		assert.Equal(t, "society", n.NormalizeBase("/Users/foo/Development/society"))
	})

	t.Run("case sensitivity", func(t *testing.T) {
		n := NewNormalizer("Dev")
		assert.Equal(t, "Development", n.NormalizeBase("/Users/foo/Development"))

		n = NewNormalizer("dev")
		assert.Equal(t, "development", n.NormalizeBase("/Users/foo/Development"))
	})
}

func TestScore(t *testing.T) {
	t.Run("no separators matches the base name only", func(t *testing.T) {
		n := NewNormalizer("soc")
		assert.Equal(t, 3, n.Score("/Users/foo/Development/society"))
		assert.True(t, n.Score("/Users/foo/society/website") < 3)
	})

	t.Run("components are matched right-aligned", func(t *testing.T) {
		n := NewNormalizer("dev/soc")
		assert.Equal(t, 6, n.Score("/Users/foo/Development/society"))
		// "dev" is matched against "society" and "soc" against "website" here.
		assert.True(t, n.Score("/Users/foo/society/website") < 6)
	})

	t.Run("absolute terms line up with the same prefix", func(t *testing.T) {
		n := NewNormalizer("/home/foo/Dow")
		downloads := n.Score("/home/foo/Downloads")
		wallpaper := n.Score("/home/foo/Development/ponor-wallpaper")
		assert.True(t, downloads > wallpaper)
	})

	t.Run("case sensitivity", func(t *testing.T) {
		assert.Equal(t, 3, NewNormalizer("Dev").Score("/Users/foo/Development"))
		assert.True(t, NewNormalizer("Dev").Score("/Users/foo/development") < 3)
		assert.Equal(t, 3, NewNormalizer("dev").Score("/Users/foo/Development"))
	})
}
