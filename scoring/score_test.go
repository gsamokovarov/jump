package scoring

import (
	"fmt"
	"testing"
)

func TestScoreCalculate(t *testing.T) {
	score1 := Score{2, Now}
	score2 := Score{4, Now}

	if expected, got := 1.38, score2.Calculate()-score1.Calculate(); !inDelta(expected, got) {
		t.Errorf("Expected score2 - score1 to be around expected %f, got %f", expected, got)
	}
}

func TestScoreRelevance(t *testing.T) {
	score := Score{2, Now}

	if !inDelta(score.Relevance(), 2) {
		t.Errorf("Expected relevance %v to be 0", score.Relevance())
	}
}

func TestScoreUpdate(t *testing.T) {
	score := Score{2, Now}
	score.Update()

	if score.Age != Now || score.Weight != 3 {
		t.Errorf("Expected age to be %s and weight to be %d", score.Age, score.Weight)
	}
}

func TestScoreString(t *testing.T) {
	score := Score{2, Now}

	if score.String() == "" {
		t.Errorf("Expected string representation to be ")
	}
}

func TestNewScore(t *testing.T) {
	score := NewScore()
	str := fmt.Sprintf("{1 %s}", score.Age)

	if str != score.String() {
		t.Errorf("Expected %v to be %v", str, score.String())
	}
}

func inDelta(delta, expr float64) bool {
	return delta-0.01 < expr && expr < delta+0.01
}
