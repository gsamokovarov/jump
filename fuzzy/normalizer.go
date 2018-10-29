package fuzzy

import (
	"os"
	"path/filepath"
	"strings"
)

const osSeparator = string(os.PathSeparator)
const glob = "*"

// Normalizer holds utilities for search term transformation.
type Normalizer struct {
	term string
}

// NormalizePath normalizes a path for fuzzy searching.
//
// We wanna do this because paths are long and a small fuzzy term is quite
// likely to match almost any path. Most of the times we are either looking for
// the base name of the path or for the last important bits of it.
//
// We apply a couple of conventions:
//
// - If the term has capital letter , do a case sensitive search. By default
//   we do a case insensitive one to save keystrokes.
//
// - If the term contains an OS separator, we extract the very last bits of the
//   path that contains that many separators. This helps us to keep the
//   searchable path in a reasonable size, guided by the idea that you may
//   store projects in product/component or client/project kind of structure.
//
// - If the term doesn't contain any OS separators, match on the base name of
//   the path.
func (m Normalizer) NormalizePath(path string) string {
	if caseInsensitiveSearch(m.term) {
		path = strings.ToLower(path)
	}

	if containsGlobs(m.term) {
		return path
	}

	if containsOsSeparators(m.term) {
		return finalChunks(path, strings.Count(m.term, osSeparator))
	}

	return filepath.Base(path)
}

// NormalizeTerm normalizes the search term.
//
// The normalization consists only of returning a case insensitive (lowered) or
// sensitive string.
func (m Normalizer) NormalizeTerm() string {
	if containsGlobs(m.term) {
		return strings.Replace(m.term, "*", "", -1)
	}

	return m.term
}

// NewNormalizer create a new normalizer from a stringy term.
func NewNormalizer(term string) *Normalizer {
	return &Normalizer{term}
}

func caseInsensitiveSearch(str string) bool {
	return strings.ToLower(str) == str
}

func containsOsSeparators(str string) bool {
	return strings.ContainsAny(str, osSeparator)
}

func containsGlobs(str string) bool {
	return strings.ContainsAny(str, glob)
}

func finalChunks(path string, sepCount int) string {
	var chunk string

	for i := 0; i <= sepCount; i++ {
		dir, file := filepath.Split(path)
		chunk = file + chunk
		path = strings.TrimSuffix(dir, osSeparator)
	}

	return chunk
}
