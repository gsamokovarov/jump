package cmd

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/fuzzy"
	"github.com/gsamokovarov/jump/scoring"
)

var errNoEntries = errors.New("no entries in the database")

const noEntriesMessage = `Jump's database is empty. This could mean:

1. You are running jump for the first time. Have you integrated jump with your
   shell? Run the following command for help:
	
       $ jump shell

   If you have run the integration, enter a few directories in a new shell to
   populate the database.

   Are you coming from autojump or z? You can import their existing scoring
   databases into jump with:

       $ jump import

2. The database is corrupted. You can open an issue on jump's issue tracker
   at https://github.com/gsamokovarov/jump/issues.
`

func cdCmd(args cli.Args, conf config.Config) error {
	var entry *scoring.Entry
	var err error
	var term, baseDir string

	if filepath.IsAbs(args.First()) && dirIsAccessible(args.First()) && len(args) > 1 {
		baseDir = args.First()
		term = termFromArgs(args.Rest(), conf)
	} else {
		term = termFromArgs(args, conf)
	}

	entry, err = cdEntry(term, baseDir, conf)
	if errors.Is(err, errNoEntries) {
		cli.Errf(noEntriesMessage)
		return nil
	} else if err != nil {
		return err
	}

	cli.Outf("%s\n", entry.Path)

	return nil
}

func cdEntry(term, baseDir string, conf config.Config) (*scoring.Entry, error) {
	entries, err := conf.ReadEntries()
	if err != nil {
		return nil, err
	}
	if baseDir != "" {
		entries = entries.Under(baseDir)
	}

	// If an auto-completion triggered a full path, just go there.
	if filepath.IsAbs(term) {
		return scoring.NewEntry(term), nil
	}

	index, search := 0, conf.ReadSearch()

	// If we happen to match the last term, e.g. j is called with no
	// arguments then jump to the previous search.
	if term == "" {
		term, index = search.Term, search.Index+1
	} else {
		// If there is a term given, first see if there is a pin for
		// it and if so, always follow it.
		if dir, found := conf.FindPin(term); found {
			// Except if we land on the current directory again. Then
			// ignore the term.
			if !fwdPathIsCwd(dir) {
				return scoring.NewEntry(dir), nil
			}
		}
	}

	// If the directory exists under the cwd, go there.
	cwd, err := os.Getwd()
	if termIsRelative(term) && err == nil {
		relativeDir := filepath.Join(cwd, term)
		if dirIsAccessible(relativeDir) {
			return scoring.NewEntry(relativeDir), nil
		}
	}

	settings := conf.ReadSettings()
	fuzzyEntries := scoring.NewFuzzyEntries(entries, term)

	// Prefer an exact match if it's in a reasonable proximity of the best
	// match. Useful for jumping to (2...4) letter directories, which you
	// may just type in their exact form anyway.
	index = exactMatchInProximity(fuzzyEntries, term, index)

	for {
		if entry, ok := fuzzyEntries.Select(index); ok {
			// Remove the entries that no longer exists.
			if _, err := os.Stat(entry.Path); os.IsNotExist(err) && !settings.Preserve {
				entries.Remove(entry.Path)
				conf.WriteEntries(entries)

				index++
				continue
			}

			// Jump to the next entry, if the jump is going to land on the
			// current directory.
			if fwdPathIsCwd(entry.Path) {
				index++
				continue
			}

			return entry, conf.WriteSearch(term, index)
		}

		// If we're given a base directory, and there is no match, go to the base.
		if baseDir != "" {
			return scoring.NewEntry(baseDir), nil
		}

		break
	}

	return nil, errNoEntries
}

const exactMatchProximity = 5
const exactMatchLenRequirement = 5

func exactMatchInProximity(entries *scoring.FuzzyEntries, term string, offset int) int {
	norm := fuzzy.NewNormalizer(filepath.Base(term))
	normalizedTerm := norm.NormalizeTerm()

	if len(normalizedTerm) < exactMatchLenRequirement {
		return offset
	}

	for index := offset; index <= offset+exactMatchProximity; index++ {
		entry, ok := entries.Select(index)
		if !ok {
			continue
		}

		// Take only the base part, if you wanna do a deep search like Dev/nes.
		basePath := filepath.Base(norm.NormalizePath(entry.Path))
		if basePath == normalizedTerm {
			return index
		}
	}

	return offset
}

func fwdPathIsCwd(path string) bool {
	cwd, err := os.Getwd()
	if err != nil {
		return false
	}

	fwdPath, err := filepath.EvalSymlinks(path)
	if err != nil {
		return false
	}

	return fwdPath == cwd
}

func dirIsAccessible(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fi.IsDir()
}

func termIsRelative(term string) bool {
	return strings.Contains(term, "/")
}

func init() {
	cli.RegisterCommand("cd", "Fuzzy match a directory to jump to.", cdCmd)
}
