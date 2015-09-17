package match

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gsamokovarov/jump/lcs"
)

const osSeparator = string(os.PathSeparator)

type Matcher struct {
	term string
}

func (m Matcher) Match(left, right string) bool {
	if caseSensitiveSearch(m.term) {
		m.term = strings.ToLower(m.term)
		left, right = strings.ToLower(left), strings.ToLower(right)
	}

	if containsOsSeparators(m.term) {
		count := strings.Count(m.term, osSeparator)
		left, right = finalChunks(left, count), finalChunks(right, count)
	} else {
		left, right = filepath.Base(left), filepath.Base(right)
	}

	return lcs.Length(left, m.term) >= lcs.Length(right, m.term)
}

func NewMatcher(term string) *Matcher {
	return &Matcher{term}
}

func caseSensitiveSearch(str string) bool {
	return strings.ToLower(str) == str || str == strings.ToUpper(str)
}

func containsOsSeparators(str string) bool {
	return strings.ContainsAny(str, osSeparator)
}

func finalChunks(path string, sepCount int) string {
	chunk := ""

	for i := 0; i < sepCount; i++ {
		dir, file := filepath.Split(path)
		chunk = filepath.Base(dir) + file + chunk
	}

	return chunk
}
