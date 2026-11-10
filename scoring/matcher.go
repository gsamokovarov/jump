package scoring

import (
	"path/filepath"
	"strings"
)

// exactMatchLength is the term length from which an exact base name match
// outranks the fuzzy matches. Shorter names like app, src or test are common
// to many projects and the heavier entry should still win for them.
const exactMatchLength = 5

// matcher computes how well paths match a term.
//
// Ranking matches every entry against the same term, so the term is split
// once and the LCS rows are reused between the paths.
type matcher struct {
	parts []string
	fold  bool

	current, previous []uint64
}

// newMatcher splits the term into its non-empty components, so repeated and
// trailing separators, as in dev//web or web/, don't shift the matching to
// the wrong ancestor. On Windows both / and \ are separators.
//
// The term is matched case insensitively, unless it has a capital letter in
// it, to save keystrokes.
func newMatcher(term string) *matcher {
	return &matcher{
		parts: strings.FieldsFunc(filepath.ToSlash(term), func(r rune) bool { return r == '/' }),
		fold:  strings.ToLower(term) == term,
	}
}

// match computes how well the path matches the term.
//
// The components of the term and the path are compared right-aligned. The
// last term component is matched against the path's base name, the one before
// it against the parent and so on. The result is the sum of the LCS lengths
// of every pair. The path components the term doesn't reach are ignored.
//
// Keeping the components apart means a term like Dev/jum has to find "Dev"
// in the parent and "jum" in the base name, instead of anywhere in the
// concatenated tail of the path. It also means that an absolute term, as
// produced by a shell expanding ~/Dow, lines up with the entries under the
// same prefix without any special treatment.
//
// An exact match of a base name of at least exactMatchLength counts one
// extra, so a directory can be reached by typing its full name even if a
// heavier entry contains the same letters.
func (m *matcher) match(path string) uint64 {
	path = filepath.ToSlash(path)

	var length uint64
	for i, end := len(m.parts)-1, len(path); i >= 0; i-- {
		for end > 0 && path[end-1] == '/' {
			end--
		}
		if end == 0 {
			break
		}

		start := strings.LastIndexByte(path[:end], '/') + 1

		termPart, pathPart := m.parts[i], path[start:end]
		if m.fold {
			pathPart = strings.ToLower(pathPart)
		}

		length += m.lcs(pathPart, termPart)

		isBase := i == len(m.parts)-1
		if isBase && len(termPart) >= exactMatchLength && termPart == pathPart {
			length++
		}

		end = start
	}

	return length
}

// lcs computes the LCS length of two strings using a two row version of the
// algorithm described at:
//
// https://en.wikipedia.org/wiki/Longest_common_subsequence_problem#Computing_the_length_of_the_LCS.
//
// The rows are kept in the matcher and reused between the calls.
func (m *matcher) lcs(left, right string) uint64 {
	var shorter, longer string

	if len(left) < len(right) {
		shorter, longer = left, right
	} else {
		longer, shorter = left, right
	}

	shorterLen, longerLen := len(shorter), len(longer)

	if cap(m.current) < shorterLen+1 {
		m.current = make([]uint64, shorterLen+1)
		m.previous = make([]uint64, shorterLen+1)
	}

	current, previous := m.current[:shorterLen+1], m.previous[:shorterLen+1]
	clear(current)
	clear(previous)

	for i := range longerLen {
		for j := range shorterLen {
			if shorter[j] == longer[i] {
				current[j+1] = previous[j] + 1
			} else {
				current[j+1] = max(current[j], previous[j+1])
			}
		}

		current, previous = previous, current
	}

	// The rows are swapped after every pass, so the result is in previous.
	return previous[shorterLen]
}
