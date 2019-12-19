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
	term := strings.Join(args.Raw(), osSeparator)

	entry, err := cdEntry(term, conf)
	if errors.Is(err, errNoEntries) {
		cli.Errf(noEntriesMessage)
		return nil
	}
	if err != nil {
		return err
	}

	cli.Outf("%s\n", entry.Path)

	return nil
}

func cdEntry(term string, conf config.Config) (*scoring.Entry, error) {
	entries, err := conf.ReadEntries()
	if err != nil {
		return nil, err
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

			if err := conf.WriteSearch(term, index); err != nil {
				return nil, err
			}

			return entry, nil
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

func init() {
	cli.RegisterCommand("cd", "Fuzzy match a directory to jump to.", cdCmd)
}
