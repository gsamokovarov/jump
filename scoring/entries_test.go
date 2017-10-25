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

func TestEntriesSwap(t *testing.T) {
	e1 := &Entry{"/foo", &Score{100, Now}}
	e2 := &Entry{"/foo/bar", &Score{200, Now}}

	entries := Entries{e1, e2}
	entries.Swap(0, 1)

	assert.Equal(t, Entries{e2, e1}, entries)
}

func TestEntriesLess(t *testing.T) {
	e1 := &Entry{"/foo", &Score{100, Now}}
	e2 := &Entry{"/foo/bar", &Score{200, Now}}

	entries := Entries{e1, e2}

	assert.True(t, entries.Less(0, 1))
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
