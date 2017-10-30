package config

import (
	"github.com/gsamokovarov/jump/config/jsonio"
	"github.com/gsamokovarov/jump/scoring"
)

// ReadEntries returns the current entries for the config.
//
// If the scores file is empty, the returned entries are empty.
func (c *fileConfig) ReadEntries() (entries scoring.Entries, err error) {
	scoresFile, err := createOrOpenLockedFile(c.Scores)
	if err != nil {
		return
	}
	defer closeLockedFile(scoresFile)

	if err = jsonio.Decode(scoresFile, &entries); err != nil {
		return
	}

	return
}

// WriteEntries the input scoring entries to a file.
//
// Sorts the entries before writing them to disk.
func (c *fileConfig) WriteEntries(entries scoring.Entries) error {
	scoresFile, err := createOrOpenLockedFile(c.Scores)
	if err != nil {
		return err
	}
	defer closeLockedFile(scoresFile)

	entries.Sort()

	return jsonio.Encode(scoresFile, entries)
}
