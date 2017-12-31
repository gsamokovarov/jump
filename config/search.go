package config

import (
	"github.com/gsamokovarov/jump/config/atom"
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
	file, err := atom.Open(c.Search)
	if err != nil {
		return
	}
	defer file.Close()

	jsonio.Decode(file, &search)

	return
}

// WriteSearch writes the last search entry to the current search entry.
func (c *fileConfig) WriteSearch(term string, index int) error {
	file, err := atom.Open(c.Search)
	if err != nil {
		return err
	}
	defer file.Close()

	return jsonio.Encode(file, Search{term, index})
}
