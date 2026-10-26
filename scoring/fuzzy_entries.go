package scoring

import (
	"sort"

	"github.com/gsamokovarov/jump/fuzzy"
)

// FuzzyEntries represents entries slice that supports sort.Interface which can
// be sorted by a fuzzy search term.
type FuzzyEntries struct {
	Entries
	Term string
}

// Less compares the Longest Subsequence Length between the term string and
// every entry. The entries with greater LCS come first.
func (fe *FuzzyEntries) Less(i, j int) bool {
	return fuzzy.Score(fe.Term, fe.Entries[i].Path) >= fuzzy.Score(fe.Term, fe.Entries[j].Path)
}

// Sort sorts the entries collection.
func (fe *FuzzyEntries) Sort() {
	// If this method is left undefined, when fe.Sort() is called, the
	// Entries.Sort method will be called. In its context, the receiver is
	// Entries, therefore, Entries.Less, and not FuzzyEntries.Less, will be
	// called during sorting.
	sort.Stable(fe)
}

// Select selects the entry with greatest LCS score at index.
func (fe *FuzzyEntries) Select(index int) (*Entry, bool) {
	if length := fe.Len(); length == 0 || index >= length {
		return nil, false
	}

	return fe.Entries[index], true
}

// NewFuzzyEntries converts a FuzzyEntries and a target string to a
// FuzzyEntries struct.
//
// Entries is expected to be sorted in ASC before creating the FuzzyEntries.
// This gives us the best match. This is not enforced, however.
func NewFuzzyEntries(entries Entries, target string) *FuzzyEntries {
	fuzzyEntries := &FuzzyEntries{entries, target}
	fuzzyEntries.Sort()

	return fuzzyEntries
}
