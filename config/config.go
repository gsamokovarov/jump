package config

import (
	"os"
	"path/filepath"

	"github.com/gsamokovarov/jump/scoring"
)

const (
	defaultScoreFile    = "scores.json"
	defaultSearchFile   = "search.json"
	defaultSettingsFile = "settings.json"
	defaultPinsFile     = "pins.json"
	defaultHomeDir      = ".jump"
	defaultXDGDir       = "jump"
)

// Config represents the config directory and all the miscellaneous
// configuration files we can have in there.
type Config interface {
	ReadEntries() (scoring.Entries, error)
	WriteEntries(scoring.Entries) error

	ReadSearch() Search
	WriteSearch(string, int) error

	ReadPins() (map[string]string, error)
	FindPin(string) (string, bool)
	WritePin(string, string) error
	RemovePin(string) error

	ReadSettings() Settings
	WriteSettings(Settings) error
}

type fileConfig struct {
	Dir      string
	Scores   string
	Search   string
	Pins     string
	Settings string
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
		Search:   filepath.Join(dir, defaultSearchFile),
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
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if xdgHome := os.Getenv("XDG_CONFIG_HOME"); xdgHome != "" {
			configDir = filepath.Join(xdgHome, defaultXDGDir)
		}
	}

	return configDir, nil
}
