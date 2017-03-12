package scoring

import (
	"reflect"
	"testing"
)

func TestEntriesLen(t *testing.T) {
	e1 := Entry{"/foo", &Score{100, Now}}
	e2 := Entry{"/foo/bar", &Score{200, Now}}

	entries := Entries{e1, e2}

	if entries.Len() != 2 {
		t.Errorf("Expected entries length to be 2, got %d", entries.Len())
	}
}

func TestEntriesSwap(t *testing.T) {
	e1 := Entry{"/foo", &Score{100, Now}}
	e2 := Entry{"/foo/bar", &Score{200, Now}}

	entries := Entries{e1, e2}
	expectedEntries := Entries{e2, e1}

	entries.Swap(0, 1)

	if !reflect.DeepEqual(expectedEntries, entries) {
		t.Errorf("Expected entries to be %v, got %v", expectedEntries, entries)
	}
}

func TestEntriesLess(t *testing.T) {
	e1 := Entry{"/foo", &Score{100, Now}}
	e2 := Entry{"/foo/bar", &Score{200, Now}}

	entries := Entries{e1, e2}

	if !entries.Less(0, 1) {
		t.Errorf("Expected %v to be less than %v", e1, e2)
	}
}

func TestEntriesFind(t *testing.T) {
	e1 := Entry{"/foo", &Score{100, Now}}
	e2 := Entry{"/foo/bar", &Score{200, Now}}

	entries := Entries{e1, e2}

	if e, ok := entries.Find("/foo"); e == nil || !ok {
		t.Errorf("Expected %v to have /foo", entries)
	}

	if e, ok := entries.Find("/bar"); e != nil || ok {
		t.Errorf("Expected %v to not have /bar", entries)
	}
}

func TestEntriesRemove(t *testing.T) {
	e1 := Entry{"/foo", &Score{100, Now}}
	e2 := Entry{"/foo/bar", &Score{200, Now}}

	entries := Entries{e1, e2}

	if !entries.Remove("/foo") && entries.Len() != 1 {
		t.Errorf("Expected %v to remove /foo", entries)
	}
}

func TestEntriesSort(t *testing.T) {
	e1 := Entry{"/foo", &Score{100, Now}}
	e2 := Entry{"/foo/bar", &Score{200, Now}}

	entries := Entries{e1, e2}
	expectedEntries := Entries{e1, e2}

	entries.Sort()

	if !reflect.DeepEqual(expectedEntries, entries) {
		t.Errorf("Expected entries to be %v, got %v", expectedEntries, entries)
	}
}
