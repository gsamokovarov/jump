package scoring

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

func TestEntriesCalculateScore(t *testing.T) {
	entry := Entry{"/foo", NewScore()}

	assert.Equal(t, entry.Score.Calculate(), entry.CalculateScore())
}

func TestEntriesUpdateScore(t *testing.T) {
	entry := Entry{"/foo", NewScore()}
	entry.UpdateScore(1)

	assert.Equal(t, 2, entry.Score.Weight)
}
