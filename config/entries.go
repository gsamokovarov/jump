package config

import (
	"github.com/gsamokovarov/jump/config/jsonio"
	"github.com/gsamokovarov/jump/scoring"
)

// ReadEntries returns the current entries for the config.
//
// If the scores file is empty, the returned entries are empty.
func (c *fileConfig) ReadEntries() (*scoring.Entries, error) {
	var entries scoring.Entries

	scoresFile, err := createOrOpenLockedFile(c.Scores)
	if err != nil {
		return &entries, nil
	}

	defer closeLockedFile(scoresFile)

	if err := jsonio.Decode(scoresFile, &entries); err != nil {
		return nil, err
	}

	return &entries, nil
}

// WriteEntries the input scoring entries to a file.
//
// Sorts the entries before writing them to disk.
func (c *fileConfig) WriteEntries(entries *scoring.Entries) error {
	scoresFile, err := createOrOpenLockedFile(c.Scores)
	if err != nil {
		return err
	}

	defer closeLockedFile(scoresFile)

	entries.Sort()

	return jsonio.Encode(scoresFile, entries)
}
