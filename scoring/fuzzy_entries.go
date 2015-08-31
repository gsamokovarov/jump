package scoring

import (
	"sort"
	"strings"

	"github.com/gsamokovarov/jump/lcs"
)

type FuzzyEntries struct {
	Entries
	Target string
}

// Less compares the Longest Subsequence Length between the Target string and
// every entry. The entries with greater LCS come first.
func (fe FuzzyEntries) Less(i, j int) bool {
	iPath := strings.ToLower(fe.Entries[i].BasePath())
	jPath := strings.ToLower(fe.Entries[j].BasePath())

	return lcs.Length(iPath, fe.Target) >= lcs.Length(jPath, fe.Target)
}

func (fe FuzzyEntries) Sort() {
	// If this method is left undefined, when fe.Sort() is called, the
	// Entries.Sort method will be called. In its context, the receiver is
	// Entries, therefore, Entries.Less, and not FuzzyEntries.Less, will be
	// called during sorting.
	sort.Sort(fe)
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
// Entries is expected to be sorted before creating the FuzzyEntries. This
// gives us the best match. This is not enforced, however.
func NewFuzzyEntries(entries Entries, target string) *FuzzyEntries {
	return &FuzzyEntries{entries, target}
}
