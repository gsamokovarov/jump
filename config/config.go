package config

import (
	"os"
	"os/user"
	"path/filepath"
)

const (
	defaultScoreFile  = "scores.json"
	defaultSearchFile = "search.json"
	defaultDirName    = ".jump"
)

// Config represents the config directory and all the misc. configuration files
// we can have in there.
type Config struct {
	Dir    string
	Scores string
	Search string
}

// Returns a file object for the saved scores path.
func (c *Config) scoresFile() (*os.File, error) {
	return createOrOpenLockedFile(c.Scores)
}

// Returns a file object for the saved term path.
func (c *Config) searchFile() (*os.File, error) {
	return createOrOpenLockedFile(c.Search)
}

// Setup setups the config folder from a directory path.
//
// If the directories don't already exists, they are created and if the score
// file is present, it is loaded.
func Setup(dir string) (*Config, error) {
	// We get the directory check for free form os.MkdirAll.
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	scores := filepath.Join(dir, defaultScoreFile)
	search := filepath.Join(dir, defaultSearchFile)

	return &Config{dir, scores, search}, nil
}

// SetupDefault setups the config folder from a directory path.
//
// If the directory path is an empty string, the path is automatically guessed.
func SetupDefault(dir string) (*Config, error) {
	dir, err := normalizeDir(dir)
	if err != nil {
		return nil, err
	}

	return Setup(dir)
}

func normalizeDir(dir string) (string, error) {
	if len(dir) == 0 {
		usr, err := user.Current()
		if err != nil {
			return dir, err
		}

		homeDir := usr.HomeDir
		return filepath.Join(homeDir, defaultDirName), nil
	}

	return dir, nil
}
