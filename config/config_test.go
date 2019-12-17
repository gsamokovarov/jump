package config

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

var td string

func tempConfig(t *testing.T) Config {
	tempDir, err := ioutil.TempDir(td, ".tmp")
	if err != nil {
		t.Fatalf("Cannot create temporary testing directory: %v", err)
	}

	conf, err := Setup(tempDir)
	if err != nil {
		t.Fatalf("Cannot setup temporary testing directory: %v", err)
	}

	return conf
}

func init() {
	cwd, _ := os.Getwd()
	td = path.Join(cwd, "testdata")
}
