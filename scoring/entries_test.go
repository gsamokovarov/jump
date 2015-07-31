package scoring

import (
	"reflect"
	"sort"
	"testing"
)

func TestEntriesSort(t *testing.T) {
	e1 := Entry{"/foo", &Score{100, Now}}
	e2 := Entry{"/foo/bar", &Score{200, Now}}

	entries := Entries([]Entry{e1, e2})
	expectedEntries := Entries([]Entry{e2, e1})

	sort.Sort(entries)

	if !reflect.DeepEqual(expectedEntries, entries) {
		t.Errorf("Expected entries to be %v, got %v", expectedEntries, entries)
	}
}
