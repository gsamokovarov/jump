package config

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

func TestSettings(t *testing.T) {
	t.Run("save and read settings from a file", func(t *testing.T) {
		conf := tempConfig(t)

		err := conf.WriteSettings(Settings{
			Space:    SpaceIgnore,
			Preserve: true,
		})
		assert.Nil(t, err)

		s := conf.ReadSettings()

		assert.Equal(t, SpaceIgnore, s.Space)
		assert.Equal(t, true, s.Preserve)
	})
}
