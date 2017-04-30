package cmd

import (
	"os"
	"strings"
	"testing"

	"github.com/gsamokovarov/jump/cli"
)

func Test_cleanCmd(t *testing.T) {
	conf := &testConfig{}

	chdirCmd(cli.Args{"/inexistent/dir/dh891n2kisdha"}, conf)

	entries, _ := conf.ReadEntries()

	if entries.Len() != 1 {
		t.Fatalf("Expected one entry, got %v", entries)
	}

	output := capture(&os.Stdout, func() {
		cleanCmd(cli.Args{}, conf)
	})

	if !strings.Contains(output, "Cleaning") {
		t.Fatalf("Expected to clean entries, got:\n%s", output)
	}
}
