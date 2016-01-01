package scoring

import "testing"

func TestEntriesCalculateScore(t *testing.T) {
	entry := Entry{"/foo", NewScore()}

	if got, expected := entry.CalculateScore(), entry.Score.Calculate(); got != expected {
		t.Errorf("Expected entry.CalculateScore to be %f, got %f", expected, got)
	}
}

func TestEntriesUpdateScore(t *testing.T) {
	entry := Entry{"/foo", NewScore()}
	entry.UpdateScore()

	if got, expected := entry.Score.Weight, int64(2); got != expected {
		t.Errorf("Expected entry.UpdateScore Weight to be %d, got %d", expected, got)
	}
}
