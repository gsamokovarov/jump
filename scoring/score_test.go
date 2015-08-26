package scoring

import "testing"

func TestNewScore(t *testing.T) {
	score := NewScore()

	if score.Age != Now && score.Weight != 1 {
		t.Errorf("Expected age to be %s and weight to be %d", score.Age, score.Weight)
	}

}

func TestCalculateScore(t *testing.T) {
	score1 := Score{2, Now}
	score2 := Score{4, Now}

	if expected, got := 1.38, score2.Calculate()-score1.Calculate(); !inDelta(expected, got) {
		t.Errorf("Expected score2 - score1 to be around expected %f, got %f", expected, got)
	}
}

func TestUpdate(t *testing.T) {
	score := Score{2, Now}
	score.Update()

	if score.Age != Now || score.Weight != 3 {
		t.Errorf("Expected age to be %s and weight to be %d", score.Age, score.Weight)
	}
}

func inDelta(delta, expr float64) bool {
	return delta-0.01 < expr && expr < delta+0.01
}
