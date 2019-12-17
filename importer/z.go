package importer

import (
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

var zDefaultConfigPaths = []string{
	"$HOME/.z",
}

// Z is an importer for the z tool.
func Z(conf config.Config, configPaths ...string) Importer {
	if len(configPaths) == 0 {
		configPaths = zDefaultConfigPaths
	}

	return &z{
		config:      conf,
		configPaths: configPaths,
	}
}

type z struct {
	config      config.Config
	configPaths []string
}

func (i *z) Import(fn Callback) error {
	zEntries, err := i.parseConfig()
	if err != nil {
		return err
	}

	jumpEntries, err := i.config.ReadEntries()
	if err != nil {
		return err
	}

	for _, entry := range zEntries {
		if _, found := jumpEntries.Find(entry.Path); found {
			continue
		}

		fn.Call(entry)

		jumpEntries = append(jumpEntries, entry)
	}

	return i.config.WriteEntries(jumpEntries)
}

func (i *z) parseConfig() (scoring.Entries, error) {
	content, err := readConfig(i.configPaths)
	if err != nil {
		return nil, err
	}

	var entries scoring.Entries

	for _, line := range strings.Split(content, "\n") {
		entry, err := i.newEntryFromLine(line)
		if errors.Is(err, io.EOF) {
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

func (i *z) newEntryFromLine(line string) (*scoring.Entry, error) {
	if line == "" {
		return nil, io.EOF
	}

	parts := strings.Split(line, "|")
	if len(parts) != 3 {
		return nil, fmt.Errorf("importer: cannot parse entry: %s", line)
	}

	path := parts[0]
	weight, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return nil, err
	}
	epoch, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		return nil, err
	}

	return &scoring.Entry{
		Path: path,
		Score: &scoring.Score{
			Weight: int64(math.Round(weight)),
			Age:    time.Unix(epoch, 0),
		},
	}, nil
}
