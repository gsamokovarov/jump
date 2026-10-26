package fuzzy

import (
	"path/filepath"
	"strings"
)

// exactMatchLength is the term length from which an exact base name match
// outranks the fuzzy matches. Shorter names like app, src or test are common
// to many projects and the heavier entry should still win for them.
const exactMatchLength = 5

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
//
// The term is matched case insensitively, unless it has a capital letter in
// it, to save keystrokes. An exact match of the base name scores above any
// fuzzy one, so a directory can be reached by typing its full name even if a
// heavier entry contains the same letters.
func Score(term, path string) uint64 {
	if strings.ToLower(term) == term {
		path = strings.ToLower(path)
	}

	termParts, pathParts := components(term), components(path)

	var score uint64
	for i, j := len(termParts)-1, len(pathParts)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		score += Length(pathParts[j], termParts[i])
	}

	if len(termParts) > 0 && len(pathParts) > 0 {
		termBase, pathBase := termParts[len(termParts)-1], pathParts[len(pathParts)-1]
		if len(termBase) >= exactMatchLength && termBase == pathBase {
			score++
		}
	}

	return score
}

// components splits a path into its non-empty components.
//
// Repeated and trailing separators, as in dev//web or web/, don't produce
// empty components that would shift the matching to the wrong ancestor. On
// Windows both / and \ are separators, so a term typed with forward slashes
// lines up with the backslash separated entries.
func components(path string) []string {
	return strings.FieldsFunc(filepath.ToSlash(path), func(r rune) bool { return r == '/' })
}
