package scoring

import (
	"sort"

	"github.com/gsamokovarov/jump/match"
)

type FuzzyEntries struct {
	Entries
	Term string
}

// Less compares the Longest Subsequence Length between the term string and
// every entry. The entries with greater LCS come first.
func (fe FuzzyEntries) Less(i, j int) bool {
	matcher := match.NewMatcher(fe.Term)
	left, right := fe.Entries[i].Path, fe.Entries[j].Path

	return matcher.Match(left, right)
}

func (fe FuzzyEntries) Sort() {
	// If this method is left undefined, when fe.Sort() is called, the
	// Entries.Sort method will be called. In its context, the receiver is
	// Entries, therefore, Entries.Less, and not FuzzyEntries.Less, will be
	// called during sorting.
	sort.Stable(fe)
}

// Select selects the entry with greatest LCS score.
func (fe FuzzyEntries) Select() (entry *Entry, empty bool) {
	fe.Sort()

	if fe.Len() == 0 {
		return nil, true
	}

	return &fe.Entries[0], false
}

// NewFuzzyEntries converts a FuzzyEntries and a target string to a
// FuzzyEntries struct.
//
// Entries is expected to be sorted in ASC before creating the FuzzyEntries.
// This gives us the best match. This is not enforced, however.
func NewFuzzyEntries(entries Entries, target string) *FuzzyEntries {
	return &FuzzyEntries{entries, target}
}
