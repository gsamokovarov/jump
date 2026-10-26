package fuzzy

import (
	"os"
	"path/filepath"
	"strings"
)

const osSeparator = string(os.PathSeparator)

// Normalizer holds utilities for search term transformation.
type Normalizer struct {
	term string
}

// NormalizeBase returns the base name of the path, folded to the case of
// the term. The term is matched case insensitively unless it has a capital
// letter in it, to save keystrokes.
func (m Normalizer) NormalizeBase(path string) string {
	if caseInsensitiveSearch(m.term) {
		path = strings.ToLower(path)
	}

	return filepath.Base(path)
}

// NewNormalizer create a new normalizer from a stringy term.
func NewNormalizer(term string) *Normalizer {
	return &Normalizer{term}
}

func caseInsensitiveSearch(str string) bool {
	return strings.ToLower(str) == str
}

// Score computes how well the path matches the term.
//
// Both the term and the path are split into their components and compared
// right-aligned, component by component. The last term component is matched
// against the path's base name, the one before it against the parent and so
// on. The score is the sum of the LCS lengths of every pair.
//
// Keeping the components apart means a term like Dev/jum has to find "Dev"
// in the parent and "jum" in the base name, instead of anywhere in the
// concatenated tail of the path. It also means that an absolute term, as
// produced by a shell expanding ~/Dow, lines up with the entries under the
// same prefix without any special treatment.
func (m Normalizer) Score(path string) uint64 {
	if caseInsensitiveSearch(m.term) {
		path = strings.ToLower(path)
	}

	termParts := strings.Split(m.term, osSeparator)
	pathParts := strings.Split(path, osSeparator)

	var score uint64
	for i, j := len(termParts)-1, len(pathParts)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		score += Length(pathParts[j], termParts[i])
	}

	return score
}
