package scoring

import "fmt"

// Entry represents a path and a score.
type Entry struct {
	Path  string
	Score *Score
}

// UpdateScore updates the score for an entry.
func (e *Entry) UpdateScore() {
	e.Score.Update()
}

// CalculateScore calculates the score for an entry.
func (e *Entry) CalculateScore() float64 {
	return e.Score.Calculate()
}

func (e *Entry) String() string {
	return fmt.Sprintf("{%s %s}", e.Path, e.Score)
}

// NewEntry creates a new entry with the specified path. The score is created
// with NewScore.
func NewEntry(path string) *Entry {
	return &Entry{path, NewScore()}
}

// NewEntryWithWeight creates a new entry with the specified path and weight.
func NewEntryWithWeight(path string, weight int64) *Entry {
	score := NewScore()
	score.Weight = weight

	return &Entry{path, score}
}
