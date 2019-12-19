package config

import (
	"os"
	"path"
	"testing"
)

var td string

func tempConfig(t *testing.T) Config {
	conf, err := Temporary(td, ".tmp")
	if err != nil {
		t.Fatalf("Cannot setup temporary testing config: %v", err)
	}

	return conf
}

func init() {
	cwd, _ := os.Getwd()
	td = path.Join(cwd, "testdata")
}
