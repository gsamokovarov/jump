package config

// ReadPins reads the pinned search terms and their directories.
//
// If the pins file is missing or empty, the returned pins are empty.
func (c *fileConfig) ReadPins() (map[string]string, error) {
	pins := map[string]string{}

	return pins, readFile(c.Pins, &pins)
}

// FindPin tries to find a directory for a pinned search term.
func (c *fileConfig) FindPin(term string) (dir string, found bool) {
	pins, err := c.ReadPins()
	if err != nil {
		return
	}

	dir, found = pins[term]

	return
}

// WritePin saves a pinned search term into a file.
func (c *fileConfig) WritePin(pin, value string) error {
	pins, err := c.ReadPins()
	if err != nil {
		return err
	}

	pins[pin] = value

	return writeFile(c.Pins, pins)
}

// RemovePin removes a pinned search term from a file.
func (c *fileConfig) RemovePin(pin string) error {
	pins, err := c.ReadPins()
	if err != nil {
		return err
	}

	delete(pins, pin)

	return writeFile(c.Pins, pins)
}
