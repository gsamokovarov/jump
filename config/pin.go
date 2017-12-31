package config

import (
	"github.com/gsamokovarov/jump/config/atom"
	"github.com/gsamokovarov/jump/config/jsonio"
)

// FindPin tries to the directory from a pinned search term.
//
// If no search pinned search term is found.
func (c *fileConfig) FindPin(term string) (dir string, found bool) {
	file, err := atom.Open(c.Pins)
	if err != nil {
		return
	}
	defer file.Close()

	pins := map[string]string{}
	if err := jsonio.Decode(file, &pins); err == nil {
		dir, found = pins[term]
	}

	return
}

// WritePin saves a pinned search term into a file.
func (c *fileConfig) WritePin(pin, value string) error {
	file, err := atom.Open(c.Pins)
	if err != nil {
		return err
	}
	defer file.Close()

	pins := map[string]string{}
	if err := jsonio.Decode(file, &pins); err != nil {
		return err
	}

	pins[pin] = value

	return jsonio.Encode(file, pins)
}
