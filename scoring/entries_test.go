package scoring

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

func TestEntriesLen(t *testing.T) {
	e1 := &Entry{"/foo", &Score{100, Now}}
	e2 := &Entry{"/foo/bar", &Score{200, Now}}

	entries := Entries{e1, e2}
	assert.Len(t, 2, entries)
}

func TestEntriesFind(t *testing.T) {
	e1 := &Entry{"/foo", &Score{100, Now}}
	e2 := &Entry{"/foo/bar", &Score{200, Now}}

	entries := Entries{e1, e2}

	e, ok := entries.Find("/foo")
	assert.True(t, ok)
	assert.Equal(t, e1, e)

	e, ok = entries.Find("/foo/bar")
	assert.True(t, ok)
	assert.Equal(t, e2, e)
}

func TestEntriesRemove(t *testing.T) {
	e1 := &Entry{"/foo", &Score{100, Now}}
	e2 := &Entry{"/foo/bar", &Score{200, Now}}

	entries := Entries{e1, e2}

	assert.True(t, entries.Remove("/foo"))
	assert.Len(t, 1, entries)
}

func TestEntriesSort(t *testing.T) {
	e1 := &Entry{"/foo", &Score{100, Now}}
	e2 := &Entry{"/foo/bar", &Score{200, Now}}

	entries := Entries{e2, e1}
	entries.Sort()

	assert.Equal(t, Entries{e1, e2}, entries)
}

func TestEntriesUnder(t *testing.T) {
	e1 := &Entry{"/home", &Score{100, Now}}
	e2 := &Entry{"/home/user", &Score{200, Now}}
	e3 := &Entry{"/home/user/documents", &Score{300, Now}}
	e4 := &Entry{"/home/user/downloads", &Score{400, Now}}
	e5 := &Entry{"/var/log", &Score{500, Now}}

	entries := Entries{e1, e2, e3, e4, e5}

	t.Run("returns entries under specified path", func(t *testing.T) {
		result := entries.Under("/home")
		expected := Entries{e2, e3, e4}
		assert.Equal(t, expected, result)
	})

	t.Run("returns entries under nested path", func(t *testing.T) {
		result := entries.Under("/home/user")
		expected := Entries{e3, e4}
		assert.Equal(t, expected, result)
	})

	t.Run("returns empty when no entries under path", func(t *testing.T) {
		result := entries.Under("/nonexistent")
		assert.Len(t, 0, result)
	})

	t.Run("excludes exact path match", func(t *testing.T) {
		result := entries.Under("/home/user")
		for _, entry := range result {
			assert.NotEqual(t, "/home/user", entry.Path)
		}
	})
}

func TestEntriesRanked(t *testing.T) {
	t.Run("best match first", func(t *testing.T) {
		e1 := &Entry{"/Development/web-console", &Score{100, Now}}
		e2 := &Entry{"/Development/website", &Score{200, Now}}

		ranked := Entries{e1, e2}.Ranked("wc")

		assert.Equal(t, Entries{e1, e2}, ranked)
	})

	t.Run("heavier entry first among equal matches", func(t *testing.T) {
		e1 := &Entry{"/Development/web-client", &Score{100, Now}}
		e2 := &Entry{"/Development/web-console", &Score{200, Now}}

		assert.Equal(t, Entries{e2, e1}, Entries{e1, e2}.Ranked("wc"))
		assert.Equal(t, Entries{e2, e1}, Entries{e2, e1}.Ranked("wc"))
	})

	t.Run("empty term orders by score alone", func(t *testing.T) {
		e1 := &Entry{"/Development/web-client", &Score{100, Now}}
		e2 := &Entry{"/Development/web-console", &Score{200, Now}}

		assert.Equal(t, Entries{e2, e1}, Entries{e1, e2}.Ranked(""))
	})

	t.Run("leaves the entries untouched", func(t *testing.T) {
		e1 := &Entry{"/Development/web-client", &Score{100, Now}}
		e2 := &Entry{"/Development/web-console", &Score{200, Now}}

		entries := Entries{e1, e2}
		entries.Ranked("wc")

		assert.Equal(t, Entries{e1, e2}, entries)
	})
}
