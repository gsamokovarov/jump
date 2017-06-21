package cmd

import (
	"os"
	p "path"
	"testing"

	"github.com/gsamokovarov/jump/cli"
)

func Test_chdirCmd(t *testing.T) {
	p1 := p.Join(td, "web-console")
	p2 := p.Join(td, "/client/website")

	conf := &testConfig{}

	entries, _ := conf.ReadEntries()

	if entries.Len() != 0 {
		t.Fatal("Expected entries to be empty")
	}

	// Test that a new entry is added to the list.
	chdirCmd(cli.Args{p1}, conf)

	entries, _ = conf.ReadEntries()

	if entries.Len() != 1 {
		t.Fatalf("Expected one entry, got %v", entries)
	}

	// Test that a new entry is added to the list.
	if err := chdirCmd(cli.Args{p2}, conf); err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	entries, _ = conf.ReadEntries()

	if entries.Len() != 2 {
		t.Fatalf("Expected two entries, got %v", entries)
	}

	// Test that once an existing path is entered again, it's not duplicated in
	// the entries.
	chdirCmd(cli.Args{p2}, conf)

	entries, _ = conf.ReadEntries()

	if entries.Len() != 2 {
		t.Fatalf("Expected two entries, got %v", entries)
	}
}

func Test_chdirCmd_cwd(t *testing.T) {
	conf := &testConfig{}

	entries, _ := conf.ReadEntries()

	if entries.Len() != 0 {
		t.Fatal("Expected entries to be empty")
	}

	// Test that the current directory is added to the list.
	if err := chdirCmd(cli.Args{}, conf); err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	entries, _ = conf.ReadEntries()

	if entries.Len() != 1 {
		t.Fatalf("Expected one entry, got %v", entries)
	}

	cwd, _ := os.Getwd()

	if (*entries)[0].Path != cwd {
		t.Fatalf("Expected entry to be %s, got %v", cwd, entries)
	}
}
