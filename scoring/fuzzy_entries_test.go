package scoring

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

func TestFuzzyEntriesLess(t *testing.T) {
	e1 := &Entry{"/Development/web-console", &Score{200, Now}}
	e2 := &Entry{"/Development/web-client", &Score{100, Now}}

	entries := NewFuzzyEntries(Entries{e1, e2}, "wc")

	assert.True(t, entries.Less(0, 1))
}

func TestFuzzyEntriesSort(t *testing.T) {
	e1 := &Entry{"/Development/web-client", &Score{100, Now}}
	e2 := &Entry{"/Development/web-console", &Score{200, Now}}

	entries := &FuzzyEntries{Entries{e1, e2}, "wc"}
	entries.Sort()

	e, ok := entries.Select(0)
	assert.True(t, ok)

	assert.Equal(t, e2, e)
}

func TestFuzzyEntriesSelect(t *testing.T) {
	e1 := &Entry{"/Development/web-client", &Score{100, Now}}
	e2 := &Entry{"/Development/web-console", &Score{200, Now}}

	entries := &FuzzyEntries{Entries{e1, e2}, "wc"}

	e, ok := entries.Select(0)
	assert.True(t, ok)
	assert.Equal(t, e1, e)

	e, ok = entries.Select(1)
	assert.True(t, ok)
	assert.Equal(t, e2, e)

	e, ok = entries.Select(2)
	assert.False(t, ok)
	assert.Nil(t, e)
}

func TestNewFuzzyEntries(t *testing.T) {
	e1 := &Entry{"/Development/web-client", &Score{100, Now}}
	e2 := &Entry{"/Development/web-console", &Score{200, Now}}

	entries := NewFuzzyEntries(Entries{e1, e2}, "wc")

	e, ok := entries.Select(0)
	assert.True(t, ok)
	assert.Equal(t, e2, e)
}
