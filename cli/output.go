package cli

import (
	"fmt"
	"os"
)

// Errf writes to os.Stderr.
//
// Shortcut for fmt.Fprintf(os.Stderr, format).
func Errf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(os.Stderr, format, a...)
}

// Outf writes to os.Stdout.
//
// Shortcut for fmt.Fprintf(os.Stdout, format).
func Outf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(os.Stdout, format, a...)
}
