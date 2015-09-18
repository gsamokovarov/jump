package fuzzy

import "testing"

func TestNoSeparators(t *testing.T) {
	norm := NewNormalizer("soc")

	if norm.NormalizePath("/Users/foo/Development/society") != "society" {
		t.Error("Expected soc to normalize to society")
	}
}

func TestSingleSeparator(t *testing.T) {
	norm := NewNormalizer("dev/soc")

	if norm.NormalizePath("/Users/foo/Development/society") != "developmentsociety" {
		t.Error("Expected dev/soc to normalize to developmentsociety")
	}
}

func TestMultipleSeparators(t *testing.T) {
	norm := NewNormalizer("dev/soc/website")

	if norm.NormalizePath("/Users/foo/Development/society/website") != "developmentsocietywebsite" {
		t.Error("Expected dev/soc/web to normalize to developmentsocietywebsite")
	}
}

func TestCaseSensitivity(t *testing.T) {
	norm := NewNormalizer("Dev")

	if norm.NormalizePath("/Users/foo/Development") != "Development" {
		t.Error("Expected Dev to normalize to Development")
	}
}
