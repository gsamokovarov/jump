package config

import (
	"github.com/gsamokovarov/jump/config/jsonio"
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
	if err := jsonio.Decode(pinsFile, &pins); err == nil {
		dir, found = pins[term]
	}

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
	if err := jsonio.Decode(pinsFile, &pins); err != nil {
		return err
	}

	pins[pin] = value

	return jsonio.Encode(pinsFile, pins)
}
