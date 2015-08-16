package cmd

import (
	"path/filepath"
	"sort"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/lcs"
	"github.com/gsamokovarov/jump/scoring"
)

type fuzzyEntries struct {
	entries scoring.Entries
	target  string
}

func (fe fuzzyEntries) Len() int {
	return len(fe.entries)
}

func (fe fuzzyEntries) Swap(i, j int) {
	fe.entries[i], fe.entries[j] = fe.entries[j], fe.entries[i]
}

func (fe fuzzyEntries) Less(i, j int) bool {
	iPath := filepath.Base(fe.entries[i].Path)
	jPath := filepath.Base(fe.entries[j].Path)

	return lcs.Length(iPath, fe.target) >= lcs.Length(jPath, fe.target)
}

func (fe fuzzyEntries) Choose() (entry *scoring.Entry, empty bool) {
	sort.Sort(fe)

	if fe.Len() == 0 {
		return nil, true
	}

	return &fe.entries[0], false
}

func cdCmd(args cli.Args, conf *config.Config) {
	if len(args) == 0 {
		return
	}

	rawEntries, err := conf.ReadEntries()
	if err != nil {
		cli.Exitf(1, "%s\n", err)
	}

	entries := fuzzyEntries{rawEntries, args.CommandName()}
	if entry, empty := entries.Choose(); !empty {
		cli.Outf("%s\n", entry.Path)
	}
}

func init() {
	cli.RegisterCommand("cd", "Fuzzy match a directory to jump to.", cdCmd)
}
