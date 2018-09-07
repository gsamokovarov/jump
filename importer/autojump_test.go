package importer

import (
	p "path"
	"testing"

	"github.com/gsamokovarov/assert"
	"github.com/gsamokovarov/jump/config"
)

func TestAutojump(t *testing.T) {
	conf := &config.Testing{}
	configPath := p.Join(td, "autojump.txt")

	imp := Autojump(conf, configPath)

	err := imp.Import(nil)
	assert.Nil(t, err)

	assert.
		Len(t, 6, conf.Entries).
		// 0
		Equal(t, "/Users/genadi/Development/mock_last_status", conf.Entries[0].Path).
		Equal(t, 14, conf.Entries[0].Score.Weight).
		// 1
		Equal(t, "/Users/genadi/.go/src/github.com/gsamokovarov/jump", conf.Entries[1].Path).
		Equal(t, 14, conf.Entries[1].Score.Weight).
		// 2
		Equal(t, "/Users/genadi/Development/gloat", conf.Entries[2].Path).
		Equal(t, 20, conf.Entries[2].Score.Weight).
		// 3
		Equal(t, "/Users/genadi/Development", conf.Entries[3].Path).
		Equal(t, 33, conf.Entries[3].Score.Weight).
		// 4
		Equal(t, "/Users/genadi/Development/jump", conf.Entries[4].Path).
		Equal(t, 39, conf.Entries[4].Score.Weight).
		// 5
		Equal(t, "/usr/local/Cellar/autojump", conf.Entries[5].Path).
		Equal(t, 44, conf.Entries[5].Score.Weight)

	for i, j := 0, 1; i < len(conf.Entries)-1; i, j = i+1, j+1 {
		assert.True(t, conf.Entries[i].CalculateScore() <= conf.Entries[j].CalculateScore())
	}
}
