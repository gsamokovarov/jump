package importer

import (
	p "path"
	"testing"
	"time"

	"github.com/gsamokovarov/assert"
	"github.com/gsamokovarov/jump/config"
)

func TestZ(t *testing.T) {
	conf := &config.InMemory{}
	configPath := p.Join(td, "z")

	imp := Z(conf, configPath)

	err := imp.Import(nil)
	assert.Nil(t, err)

	assert.
		Len(t, 4, conf.Entries).
		// 0
		Equal(t, "/Users/genadi/.go/src/github.com/gsamokovarov/jump", conf.Entries[0].Path).
		Equal(t, 1, conf.Entries[0].Score.Weight).
		Equal(t, time.Unix(1536272492, 0), conf.Entries[0].Score.Age).
		// 1
		Equal(t, "/Users/genadi/Development/masse", conf.Entries[1].Path).
		Equal(t, 1, conf.Entries[1].Score.Weight).
		Equal(t, time.Unix(1536272502, 0), conf.Entries[1].Score.Age).
		// 2
		Equal(t, "/Users/genadi/Development", conf.Entries[2].Path).
		Equal(t, 3, conf.Entries[2].Score.Weight).
		Equal(t, time.Unix(1536272506, 0), conf.Entries[2].Score.Age).
		// 3
		Equal(t, "/Users/genadi/Development/hack", conf.Entries[3].Path).
		Equal(t, 11, conf.Entries[3].Score.Weight).
		Equal(t, time.Unix(1536272816, 0), conf.Entries[3].Score.Age)

	for i, j := 0, 1; i < len(conf.Entries)-1; i, j = i+1, j+1 {
		assert.True(t, conf.Entries[i].CalculateScore() <= conf.Entries[j].CalculateScore())
	}
}

func TestZAged(t *testing.T) {
	conf := &config.InMemory{}
	configPath := p.Join(td, "z-aged.txt")

	imp := Z(conf, configPath)

	err := imp.Import(nil)
	assert.Nil(t, err)

	assert.
		Len(t, 3, conf.Entries).
		// 0
		Equal(t, "/var/log", conf.Entries[0].Path).
		Equal(t, 79, conf.Entries[0].Score.Weight).
		Equal(t, time.Unix(1553005178, 0), conf.Entries[0].Score.Age).
		// 1
		Equal(t, "/", conf.Entries[1].Path).
		Equal(t, 90, conf.Entries[1].Score.Weight).
		Equal(t, time.Unix(1553005180, 0), conf.Entries[1].Score.Age).
		// 2
		Equal(t, "/home", conf.Entries[2].Path).
		Equal(t, 8658, conf.Entries[2].Score.Weight).
		Equal(t, time.Unix(1553005185, 0), conf.Entries[2].Score.Age)

	for i, j := 0, 1; i < len(conf.Entries)-1; i, j = i+1, j+1 {
		assert.True(t, conf.Entries[i].CalculateScore() <= conf.Entries[j].CalculateScore())
	}
}
