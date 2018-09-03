package config

import (
	"github.com/gsamokovarov/jump/config/atom"
	"github.com/gsamokovarov/jump/config/jsonio"
	"github.com/gsamokovarov/jump/scoring"
)

// ReadEntries returns the current entries for the config.
//
// If the scores file is empty, the returned entries are empty.
func (c *fileConfig) ReadEntries() (entries scoring.Entries, err error) {
	file, err := atom.Open(c.Scores)
	if err != nil {
		return
	}
	defer file.Close()

	err = jsonio.Decode(file, &entries)
	return
}

// WriteEntries the input scoring entries to a file.
//
// Sorts the entries before writing them to disk.
func (c *fileConfig) WriteEntries(entries scoring.Entries) error {
	file, err := atom.Open(c.Scores)
	if err != nil {
		return err
	}
	defer file.Close()

	entries.Sort()

	return jsonio.Encode(file, entries)
}
