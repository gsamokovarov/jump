package fuzzy

import (
	"os"
	"path/filepath"
	"strings"
)

const osSeparator = string(os.PathSeparator)

type Normalizer struct {
	term string
}

// Normalize a path for fuzzy searching.
//
// We wanna do this because paths are long and a small fuzzy term is quite
// likely to match almost any path. Most of the times we are either looking for
// the base name of the path or for the last important bits of it.
//
// We apply a couple of conventions:
//
// - If the term has different letter casing in it, do a case sensitive search.
//   By default we do a case insensitive one.
//
// - If the term contains an OS separator, we extract the very last bits of the
//   path that contains that many separators. This helps us to keep the
//   searchable path in a reasonable size, guided by the idea that you may
//   store projects in product/component or client/project kind of structure.
//
// - If the term doesn't contain any OS separators, match on the base name of
//   the path.
func (m Normalizer) NormalizePath(path string) string {
	if caseSensitiveSearch(m.term) {
		path = strings.ToLower(path)
	}

	if containsOsSeparators(m.term) {
		count := strings.Count(m.term, osSeparator)
		path = finalChunks(path, count)
	} else {
		path = filepath.Base(path)
	}

	return path
}

// Normalizes the search term.
//
// The normalization consists only of returning a case insensitive (lowered) or
// sensitive string.
func (m Normalizer) NormalizeTerm() string {
	if caseSensitiveSearch(m.term) {
		return strings.ToLower(m.term)
	}

	return m.term
}

// Create a new normalizer from a stringy term.
func NewNormalizer(term string) *Normalizer {
	return &Normalizer{term}
}

func caseSensitiveSearch(str string) bool {
	return strings.ToLower(str) == str || str == strings.ToUpper(str)
}

func containsOsSeparators(str string) bool {
	return strings.ContainsAny(str, osSeparator)
}

func finalChunks(path string, sepCount int) string {
	var chunk string

	for i := 0; i < sepCount; i++ {
		dir, file := filepath.Split(path)
		chunk = file + chunk
		path = strings.TrimSuffix(dir, osSeparator)
	}

	return chunk
}
