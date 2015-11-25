package config

import (
	"encoding/json"
	"io/ioutil"
)

// Search represents a search term used for advancing through the entries of the same
// term.
type Search struct {
	Term  string
	Index int
}

// Reads the last saved search entry.
//
// If the last search doesn't exist, a zero value Search is returned.
func (c *Config) ReadSearch() Search {
	var search Search

	if searchFile, err := c.searchFile(); err == nil {
		defer searchFile.Close()

		if content, err := ioutil.ReadAll(searchFile); err == nil {
			if err := json.Unmarshal(content, &search); err == nil {
				return search
			}
		}
	}

	return search
}

// Writes the last search entry to the current search entry.
func (c *Config) WriteSearch(term string, index int) error {
	jsonContent, err := json.Marshal(&Search{term, index})
	if err != nil {
		return err
	}

	return ioutil.WriteFile(c.Search, jsonContent, 0644)
}
