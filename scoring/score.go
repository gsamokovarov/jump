package scoring

import (
	"fmt"
	"math"
	"time"
)

// A point of reference Score.Update and Score.Relevance use to reference the
// current time. It is used in testing, so we always have the same current
// time. This is okay for this programs as it won't run for long.
var Now time.Time

// Represents a weight of a score and the age of it.
type Score struct {
	Weight int64
	Age    time.Time
}

// Update the weight and age of the current score.
func (s *Score) Update() {
	s.Weight++
	s.Age = Now
}

// Relevance of a score is the quotient of the score age and the current time.
//
// It is expected to be between 0 and 1. It can be more, though, if the age of
// the score is in the future.
func (s *Score) Relevance() float64 {
	return float64(s.Age.Unix()) / float64(Now.Unix())
}

// Calculate the final score from the score weight and the age.
func (s *Score) Calculate() float64 {
	return float64(s.Weight) * math.Log(s.Relevance())
}

// String gives a string representation to Score. Useful for debugging.
func (s *Score) String() string {
	return fmt.Sprintf("{%s %s}", s.Weight, s.Age)
}

// Create a new score object with default weight of 1 and age set to now.
func NewScore() *Score {
	return &Score{1, Now}
}

func init() {
	Now = time.Now()
}
