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
	t.Run("--option without a value", func(t *testing.T) {
		args := ParseArgs([]string{"program", "command", "--option"})

		assert.True(t, args.Has("--option"))
		assert.False(t, args.Has("--another"))
	})

	t.Run("--option=with-value", func(t *testing.T) {
		args := ParseArgs([]string{"program", "command", "--option=value"})

		assert.True(t, args.Has("--option"))
		assert.False(t, args.Has("--another"))
	})
}

func TestGet(t *testing.T) {
	args := ParseArgs([]string{"program", "--option=val"})
	assert.Equal(t, "val", args.Get("--option", Optional))

	args = ParseArgs([]string{"program", "--option", "val"})
	assert.Equal(t, "val", args.Get("--option", Optional))

	args = ParseArgs([]string{"program", "--option"})
	assert.Equal(t, "default", args.Get("--option", "default"))
}

func TestRest(t *testing.T) {
	args := ParseArgs([]string{"program", "command"})
	assert.Len(t, 0, args.Rest())

	args = ParseArgs([]string{"program", "command", "arg"})
	assert.Equal(t, []string{"arg"}, args.Rest())
}

func TestWithout(t *testing.T) {
	t.Run("removes existing option", func(t *testing.T) {
		args := ParseArgs([]string{"program", "command", "--option", "value"})
		result := args.Without("--option")

		assert.Equal(t, []string{"command", "value"}, []string(result))
	})

	t.Run("returns same args when option not found", func(t *testing.T) {
		args := ParseArgs([]string{"program", "command"})
		result := args.Without("--nonexistent")

		assert.Equal(t, []string{"command"}, []string(result))
	})
}
