package fuzzy

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

func TestNormalizePath(t *testing.T) {
	t.Run("no separators", func(t *testing.T) {
		n := NewNormalizer("soc")

		assert.Equal(t, "society",
			n.NormalizePath("/Users/foo/Development/society"))
	})

	t.Run("single separator", func(t *testing.T) {
		n := NewNormalizer("dev/soc")

		assert.Equal(t, "developmentsociety",
			n.NormalizePath("/Users/foo/Development/society"))
	})

	t.Run("multiple separator", func(t *testing.T) {
		n := NewNormalizer("dev/soc/website")

		assert.Equal(t, "developmentsocietywebsite",
			n.NormalizePath("/Users/foo/Development/society/website"))
	})

	t.Run("case sensitivity", func(t *testing.T) {
		n := NewNormalizer("Dev")

		assert.Equal(t, "Development",
			n.NormalizePath("/Users/foo/Development"))
	})
}
