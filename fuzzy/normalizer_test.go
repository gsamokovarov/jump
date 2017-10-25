package fuzzy

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

func TestNoSeparators(t *testing.T) {
	n := NewNormalizer("soc")

	assert.Equal(t, "society",
		n.NormalizePath("/Users/foo/Development/society"))
}

func TestSingleSeparator(t *testing.T) {
	n := NewNormalizer("dev/soc")

	assert.Equal(t, "developmentsociety",
		n.NormalizePath("/Users/foo/Development/society"))
}

func TestMultipleSeparators(t *testing.T) {
	n := NewNormalizer("dev/soc/website")

	assert.Equal(t, "developmentsocietywebsite",
		n.NormalizePath("/Users/foo/Development/society/website"))
}

func TestCaseSensitivity(t *testing.T) {
	n := NewNormalizer("Dev")

	assert.Equal(t, "Development",
		n.NormalizePath("/Users/foo/Development"))
}
