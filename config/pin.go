package config

import (
	"encoding/json"
	"io"
)

// FindPin tries to the directory from a pinned search term.
//
// If no search pinned search term is found.
func (c *Config) FindPin(term string) (dir string, found bool) {
	pinsFile, err := c.pinsFile()
	if err != nil {
		return
	}

	defer closeLockedFile(pinsFile)

	pins := map[string]string{}

	decoder := json.NewDecoder(pinsFile)
	if err := decoder.Decode(&pins); err != nil && err != io.EOF {
		return
	}

	dir, found = pins[term]
	return
}

// WritePin saves a pinned search term into a file.
func (c *Config) WritePin(pin, value string) error {
	pinsFile, err := c.pinsFile()
	if err != nil {
		return err
	}

	defer closeLockedFile(pinsFile)

	pins := map[string]string{}

	decoder := json.NewDecoder(pinsFile)
	if err := decoder.Decode(&pins); err != nil && err != io.EOF {
		return err
	}

	pins[pin] = value

	if _, err := pinsFile.Seek(0, io.SeekStart); err != nil {
		return err
	}

	if err := pinsFile.Truncate(0); err != nil {
		return err
	}

	// Seeking to the beginning is important here, so we don't end up writing zero by
	encoder := json.NewEncoder(pinsFile)

	return encoder.Encode(pins)
}
