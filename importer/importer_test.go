package importer

import (
	"os"
	"path"

	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

var td string

type testConfig struct {
	Entries scoring.Entries
	Search  config.Search
	Pins    map[string]string
	Pin     string
}

func (c *testConfig) ReadEntries() (scoring.Entries, error) {
	return c.Entries, nil
}

func (c *testConfig) WriteEntries(entries scoring.Entries) error {
	c.Entries = entries
	return nil
}

func (c *testConfig) ReadSearch() config.Search {
	return c.Search
}

func (c *testConfig) WriteSearch(term string, index int) error {
	c.Search.Term = term
	c.Search.Index = index
	return nil
}

func (c *testConfig) ReadPins() (map[string]string, error) {
	return c.Pins, nil
}

func (c *testConfig) FindPin(term string) (string, bool) {
	return c.Pin, c.Pin != ""
}

func (c *testConfig) WritePin(_, value string) error {
	c.Pin = value
	return nil
}

func (c *testConfig) RemovePin(term string) error {
	c.Pin = ""
	return nil
}

func init() {
	cwd, _ := os.Getwd()
	td = path.Join(cwd, "testdata")
}
