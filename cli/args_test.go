package cli

import (
	"reflect"
	"testing"
)

func TestParseArgs(t *testing.T) {
	args := ParseArgs([]string{"program", "command"})

	if args[0] != "command" {
		t.Errorf("Expected args[0] to be command, got %v", args[0])
	}
}

func TestFirst(t *testing.T) {
	args := ParseArgs([]string{"program", "command"})

	if args.First() != "command" {
		t.Error("Expected args.First() to be command, got %v", args.First())
	}
}

func TestFirstWithEmptyArgs(t *testing.T) {
	args := ParseArgs([]string{"program"})

	if args.First() != "" {
		t.Error("Expected args.First() to be empty")
	}
}

func TestArgsCommandName(t *testing.T) {
	args := ParseArgs([]string{"program", "command"})

	if command := args.CommandName(); command != "command" {
		t.Errorf("Expected args.CommandName() to be command, got %v", command)
	}
}

func TestArgsHas(t *testing.T) {
	args := ParseArgs([]string{"program", "command", "--foo"})

	if !args.Has("--foo") {
		t.Errorf("Expected %v to have --foo", args)
	}

	if args.Has("--bar") {
		t.Errorf("Expected %v to not have --bar", args)
	}
}

func TestGetWithEquals(t *testing.T) {
	args := ParseArgs([]string{"program", "--option=val"})

	if value := args.Get("--option", "none"); value != "val" {
		t.Errorf("Expected val, got: %s", value)
	}
}

func TestGetWithSpace(t *testing.T) {
	args := ParseArgs([]string{"program", "--option", "val"})

	if value := args.Get("--option", "none"); value != "val" {
		t.Errorf("Expected val, got: %s", value)
	}
}

func TestGetHittingDefaultValue(t *testing.T) {
	args := ParseArgs([]string{"program", "--option"})

	if value := args.Get("--option", "none"); value != "none" {
		t.Errorf("Expected none, got: %s", value)
	}
}

func TestArgsRest(t *testing.T) {
	args := ParseArgs([]string{"program", "command"})

	if got, want := args.Rest(), []string{"command"}; reflect.DeepEqual(got, want) {
		t.Errorf("Expected args.Rest() to be %v, got %v", want, got)
	}
}
