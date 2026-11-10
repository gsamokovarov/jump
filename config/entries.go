package config

import "github.com/gsamokovarov/jump/scoring"

// ReadEntries returns the current entries for the config.
//
// If the scores file is missing or empty, the returned entries are empty.
func (c *fileConfig) ReadEntries() (entries scoring.Entries, err error) {
	err = readFile(c.Scores, &entries)
	return
}

// WriteEntries the input scoring entries to a file.
//
// Sorts the entries before writing them to disk.
func (c *fileConfig) WriteEntries(entries scoring.Entries) error {
	entries.Sort()

	return writeFile(c.Scores, entries)
}
