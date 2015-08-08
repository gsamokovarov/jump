package scoring

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

func NewEntry(path string) *Entry {
	return &Entry{path, NewScore()}
}
