package cli

import (
	"fmt"
	"os"
)

// Errf writes to os.Stderr.
//
// Shortcut for fmt.Fprintf(os.Stderr, format).
func Errf(format string, a ...any) (int, error) {
	return fmt.Fprintf(os.Stderr, format, a...)
}

// Outf writes to os.Stdout.
//
// Shortcut for fmt.Fprintf(os.Stdout, format).
func Outf(format string, a ...any) (int, error) {
	return fmt.Fprintf(os.Stdout, format, a...)
}

// Exitf writes to os.Stderr and exits the process with a given code.
//
// Shortcut for Errf(format); os.Exit(code).
func Exitf(code int, format string, a ...any) {
	Errf(format, a...)
	os.Exit(code)
}
