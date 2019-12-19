package cmd

import (
	"os"
	"testing"

	"github.com/gsamokovarov/assert"
	"github.com/gsamokovarov/jump/config"
)

func Test_pinsCmd(t *testing.T) {
	conf := &config.InMemory{
		Pins: map[string]string{
			"r": "/home/user/projects/rails",
		},
	}

	output := capture(&os.Stdout, func() {
		assert.Nil(t, pinsCmd(nil, conf))
	})

	assert.Equal(t, "r\t/home/user/projects/rails\n", output)
}
