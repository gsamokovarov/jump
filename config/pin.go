package config

import (
	"github.com/gsamokovarov/jump/config/atom"
	"github.com/gsamokovarov/jump/config/jsonio"
)

// ReadPins tries to the directory from a pinned search term.
//
// If no search pinned search term is found.
func (c *fileConfig) ReadPins() (map[string]string, error) {
	file, err := atom.Open(c.Pins)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	pins := map[string]string{}

	return pins, jsonio.Decode(file, &pins)
}

// FindPin tries to find a directory for a pinned search term.
//
// If no search pinned search term is found.
func (c *fileConfig) FindPin(term string) (dir string, found bool) {
	file, err := atom.Open(c.Pins)
	if err != nil {
		return
	}
	defer file.Close()

	pins, err := c.ReadPins()
	if err != nil {
		return
	}

	dir, found = pins[term]

	return
}

// WritePin saves a pinned search term into a file.
func (c *fileConfig) WritePin(pin, value string) error {
	file, err := atom.Open(c.Pins)
	if err != nil {
		return err
	}
	defer file.Close()

	pins, err := c.ReadPins()
	if err != nil {
		return err
	}

	pins[pin] = value

	return jsonio.Encode(file, pins)
}

// RemovePin removes a pinned search term from a file.
func (c *fileConfig) RemovePin(pin string) error {
	file, err := atom.Open(c.Pins)
	if err != nil {
		return err
	}
	defer file.Close()

	pins := map[string]string{}
	if err := jsonio.Decode(file, &pins); err != nil {
		return err
	}

	delete(pins, pin)

	return jsonio.Encode(file, pins)
}
