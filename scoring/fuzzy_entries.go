package scoring

import (
	"os"
	"path/filepath"
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
	target := fe.Target
	left, right := fe.Entries[i].Path, fe.Entries[j].Path

	if isSameCase(target) {
		target = strings.ToLower(target)
		left, right = strings.ToLower(left), strings.ToLower(right)
	}

	if isInterestingSearch(target) {
		left, right = extractInterestingPath(left), extractInterestingPath(right)
	}

	return lcs.Length(left, target) >= lcs.Length(right, target)
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

func isSameCase(str string) bool {
	return strings.ToLower(str) == str || str == strings.ToUpper(str)
}

func isInterestingSearch(str string) bool {
	return strings.ContainsAny(str, string(os.PathSeparator))
}

func extractInterestingPath(path string) string {
	dir, file := filepath.Split(path)
	return filepath.Base(dir) + file
}
