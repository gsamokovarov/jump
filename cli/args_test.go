package cli

import (
	"testing"

	"github.com/gsamokovarov/assert"
)

func TestParseArgs(t *testing.T) {
	args := ParseArgs([]string{"program", "command"})
	assert.Len(t, 1, args)

	assert.Equal(t, "command", args[0])
}

func TestFirst(t *testing.T) {
	args := ParseArgs([]string{"program", "command"})
	assert.Len(t, 1, args)
	assert.Equal(t, "command", args.First())

	args = ParseArgs([]string{"program"})
	assert.Equal(t, "", args.First())
}

func TestCommandName(t *testing.T) {
	args := ParseArgs([]string{"program", "command"})

	assert.Equal(t, "command", args.CommandName())
}

func TestHas(t *testing.T) {
	args := ParseArgs([]string{"program", "command", "--foo"})

	assert.True(t, args.Has("--foo"))
	assert.False(t, args.Has("--bar"))
}

func TestGet(t *testing.T) {
	args := ParseArgs([]string{"program", "--option=val"})
	assert.Equal(t, "val", args.Get("--option", "none"))

	args = ParseArgs([]string{"program", "--option", "val"})
	assert.Equal(t, "val", args.Get("--option", "none"))

	args = ParseArgs([]string{"program", "--option"})
	assert.Equal(t, "none", args.Get("--option", "none"))
}

func TestRest(t *testing.T) {
	args := ParseArgs([]string{"program", "command"})
	assert.Len(t, 0, args.Rest())

	args = ParseArgs([]string{"program", "command", "arg"})
	assert.Equal(t, []string{"arg"}, args.Rest())
}
