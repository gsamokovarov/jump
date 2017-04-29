package cmd

import (
	"os"
	p "path"
	"strings"
	"testing"

	"github.com/gsamokovarov/jump/cli"
	s "github.com/gsamokovarov/jump/scoring"
)

func Test_hintCmd(t *testing.T) {
	wc := p.Join(td, "web-console")
	web := p.Join(td, "/client/website")

	conf := &testConfig{
		Entries: s.Entries{
			s.Entry{wc, &s.Score{Weight: 100, Age: s.Now}},
			s.Entry{web, &s.Score{Weight: 90, Age: s.Now}},
		},
	}

	output := capture(&os.Stdout, func() {
		hintCmd(cli.Args{}, conf)
	})

	lines := strings.Split(output, "\n")

	if lines[0] != wc {
		t.Fatalf("Expected first line to be %s, got %s", wc, lines[0])
	}

	if lines[1] != web {
		t.Fatalf("Expected first line to be %s, got %s", web, lines[1])
	}
}
