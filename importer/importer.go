package importer

import (
	"errors"
	"os"

	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

// ErrNotFound is an error returned when the importing tool is not found.
var ErrNotFound = errors.New("importer: cannot find autojump or z data files")

// Callback is called on every import.
type Callback func(*scoring.Entry)

// Call calls the callback without the need of nil checks.
func (fn Callback) Call(entry *scoring.Entry) {
	if fn != nil {
		fn(entry)
	}
}

// Importer imports a configuration from an external tool into jump.
type Importer interface {
	Import(fn Callback) error
}

// Guess tries to guess the importer to use based on a hint.
func Guess(hint string, conf config.Config) Importer {
	var imp Importer

	switch hint {
	case "autojump":
		imp = Autojump(conf)
	case "z":
		imp = Z(conf)
	default:
		// First try z, then try autojump.
		imp = multiImporter{Z(conf), Autojump(conf)}
	}

	return imp
}

func readConfig(paths []string) (string, error) {
	path, err := findPath(paths)
	if err != nil {
		return "", err
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func findPath(paths []string) (string, error) {
	for _, path := range paths {
		path = os.ExpandEnv(path)

		if _, err := os.Stat(path); !os.IsNotExist(err) {
			return path, nil
		}
	}

	return "", ErrNotFound
}
