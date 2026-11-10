package config

import (
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gsamokovarov/jump/scoring"
)

// Config represents the config directory and all the miscellaneous
// configuration files we can have in there.
type Config interface {
	ReadEntries() (scoring.Entries, error)
	WriteEntries(scoring.Entries) error

	ReadSearch() Search

	ReadPins() (map[string]string, error)
	FindPin(string) (string, bool)
	WritePin(string, string) error
	RemovePin(string) error

	ReadSettings() Settings
	WriteSettings(Settings) error
}

// SpaceSetting controls how space in a jump term is treated.
type SpaceSetting int

const (
	// SpaceSlash treats `j receipt app` like `j receipt/app`, which issues a
	// deep search.
	SpaceSlash SpaceSetting = iota

	// SpaceIgnore treats `j receipt app` like `j receiptapp`, which mimics z
	// space behavior, because of the fuzzy matching offered in jump.
	SpaceIgnore
)

// String implements the fmt.Stringer interface.
func (s SpaceSetting) String() string {
	switch s {
	case SpaceSlash:
		return "slash"
	case SpaceIgnore:
		return "ignore"
	default:
		return ""
	}
}

// Settings represents user configurable behaviours for jump.
//
// Keep the default setting values the zero value of their type.
type Settings struct {
	Space    SpaceSetting
	Preserve bool
}

// Search represents a search term used for advancing through the entries of
// the same term.
type Search struct {
	Term  string
	Index int
}

// Setup setups the config folder from a directory path.
//
// If the directories don't already exists, they are created and if the score
// file is present, it is loaded.
func Setup(dir string) (Config, error) {
	// We get the directory check for free form os.MkdirAll.
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	return &fileConfig{
		Dir:      dir,
		Scores:   filepath.Join(dir, defaultScoreFile),
		Pins:     filepath.Join(dir, defaultPinsFile),
		Settings: filepath.Join(dir, defaultSettingsFile),
	}, nil
}

// SetupDefault setups the config folder from a directory path.
//
// If the directory path is an empty string, the path is automatically guessed.
func SetupDefault(dir string) (Config, error) {
	dir, err := findConfigDir(dir)
	if err != nil {
		return nil, err
	}

	return Setup(dir)
}

const (
	defaultScoreFile    = "scores.json"
	defaultSettingsFile = "settings.json"
	defaultPinsFile     = "pins.json"
	defaultHomeDir      = ".jump"
	defaultXDGDir       = "jump"
)

type fileConfig struct {
	Dir      string
	Scores   string
	Pins     string
	Settings string
}

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

// ReadSearch reads the search to retry.
//
// A process cannot change the environment of its shell, so the shell
// integration remembers the last search itself. It passes the term in
// $JUMP_TERM and the number of the retry in $JUMP_TRY, only when j is called
// without arguments. If they are not set, a zero value Search is returned.
func (c *fileConfig) ReadSearch() Search {
	index, err := strconv.Atoi(os.Getenv("JUMP_TRY"))
	if err != nil || index < 0 {
		index = 0
	}

	return Search{os.Getenv("JUMP_TERM"), index}
}

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

// ReadSettings reads the current user settings.
//
// If the settings don't exist, the zero value Settings is returned.
func (c *fileConfig) ReadSettings() (settings Settings) {
	readFile(c.Settings, &settings)

	return
}

// WriteSettings preserves the user settings.
func (c *fileConfig) WriteSettings(settings Settings) error {
	return writeFile(c.Settings, settings)
}

// findConfigDir finds the jump configuration directory.
//
// The search algorithm tries the directories in order:
//
// - $JUMP_HOME (if given)
// - $HOME/.jump (if already exists)
// - $XDG_CONFIG_HOME/jump (prefer for new installs)
//
// We're moving towards XDG, but for existing installs or non-XDG supported
// systems, the ~/.jump dir will be used.
func findConfigDir(dir string) (string, error) {
	if dir != "" {
		return dir, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return dir, err
	}

	configDir := filepath.Join(home, defaultHomeDir)
	if _, err := os.Stat(configDir); errors.Is(err, fs.ErrNotExist) {
		if xdgHome := os.Getenv("XDG_CONFIG_HOME"); xdgHome != "" {
			configDir = filepath.Join(xdgHome, defaultXDGDir)
		}
	}

	return configDir, nil
}

// readFile decodes the JSON in the named file into v.
//
// A missing or an empty file leaves v as it is.
func readFile(name string, v any) error {
	file, err := os.Open(name)
	if errors.Is(err, fs.ErrNotExist) {
		return nil
	} else if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	for {
		if err := decoder.Decode(v); errors.Is(err, io.EOF) {
			return nil
		} else if err != nil {
			return err
		}
	}
}

// writeFile encodes v as JSON into the named file.
//
// The value is written to a temporary file that is then renamed over the
// named one, so a concurrent reader never sees a half-written file.
func writeFile(name string, v any) error {
	tmp, err := os.CreateTemp(filepath.Dir(name), ".jump")
	if err != nil {
		return err
	}

	err = json.NewEncoder(tmp).Encode(v)
	if closeErr := tmp.Close(); err == nil {
		err = closeErr
	}
	if err == nil {
		err = os.Rename(tmp.Name(), name)
	}
	if err != nil {
		os.Remove(tmp.Name())
	}

	return err
}
