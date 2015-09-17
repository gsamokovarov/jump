package fuzzy

import "testing"

func TestTwoNaiveStrings(t *testing.T) {
	if Length("fd", "falcon-dev") != 2 {
		t.Error("Expected fd and falcon-dev to have an LCS length of 2")
	}
}

func TestLongerAfterShorter(t *testing.T) {
	if Length("falcon-dev", "fd") != Length("fd", "falcon-dev") {
		t.Error("Expected auto sorting of left and right")
	}
}
