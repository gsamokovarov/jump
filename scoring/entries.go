package scoring

import (
	"cmp"
	"slices"
	"strings"

	"github.com/gsamokovarov/jump/fuzzy"
)

// Entries represents the scored directories in the database.
type Entries []*Entry

// Find finds an entry by a given path.
//
// If the entry isn't found, (nil, false) is returned.
func (e Entries) Find(path string) (*Entry, bool) {
	for _, entry := range e {
		if entry.Path == path {
			return entry, true
		}
	}

	return nil, false
}

// Remove an entry by its path.
func (e *Entries) Remove(path string) bool {
	for i, entry := range *e {
		if entry.Path == path {
			*e = append((*e)[:i], (*e)[i+1:]...)
			return true
		}
	}

	return false
}

// Under selects entries within a specific path.
func (e Entries) Under(path string) Entries {
	var entries Entries

	for _, entry := range e {
		if strings.HasPrefix(entry.Path, path) && entry.Path != path {
			entries = append(entries, entry)
		}
	}

	return entries
}

// Ranked returns a copy of the entries ordered by how well they match the
// term, best match first. Entries that match equally well are ordered by
// their score, so the more frequently visited one comes first.
func (e Entries) Ranked(term string) Entries {
	ranked := slices.Clone(e)
	slices.SortStableFunc(ranked, func(a, b *Entry) int {
		return cmp.Or(
			cmp.Compare(fuzzy.Score(term, b.Path), fuzzy.Score(term, a.Path)),
			cmp.Compare(b.CalculateScore(), a.CalculateScore()),
		)
	})

	return ranked
}

// Sort orders the entries by their score, lightest first. The database is
// kept in that order on disk.
func (e Entries) Sort() {
	slices.SortFunc(e, func(a, b *Entry) int {
		return cmp.Compare(a.CalculateScore(), b.CalculateScore())
	})
}
