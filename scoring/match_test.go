package scoring

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

func TestMatch(t *testing.T) {
	t.Run("no separators matches the base name only", func(t *testing.T) {
		assert.Equal(t, 3, match("soc", "/Users/foo/Development/society"))
		assert.True(t, match("soc", "/Users/foo/society/website") < 3)
	})

	t.Run("components are matched right-aligned", func(t *testing.T) {
		assert.Equal(t, 6, match("dev/soc", "/Users/foo/Development/society"))
		// "dev" is matched against "society" and "soc" against "website" here.
		assert.True(t, match("dev/soc", "/Users/foo/society/website") < 6)
	})

	t.Run("absolute terms line up with the same prefix", func(t *testing.T) {
		downloads := match("/home/foo/Dow", "/home/foo/Downloads")
		wallpaper := match("/home/foo/Dow", "/home/foo/Development/ponor-wallpaper")
		assert.True(t, downloads > wallpaper)
	})

	t.Run("trailing separators are ignored", func(t *testing.T) {
		assert.Equal(t, match("website", "/Users/foo/website"), match("website/", "/Users/foo/website"))
	})

	t.Run("repeated separators are ignored", func(t *testing.T) {
		// A j dev/ web turns into dev//web and dev still matches the parent.
		assert.Equal(t, match("dev/web", "/Users/foo/Development/website"), match("dev//web", "/Users/foo/Development/website"))
		assert.True(t, match("dev//web", "/Users/foo/Development/website") > match("dev//web", "/Users/foo/Development/foo/website"))
	})

	t.Run("empty term and root path match nothing", func(t *testing.T) {
		assert.Equal(t, 0, match("", "/Users/foo/website"))
		assert.Equal(t, 0, match("website", "/"))
		assert.Equal(t, 0, match("", ""))
	})

	t.Run("case sensitivity", func(t *testing.T) {
		assert.Equal(t, 3, match("Dev", "/Users/foo/Development"))
		assert.True(t, match("Dev", "/Users/foo/development") < 3)
		assert.Equal(t, 3, match("dev", "/Users/foo/Development"))
	})

	t.Run("exact base name matches rank above fuzzy ones", func(t *testing.T) {
		assert.True(t, match("webcon", "/Users/foo/webcon") > match("webcon", "/Users/foo/web-console"))
		assert.True(t, match("cli/webcon", "/Users/foo/client/webcon") > match("cli/webcon", "/Users/foo/client/web-console"))
	})

	t.Run("short exact matches get no preference", func(t *testing.T) {
		assert.Equal(t, match("web", "/Users/foo/web"), match("web", "/Users/foo/web-console"))
	})
}

func TestLCS(t *testing.T) {
	t.Run("two naive strings", func(t *testing.T) {
		assert.Equal(t, 2, lcs("fd", "falcon-dev"))
	})

	t.Run("longer after shorter", func(t *testing.T) {
		assert.Equal(t, lcs("falcon-dev", "fd"), lcs("fd", "falcon-dev"))
	})
}
