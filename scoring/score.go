package scoring

import (
	"fmt"
	"math"
	"time"
)

// Now is point of reference Score.Update and Score.Relevance use to reference
// the current time. It is used in testing, so we always have the same current
// time. This is okay for this program as it won't run for long.
var Now time.Time

// Score represents a weight of a score and the age of it.
type Score struct {
	Weight int64
	Age    time.Time
}

// Update the weight and age of the current score.
func (s *Score) Update(weight int64) {
	s.Weight += weight
	s.Age = Now
}

// Relevance is a multiplier based on how recently the score was accessed.
//
// 4 if accessed within the last hour, 2 if within the last day, 0.5 if within
// the last week, and 0.25 otherwise.
func (s *Score) Relevance() float64 {
	timeSinceAccess := Now.Sub(s.Age)

	if timeSinceAccess < time.Hour {
		return 4.0
	} else if timeSinceAccess < 24*time.Hour {
		return 2.0
	} else if timeSinceAccess < 7*24*time.Hour {
		return 0.5
	} else {
		return 0.25
	}
}

// Calculate the final score from the score weight and the age.
func (s *Score) Calculate() float64 {
	return float64(s.Weight) * math.Log(s.Relevance())
}

// String gives a string representation to Score. Useful for debugging.
func (s *Score) String() string {
	return fmt.Sprintf("{%d %s}", s.Weight, s.Age)
}

// NewScore creates a new score object with default weight of 1 and age set to
// now.
func NewScore() *Score {
	return &Score{1, Now}
}

func init() {
	Now = time.Now()
}
