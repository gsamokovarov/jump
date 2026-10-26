package fuzzy

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

func TestScore(t *testing.T) {
	t.Run("no separators matches the base name only", func(t *testing.T) {
		assert.Equal(t, 3, Score("soc", "/Users/foo/Development/society"))
		assert.True(t, Score("soc", "/Users/foo/society/website") < 3)
	})

	t.Run("components are matched right-aligned", func(t *testing.T) {
		assert.Equal(t, 6, Score("dev/soc", "/Users/foo/Development/society"))
		// "dev" is matched against "society" and "soc" against "website" here.
		assert.True(t, Score("dev/soc", "/Users/foo/society/website") < 6)
	})

	t.Run("absolute terms line up with the same prefix", func(t *testing.T) {
		downloads := Score("/home/foo/Dow", "/home/foo/Downloads")
		wallpaper := Score("/home/foo/Dow", "/home/foo/Development/ponor-wallpaper")
		assert.True(t, downloads > wallpaper)
	})

	t.Run("trailing separators are ignored", func(t *testing.T) {
		assert.Equal(t, Score("website", "/Users/foo/website"), Score("website/", "/Users/foo/website"))
	})

	t.Run("repeated separators are ignored", func(t *testing.T) {
		// A j dev/ web turns into dev//web and dev still matches the parent.
		assert.Equal(t, Score("dev/web", "/Users/foo/Development/website"), Score("dev//web", "/Users/foo/Development/website"))
		assert.True(t, Score("dev//web", "/Users/foo/Development/website") > Score("dev//web", "/Users/foo/Development/foo/website"))
	})

	t.Run("empty term and root path score nothing", func(t *testing.T) {
		assert.Equal(t, 0, Score("", "/Users/foo/website"))
		assert.Equal(t, 0, Score("website", "/"))
		assert.Equal(t, 0, Score("", ""))
	})

	t.Run("case sensitivity", func(t *testing.T) {
		assert.Equal(t, 3, Score("Dev", "/Users/foo/Development"))
		assert.True(t, Score("Dev", "/Users/foo/development") < 3)
		assert.Equal(t, 3, Score("dev", "/Users/foo/Development"))
	})

	t.Run("exact base name matches score above fuzzy ones", func(t *testing.T) {
		assert.True(t, Score("webcon", "/Users/foo/webcon") > Score("webcon", "/Users/foo/web-console"))
		assert.True(t, Score("cli/webcon", "/Users/foo/client/webcon") > Score("cli/webcon", "/Users/foo/client/web-console"))
	})

	t.Run("short exact matches get no preference", func(t *testing.T) {
		assert.Equal(t, Score("web", "/Users/foo/web"), Score("web", "/Users/foo/web-console"))
	})
}
