package cmd

import (
	"os"
	p "path"
	"strings"
	"testing"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/scoring"
)

func Test_forgetCmd(t *testing.T) {
	p := p.Join(td, "web-console")

	conf := &testConfig{
		Entries: scoring.Entries{scoring.NewEntry(p)},
	}

	output := capture(&os.Stdout, func() {
		if err := forgetCmd(cli.Args{p}, conf); err != nil {
			t.Errorf("Unexpected error %v", err)
		}

		if err := cleanCmd(cli.Args{}, conf); err != nil {
			t.Errorf("Unexpected error %v", err)
		}
	})

	if !strings.Contains(output, "Cleaning") {
		t.Fatalf("Expected to clean entries, got:\n%s", output)
	}

	entries, _ := conf.ReadEntries()

	if entries.Len() != 0 {
		t.Fatalf("Expected no entries, got %v", entries)
	}
}
