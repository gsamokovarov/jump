package importer

import (
	"os"
	"path"
	"testing"

	"github.com/gsamokovarov/assert"
	"github.com/gsamokovarov/jump/config"
)

var td string

func TestGuess_Autojump(t *testing.T) {
	imp := Guess("autojump", &config.Testing{})

	_, ok := imp.(*autojump)
	assert.True(t, ok)
}

func TestGuess_Z(t *testing.T) {
	imp := Guess("z", &config.Testing{})

	_, ok := imp.(*z)
	assert.True(t, ok)
}

func TestGuess_Both(t *testing.T) {
	imp := Guess("", &config.Testing{})

	_, ok := imp.(multiImporter)
	assert.True(t, ok)
}

func TestCallback(t *testing.T) {
	var fn Callback

	// Does not crash when fn is nil.
	fn.Call(nil)
}

func init() {
	cwd, _ := os.Getwd()
	td = path.Join(cwd, "testdata")
}
