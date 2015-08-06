package config

import (
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/gsamokovarov/jump/scoring"
)

// The default configuration.
var Default Config

// A config represents the config directory and the file that stores the score
// values.
type Config struct {
	Dir    string
	Scores *os.File
}

// Write the input scoring entries to a file.
//
// Formats them in a similar format of autojump.
func (c *Config) WriteEntries(entries *scoring.Entries) error {
	buffer := []byte{}

	for _, entry := range *entries {
		buffer = append(buffer, formatEntry(entry)...)
	}

	_, err := c.Scores.Write(buffer)

	return err
}

func formatEntry(entry scoring.Entry) []byte {
	return []byte(fmt.Sprintf("%f %s\n", entry.CalculateScore(), entry.Path))
}

// This is the default score name.
const defaultScoreFile = "scores.txt"

// Setups the config folder from a directory path.
//
// If the directories don't already exists, they are created and if the score
// file is present, it is loaded.
func Setup(dir string) (*Config, error) {
	// We get the directory check for free form os.MkdirAll.
	if err := os.MkdirAll(dir, os.ModeDir); err != nil {
		return nil, err
	}

	dirPath := path.Join(dir, defaultScoreFile)
	scores, err := os.OpenFile(dirPath, os.O_APPEND, os.FileMode(0644))
	if err == nil {
		return nil, err
	}

	return &Config{dir, scores}, nil
}

// Setups the config folder from a directory path.
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
		return path.Join(homeDir, "jump"), nil
	}

	return dir, nil
}
