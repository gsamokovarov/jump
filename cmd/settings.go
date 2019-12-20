package cmd

import (
	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
)

const settingsUsage = `Usage: jump settings --setting[=value]

Jump is opinionated and we would recommend you to stick to the sweet
hand-tuned defaults we have provided after years of research, however,
we provide a few options few settings that may be useful to hand-tune
yourself:

--space (values: slash (default), ignore)

	The calls "j parent child" and "j parent/child" are equivalent by
	default because spaces are treated as OS separators (/ in Unix). You
	can choose to ignore spaces in searches by setting the "spaces" option
	to "ignore":

    jump settings --space=ignore

--preserve (values: false (default), true)

	By default, landing in a directory that is no-longer available on disk
	will cause jump to remove that directory from its database. If a jump
	lands in unmounted drive, the changing of directory will timeout. This
	is why this is turned off (false) by default.

    jump settings --preserve=true

--reset

  Reset jump settings to their default values.

    jump settings --reset
`

func cmdSettings(args cli.Args, conf config.Config) error {
	validOptionsUsed := false

	if args.Has("--space") {
		err := cmdSettingSpace(conf, args.Get("--space", cli.Optional))
		if err != nil {
			return err
		}

		validOptionsUsed = true
	}

	if args.Has("--preserve") {
		err := cmdSettingPreserve(conf, args.Get("--preserve", cli.Optional))
		if err != nil {
			return err
		}

		validOptionsUsed = true
	}

	if args.Has("--reset") {
		err := cmdSettingReset(conf)
		if err != nil {
			return err
		}

		validOptionsUsed = true
	}

	if !validOptionsUsed {
		cli.Exitf(1, settingsUsage)
	}

	return nil
}

func cmdSettingSpace(conf config.Config, value string) error {
	settings := conf.ReadSettings()

	switch value {
	case "slash":
		settings.Space = config.SpaceSlash
	case "ignore":
		settings.Space = config.SpaceIgnore
	case cli.Optional:
		cli.Outf("--space=%v", settings.Space)
		return nil
	default:
		cli.Exitf(1, "Invalid value: %s; valid values: slash, ignore", value)
		return nil
	}

	return conf.WriteSettings(settings)
}

func cmdSettingPreserve(conf config.Config, value string) error {
	settings := conf.ReadSettings()
	switch value {
	case "true":
		settings.Preserve = true
	case "false":
		settings.Preserve = false
	case cli.Optional:
		cli.Outf("--preserve=%v", settings.Preserve)
		return nil
	default:
		cli.Exitf(1, "Invalid value: %s; valid values: slash, ignore", value)
		return nil
	}

	return conf.WriteSettings(settings)
}

func cmdSettingReset(conf config.Config) error {
	// The zero value of config.Settings is actually the default settings. Make
	// sure to keep it that way, because it's a nice constraint to have.
	var defaultSettings config.Settings

	return conf.WriteSettings(defaultSettings)
}

func init() {
	cli.RegisterCommand("settings", "Configure jump settings.", cmdSettings)
}
