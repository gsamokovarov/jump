package scoring

import (
	"path/filepath"
	"strings"
)

// exactMatchLength is the term length from which an exact base name match
// outranks the fuzzy matches. Shorter names like app, src or test are common
// to many projects and the heavier entry should still win for them.
const exactMatchLength = 5

// match computes how well the path matches the term.
//
// Both the term and the path are split into their components and compared
// right-aligned, component by component. The last term component is matched
// against the path's base name, the one before it against the parent and so
// on. The result is the sum of the LCS lengths of every pair.
//
// Keeping the components apart means a term like Dev/jum has to find "Dev"
// in the parent and "jum" in the base name, instead of anywhere in the
// concatenated tail of the path. It also means that an absolute term, as
// produced by a shell expanding ~/Dow, lines up with the entries under the
// same prefix without any special treatment.
//
// The term is matched case insensitively, unless it has a capital letter in
// it, to save keystrokes. An exact match of the base name ranks above any
// fuzzy one, so a directory can be reached by typing its full name even if a
// heavier entry contains the same letters.
func match(term, path string) uint64 {
	if strings.ToLower(term) == term {
		path = strings.ToLower(path)
	}

	termParts, pathParts := components(term), components(path)

	var length uint64
	for i, j := len(termParts)-1, len(pathParts)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		length += lcs(pathParts[j], termParts[i])
	}

	if len(termParts) > 0 && len(pathParts) > 0 {
		termBase, pathBase := termParts[len(termParts)-1], pathParts[len(pathParts)-1]
		if len(termBase) >= exactMatchLength && termBase == pathBase {
			length++
		}
	}

	return length
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

// lcs computes the LCS length of two strings using an altered (and memoized)
// version of the algorithm described at:
//
// https://en.wikipedia.org/wiki/Longest_common_subsequence_problem#Computing_the_length_of_the_LCS.
func lcs(left, right string) uint64 {
	var shorter, longer string

	if len(left) < len(right) {
		shorter, longer = left, right
	} else {
		longer, shorter = left, right
	}

	shorterLen, longerLen := len(shorter), len(longer)

	current := make([]uint64, shorterLen+1)
	previous := make([]uint64, shorterLen+1)

	for i := range longerLen {
		for j := range shorterLen {
			if shorter[j] == longer[i] {
				current[j+1] = previous[j] + 1
			} else {
				current[j+1] = max(current[j], previous[j+1])
			}
		}

		for j := range shorterLen {
			previous[j+1] = current[j+1]
		}
	}

	return current[shorterLen]
}
