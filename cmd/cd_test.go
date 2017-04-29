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
			s.Entry{p.Join(td, "web-console"), &s.Score{Weight: 100, Age: s.Now}},
			s.Entry{p.Join(td, "/client/website"), &s.Score{Weight: 90, Age: s.Now}},
		},
	}

	output := capture(&os.Stdout, func() {
		cdCmd(cli.Args{"wc"}, conf)
	})

	expectedPath := p.Join(td, "web-console")

	if !strings.Contains(output, expectedPath) {
		t.Fatalf("Expected path to be %s, got %s", expectedPath, output)
	}
}

func Test_cdCmd_absolutePath(t *testing.T) {
	conf := &testConfig{
		Entries: s.Entries{
			s.Entry{p.Join(td, "web-console"), &s.Score{Weight: 100, Age: s.Now}},
			s.Entry{p.Join(td, "/client/website"), &s.Score{Weight: 90, Age: s.Now}},
		},
	}

	output := capture(&os.Stdout, func() {
		cdCmd(cli.Args{"/absolute/path"}, conf)
	})

	expectedPath := "/absolute/path\n"

	if output != expectedPath {
		t.Fatalf("Expected path to be %s, got %s", expectedPath, output)
	}
}
