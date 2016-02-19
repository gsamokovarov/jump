package config

import (
	"encoding/json"
	"io"

	"github.com/gsamokovarov/jump/scoring"
)

// ReadEntries returns the current entries for the config.
//
// If the scores file is empty, the returned entries are empty.
func (c *Config) ReadEntries() (*scoring.Entries, error) {
	var entries scoring.Entries

	scoresFile, err := c.scoresFile()
	if err != nil {
		return &entries, nil
	}

	defer closeLockedFile(scoresFile)

	decoder := json.NewDecoder(scoresFile)
	for {
		if err := decoder.Decode(&entries); err == io.EOF {
			break
		} else if err != nil {
			return &entries, err
		}
	}

	return &entries, nil
}

// WriteEntries the input scoring entries to a file.
//
// Sorts the entries before writing them to disk.
func (c *Config) WriteEntries(entries *scoring.Entries) error {
	scoresFile, err := c.scoresFile()
	if err != nil {
		return err
	}

	defer closeLockedFile(scoresFile)

	if err := scoresFile.Truncate(0); err != nil {
		return err
	}

	entries.Sort()
	encoder := json.NewEncoder(scoresFile)

	return encoder.Encode(entries)
}
