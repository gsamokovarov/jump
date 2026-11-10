package config

// SpaceSetting controls how space in a jump term is treated.
type SpaceSetting int

const (
	// SpaceSlash treats `j receipt app` like `j receipt/app`, which issues a
	// deep search.
	SpaceSlash SpaceSetting = iota

	// SpaceIgnore treats `j receipt app` like `j receiptapp`, which mimics z
	// space behavior, because of the fuzzy matching offered in jump.
	SpaceIgnore
)

// String implements the fmt.Stringer interface.
func (s SpaceSetting) String() string {
	switch s {
	case SpaceSlash:
		return "slash"
	case SpaceIgnore:
		return "ignore"
	default:
		return ""
	}
}

// Settings represents user configurable behaviours for jump.
//
// Keep the default setting values the zero value of their type.
type Settings struct {
	Space    SpaceSetting
	Preserve bool
}

// ReadSettings reads the current user settings.
//
// If the settings don't exist, the zero value Settings is returned.
func (c *fileConfig) ReadSettings() (settings Settings) {
	readFile(c.Settings, &settings)

	return
}

// WriteSettings preserves the user settings.
func (c *fileConfig) WriteSettings(settings Settings) error {
	return writeFile(c.Settings, settings)
}
