package cmd

import (
	"os"
	p "path"
	"strings"
	"testing"

	"github.com/gsamokovarov/jump/cli"
	s "github.com/gsamokovarov/jump/scoring"
)

func Test_cdCmd(t *testing.T) {
	conf := &testConfig{
		Entries: s.Entries{
			&s.Entry{p.Join(td, "web-console"), &s.Score{Weight: 100, Age: s.Now}},
			&s.Entry{p.Join(td, "/client/website"), &s.Score{Weight: 90, Age: s.Now}},
		},
	}

	output := capture(&os.Stdout, func() {
		if err := cdCmd(cli.Args{"wc"}, conf); err != nil {
			t.Errorf("Unexpected error %v", err)
		}
	})

	expectedPath := p.Join(td, "web-console")

	if !strings.Contains(output, expectedPath) {
		t.Fatalf("Expected path to be %s, got %s", expectedPath, output)
	}
}

func Test_cdCmd_absolutePath(t *testing.T) {
	conf := &testConfig{
		Entries: s.Entries{
			&s.Entry{p.Join(td, "web-console"), &s.Score{Weight: 100, Age: s.Now}},
			&s.Entry{p.Join(td, "/client/website"), &s.Score{Weight: 90, Age: s.Now}},
		},
	}

	output := capture(&os.Stdout, func() {
		if err := cdCmd(cli.Args{"/absolute/path"}, conf); err != nil {
			t.Errorf("Unexpected error %v", err)
		}
	})

	expectedPath := "/absolute/path\n"

	if output != expectedPath {
		t.Fatalf("Expected path to be %s, got %s", expectedPath, output)
	}
}

func Test_cdCmd_exactMatch(t *testing.T) {
	p1 := p.Join(td, "web-console")
	p2 := p.Join(td, "/client/website")
	p3 := p.Join(td, "web")

	conf := &testConfig{
		Entries: s.Entries{
			&s.Entry{p1, &s.Score{Weight: 100, Age: s.Now}},
			&s.Entry{p2, &s.Score{Weight: 90, Age: s.Now}},
			&s.Entry{p3, &s.Score{Weight: 80, Age: s.Now}},
		},
	}

	output := capture(&os.Stdout, func() {
		if err := cdCmd(cli.Args{"web"}, conf); err != nil {
			t.Errorf("Unexpected error %v", err)
		}
	})

	// If someone typed a dir exactly, jump straight to it. Not good for short
	// names like this test here, but pretty useful for most of the catch-all
	// directories.
	expectedPath := p3 + "\n"

	if output != expectedPath {
		t.Fatalf("Expected path to be %s, got %s", expectedPath, output)
	}
}
