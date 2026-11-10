package config

import (
	"os"
	"strconv"
)

// Search represents a search term used for advancing through the entries of
// the same term.
type Search struct {
	Term  string
	Index int
}

// ReadSearch reads the search to retry.
//
// A process cannot change the environment of its shell, so the shell
// integration remembers the last search itself. It passes the term in
// $JUMP_TERM and the number of the retry in $JUMP_TRY, only when j is called
// without arguments. If they are not set, a zero value Search is returned.
func (c *fileConfig) ReadSearch() Search {
	index, err := strconv.Atoi(os.Getenv("JUMP_TRY"))
	if err != nil || index < 0 {
		index = 0
	}

	return Search{os.Getenv("JUMP_TERM"), index}
}
