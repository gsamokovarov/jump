package cli

import (
	"reflect"
	"testing"
)

func TestArgsCommandName(t *testing.T) {
	args := ParseArgs([]string{"program", "command-name"})

	if command := args.CommandName(); command != "command-name" {
		t.Errorf("Expected args.CommandName() to be command-name, got %v", command)
	}
}

func TestArgsRest(t *testing.T) {
	args := ParseArgs([]string{"program", "command-name"})

	if got, want := args.Rest(), []string{"command-name"}; reflect.DeepEqual(got, want) {
		t.Errorf("Expected args.Rest() to be %v, got %v", want, got)
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
