package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"sort"

	"github.com/gsamokovarov/jump/scoring"
)

const (
	defaultScoreFile = "scores.json"
	defaultDirName   = ".jump"
)

// A config represents the config directory and the file that stores the score
// values.
type Config struct {
	Dir    string
	Scores string
}

// Read entries returns the current entries for the config.
//
// If the scores file is empty, the returned entries are empty.
func (c *Config) ReadEntries() (scoring.Entries, error) {
	var entries scoring.Entries

	scoresFile, err := c.scoresFile()
	if err != nil {
		return entries, nil
	}

	defer scoresFile.Close()

	decoder := json.NewDecoder(scoresFile)
	for {
		if err := decoder.Decode(&entries); err == io.EOF {
			break
		} else if err != nil {
			return entries, err
		}
	}

	return entries, nil
}

// Write the input scoring entries to a file.
//
// Sorts the entries before writing them to disk.
func (c *Config) WriteEntries(entries scoring.Entries) error {
	scoresFile, err := c.scoresFile()
	if err != nil {
		return err
	}

	defer scoresFile.Close()

	if err := scoresFile.Truncate(0); err != nil {
		return err
	}

	sort.Sort(entries)
	encoder := json.NewEncoder(scoresFile)

	return encoder.Encode(&entries)
}

// Returns a file object for the saved scores path.
func (c *Config) scoresFile() (*os.File, error) {
	return createOrOpenFile(c.Scores)
}

// Setups the config folder from a directory path.
//
// If the directories don't already exists, they are created and if the score
// file is present, it is loaded.
func Setup(dir string) (*Config, error) {
	// We get the directory check for free form os.MkdirAll.
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	scores := filepath.Join(dir, defaultScoreFile)

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

func createOrOpenFile(name string) (file *os.File, err error) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return os.Create(name)
	}

	return os.OpenFile(name, os.O_RDWR, 0644)
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

func formatEntry(entry scoring.Entry) []byte {
	return []byte(fmt.Sprintf("%f %s\n", entry.CalculateScore(), entry.Path))
}
