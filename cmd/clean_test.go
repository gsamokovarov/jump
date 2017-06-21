package cmd

import (
	"os"
	"strings"
	"testing"

	"github.com/gsamokovarov/jump/cli"
)

func Test_cleanCmd(t *testing.T) {
	conf := &testConfig{}

	if err := chdirCmd(cli.Args{"/inexistent/dir/dh891n2kisdha"}, conf); err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	entries, _ := conf.ReadEntries()

	if entries.Len() != 1 {
		t.Fatalf("Expected one entry, got %v", entries)
	}

	output := capture(&os.Stdout, func() {
		if err := cleanCmd(cli.Args{}, conf); err != nil {
			t.Errorf("Unexpected error %v", err)
		}
	})

	if !strings.Contains(output, "Cleaning") {
		t.Fatalf("Expected to clean entries, got:\n%s", output)
	}
}
