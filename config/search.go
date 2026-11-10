package config

// Search represents a search term used for advancing through the entries of
// the same term.
type Search struct {
	Term  string
	Index int
}

// ReadSearch reads the last saved search entry.
//
// If the last search doesn't exist, a zero value Search is returned.
func (c *fileConfig) ReadSearch() (search Search) {
	readFile(c.Search, &search)

	return
}

// WriteSearch writes the last search entry to the current search entry.
func (c *fileConfig) WriteSearch(term string, index int) error {
	return writeFile(c.Search, Search{term, index})
}
