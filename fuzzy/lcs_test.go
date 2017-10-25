package fuzzy

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

func TestTwoNaiveStrings(t *testing.T) {
	assert.Equal(t, 2, Length("fd", "falcon-dev"))
}

func TestLongerAfterShorter(t *testing.T) {
	assert.Equal(t, Length("falcon-dev", "fd"), Length("fd", "falcon-dev"))
}
