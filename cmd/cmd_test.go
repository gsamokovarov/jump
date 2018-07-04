package cmd

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

var td string

type testConfig struct {
	Entries scoring.Entries
	Search  config.Search
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

func capture(stream **os.File, fn func()) string {
	rescue := *stream
	r, w, _ := os.Pipe()

	*stream = w
	defer func() {
		*stream = rescue
	}()

	fn()

	w.Close()
	out, _ := ioutil.ReadAll(r)

	return string(out)
}

func inside(dir string, fn func()) {
	previousCwd, err := os.Getwd()
	if err != nil {
		return
	}

	if os.Chdir(dir) != nil {
		return
	}

	fn()

	os.Chdir(previousCwd)
}

func init() {
	cwd, _ := os.Getwd()
	td = path.Join(cwd, "testdata")
}
