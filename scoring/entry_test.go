package scoring

import "testing"

func TestEntriesCalculateScore(t *testing.T) {
	entry := Entry{"/foo", NewScore()}

	if got, expected := entry.CalculateScore(), entry.Score.Calculate(); got != expected {
		t.Errorf("Expected entry.CalculateScore to be %f, got %f", expected, got)
	}
}
