package config

import (
	"github.com/gsamokovarov/jump/config/jsonio"
)

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
	searchFile, err := createOrOpenLockedFile(c.Search)
	if err != nil {
		return
	}
	defer closeLockedFile(searchFile)

	jsonio.Decode(searchFile, &search)

	return
}

// WriteSearch writes the last search entry to the current search entry.
func (c *fileConfig) WriteSearch(term string, index int) error {
	searchFile, err := createOrOpenLockedFile(c.Search)
	if err != nil {
		return err
	}
	defer closeLockedFile(searchFile)

	return jsonio.Encode(searchFile, Search{term, index})
}
