package cmd

import (
	"os"
	p "path"
	"strings"
	"testing"

	"github.com/gsamokovarov/assert"
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	s "github.com/gsamokovarov/jump/scoring"
)

func Test_cdSettings(t *testing.T) {
	t.Run("jump settings --space=ignore", func(t *testing.T) {
		conf := tempConfig(t)

		capture(&os.Stdout, func() {
			assert.Nil(t, cmdSettings(cli.Args{"settings", "--space=ignore"}, conf))
		})

		output := capture(&os.Stdout, func() {
			assert.Nil(t, cmdSettings(cli.Args{"settings", "--space"}, conf))
		})

		assert.True(t, strings.Contains(output, "--space=ignore"))
	})

	t.Run("jump settings --space=ignore ignores whitespace in arguments", func(t *testing.T) {
		conf := &config.InMemory{
			Entries: s.Entries{
				&s.Entry{p.Join(td, "web-console"), &s.Score{Weight: 100, Age: s.Now}},
				&s.Entry{p.Join(td, "/client/website"), &s.Score{Weight: 90, Age: s.Now}},
			},
			Settings: config.Settings{
				Space: config.SpaceIgnore,
			},
		}

		output := capture(&os.Stdout, func() {
			assert.Nil(t, cdCmd(cli.Args{"web", "console"}, conf))
		})

		assert.True(t, strings.Contains(output, p.Join(td, "web-console")))
	})

	t.Run("jump settings --preserve=true", func(t *testing.T) {
		conf := tempConfig(t)

		capture(&os.Stdout, func() {
			assert.Nil(t, cmdSettings(cli.Args{"settings", "--preserve=true"}, conf))
		})

		output := capture(&os.Stdout, func() {
			assert.Nil(t, cmdSettings(cli.Args{"settings", "--preserve"}, conf))
		})

		assert.True(t, strings.Contains(output, "--preserve=true"))
	})

	t.Run("jump settings --preserve=true does not delete missing files", func(t *testing.T) {
		conf := &config.InMemory{
			Entries: s.Entries{
				&s.Entry{p.Join(td, "webview"), &s.Score{Weight: 100, Age: s.Now}},
				&s.Entry{p.Join(td, "/client/website"), &s.Score{Weight: 90, Age: s.Now}},
			},
			Settings: config.Settings{
				Preserve: true,
			},
		}

		output := capture(&os.Stdout, func() {
			assert.Nil(t, cdCmd(cli.Args{"webview"}, conf))
		})

		assert.True(t, strings.Contains(output, p.Join(td, "webview")))
	})

	t.Run("jump settings --reset resets the settings to their default values", func(t *testing.T) {
		conf := &config.InMemory{
			Entries: s.Entries{
				&s.Entry{p.Join(td, "webview"), &s.Score{Weight: 100, Age: s.Now}},
				&s.Entry{p.Join(td, "/client/website"), &s.Score{Weight: 90, Age: s.Now}},
			},
			Settings: config.Settings{
				Space:    config.SpaceIgnore,
				Preserve: true,
			},
		}

		capture(&os.Stdout, func() {
			assert.Nil(t, cmdSettings(cli.Args{"--reset"}, conf))
		})

		assert.Equal(t, config.SpaceSlash, conf.Settings.Space)
		assert.Equal(t, false, conf.Settings.Preserve)
	})

}
