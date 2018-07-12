package cmd

import (
	"os"
	"testing"

	"github.com/gsamokovarov/assert"
)

func Test_pinsCmd(t *testing.T) {
	conf := &testConfig{
		Pins: map[string]string{
			"r": "/home/user/projects/rails",
		},
	}

	output := capture(&os.Stdout, func() {
		assert.Nil(t, pinsCmd(nil, conf))
	})

	assert.Equal(t, "r\t/home/user/projects/rails\n", output)
}
