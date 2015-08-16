package scoring

import (
	"fmt"
	"path/filepath"
)

// An entry represents a path and a score.
type Entry struct {
	Path  string
	Score *Score
}

// Update the score for an entry.
func (e *Entry) UpdateScore() {
	e.Score.Update()
}

// Calculates the score for an entry.
func (e *Entry) CalculateScore() float64 {
	return e.Score.Calculate()
}

// BasePath returns the Base path (the part after the last separator) of the
// current Entry path.
func (e *Entry) BasePath() string {
	return filepath.Base(e.Path)
}

func (e *Entry) String() string {
	return fmt.Sprintf("{%s %s}", e.Path, e.Score)
}

// Create a new entry with the specified path. The score is created with
// NewScore.
func NewEntry(path string) *Entry {
	return &Entry{path, NewScore()}
}
