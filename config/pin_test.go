package config

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

func TestPins(t *testing.T) {
	t.Run("save and read last search term from a file", func(t *testing.T) {
		conf := tempConfig(t)

		err := conf.WritePin("wc", "/test/web-console")
		assert.Nil(t, err)

		pin, found := conf.FindPin("wc")
		assert.True(t, found)

		assert.Equal(t, "/test/web-console", pin)
	})
}
