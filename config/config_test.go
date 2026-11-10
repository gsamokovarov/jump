package config

import "testing"

func tempConfig(t *testing.T) Config {
	conf, err := Setup(t.TempDir())
	if err != nil {
		t.Fatalf("Cannot setup temporary testing config: %v", err)
	}

	return conf
}
