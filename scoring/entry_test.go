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
	entry.UpdateScore()

	assert.Equal(t, 2, entry.Score.Weight)
}

func TestNewEntriesWithWeight(t *testing.T) {
	entry := NewEntryWithWeight("/foo", 2)

	assert.Equal(t, "/foo", entry.Path)
	assert.Equal(t, 2, entry.Score.Weight)
}
