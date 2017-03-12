package scoring

import (
	"testing"
)

func TestFuzzyEntriesLess(t *testing.T) {
	e1 := Entry{"/Development/web-console", &Score{200, Now}}
	e2 := Entry{"/Development/web-client", &Score{100, Now}}

	entries := NewFuzzyEntries(&Entries{e1, e2}, "wc")

	if !entries.Less(0, 1) {
		t.Errorf("Expected %v to come before %v", e1, e2)
	}
}

func TestFuzzyEntriesSort(t *testing.T) {
	e1 := Entry{"/Development/web-client", &Score{100, Now}}
	e2 := Entry{"/Development/web-console", &Score{200, Now}}

	entries := &FuzzyEntries{Entries{e1, e2}, "wc"}
	entries.Sort()

	if e, ok := entries.Select(0); e.Path != "/Development/web-console" || !ok {
		t.Errorf("Expected %v to be %v", e, e2)
	}
}

func TestFuzzyEntriesSelect(t *testing.T) {
	e1 := Entry{"/Development/web-client", &Score{100, Now}}
	e2 := Entry{"/Development/web-console", &Score{200, Now}}

	entries := &FuzzyEntries{Entries{e1, e2}, "wc"}

	if e, ok := entries.Select(0); e.Path != "/Development/web-client" || !ok {
		t.Errorf("Expected %v to be %v", e, e1)
	}

	if e, ok := entries.Select(1); e.Path != "/Development/web-console" || !ok {
		t.Errorf("Expected %v to be %v", e, e2)
	}

	if e, ok := entries.Select(2); e != nil || ok {
		t.Errorf("Expected %v to be nil", e)
	}
}

func TestNewFuzzyEntries(t *testing.T) {
	e1 := Entry{"/Development/web-client", &Score{100, Now}}
	e2 := Entry{"/Development/web-console", &Score{200, Now}}

	entries := NewFuzzyEntries(&Entries{e1, e2}, "wc")

	if e, ok := entries.Select(0); e.Path != "/Development/web-console" || !ok {
		t.Errorf("Expected %v to be %v", e, e2)
	}
}
