package importer

import (
	"errors"
	p "path"
	"testing"

	"github.com/gsamokovarov/assert"
	"github.com/gsamokovarov/jump/config"
)

type failingImporter struct{}

func (failingImporter) Import(Callback) error { return errors.New("importer: failing") }

func Test_multiImporter(t *testing.T) {
	conf := &config.InMemory{}
	autojumpPath := p.Join(td, "autojump.txt")
	zPath := p.Join(td, "z")

	imp := multiImporter{
		Autojump(conf, autojumpPath),
		Z(conf, zPath),
	}

	err := imp.Import(nil)
	assert.Nil(t, err)

	assert.Len(t, 8, conf.Entries)
}

func Test_multiImporter_oneErrored(t *testing.T) {
	conf := &config.InMemory{}
	autojumpPath := p.Join(td, "autojump.txt")

	imp := multiImporter{
		failingImporter{},
		Autojump(conf, autojumpPath),
	}

	err := imp.Import(nil)
	assert.Nil(t, err)

	assert.Len(t, 6, conf.Entries)
}
