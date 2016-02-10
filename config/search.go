package config

import (
	"encoding/json"
	"io/ioutil"
	"syscall"
)

// Search represents a search term used for advancing through the entries of the same
// term.
type Search struct {
	Term  string
	Index int
}

// ReadSearch reads the last saved search entry.
//
// If the last search doesn't exist, a zero value Search is returned.
func (c *Config) ReadSearch() (search Search) {
	searchFile, err := c.searchFile()
	if err != nil {
		return
	}

	syscall.Flock(int(searchFile.Fd()), syscall.LOCK_EX)

	defer searchFile.Close()
	defer syscall.Flock(int(searchFile.Fd()), syscall.LOCK_UN)

	if content, err := ioutil.ReadAll(searchFile); err == nil {
		if err := json.Unmarshal(content, &search); err == nil {
			return
		}
	}

	return
}

// WriteSearch writes the last search entry to the current search entry.
func (c *Config) WriteSearch(term string, index int) error {
	searchFile, err := c.searchFile()
	if err != nil {
		return err
	}

	syscall.Flock(int(searchFile.Fd()), syscall.LOCK_EX)

	defer searchFile.Close()
	defer syscall.Flock(int(searchFile.Fd()), syscall.LOCK_UN)

	jsonContent, err := json.Marshal(&Search{term, index})
	if err != nil {
		return err
	}

	_, ferr := searchFile.Write(jsonContent)
	return ferr
}
