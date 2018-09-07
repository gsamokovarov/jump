package importer

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

var autojumpDefaultConfigPaths = []string{
	"$HOME/.local/share/autojump/autojump.txt",
	"$HOME/Library/autojump/autojump.txt",
}

// Autojump is an importer for the autojump tool.
func Autojump(conf config.Config, configPaths ...string) Importer {
	if len(configPaths) == 0 {
		configPaths = autojumpDefaultConfigPaths
	}

	return &autojump{
		config:      conf,
		configPaths: configPaths,
	}
}

type autojump struct {
	config      config.Config
	configPaths []string
}

func (i *autojump) Import(fn Callback) error {
	autojumpEntries, err := i.parseConfig()
	if err != nil {
		return err
	}

	jumpEntries, err := i.config.ReadEntries()
	if err != nil {
		return err
	}

	for _, entry := range autojumpEntries {
		if _, found := jumpEntries.Find(entry.Path); found {
			continue
		}

		fn.Call(entry)

		jumpEntries = append(jumpEntries, entry)
	}

	return i.config.WriteEntries(jumpEntries)
}

func (i *autojump) parseConfig() (scoring.Entries, error) {
	content, err := readConfig(i.configPaths)
	if err != nil {
		return nil, err
	}

	var entries scoring.Entries

	for _, line := range strings.Split(content, "\n") {
		entry, err := i.newEntryFromLine(line)
		if err == io.EOF {
			continue
		}
		if err != nil {
			return nil, err
		}

		if _, found := entries.Find(entry.Path); found {
			continue
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

func (i *autojump) newEntryFromLine(line string) (*scoring.Entry, error) {
	if line == "" {
		return nil, io.EOF
	}

	parts := strings.Split(line, "\t")
	if len(parts) != 2 {
		return nil, fmt.Errorf("importer: cannot parse entry: %s", line)
	}

	path := parts[1]
	weight, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return nil, err
	}

	return &scoring.Entry{
		Path: path,
		Score: &scoring.Score{
			Weight: int64(math.Round(weight)),
			Age:    scoring.Now,
		},
	}, nil
}
