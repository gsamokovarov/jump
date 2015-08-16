package scoring

import "sort"

// An entries slice that supports sort.Interface.
type Entries []Entry

// Len returns the length of the entries slice.
func (e Entries) Len() int {
	return len(e)
}

// Swaps the values at two indexes in the entries slice.
func (e Entries) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// Less compares two indexes in the entries and returns true if the value at i
// is greater than the value at j.
//
// This is to ensure a descending order as we want the bigger elements on top.
func (e Entries) Less(i, j int) bool {
	return e[i].CalculateScore() >= e[j].CalculateScore()
}

// Find finds an entry by a given path.
//
// If the entry isn't found, (nil, false) is returned.
func (e Entries) Find(path string) (*Entry, bool) {
	length := len(e)
	index := sort.Search(length, func(i int) bool {
		return e[i].Path == path
	})

	if index < length && e[index].Path == path {
		return &e[index], true
	}

	return nil, false
}
