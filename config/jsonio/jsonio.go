package jsonio

import (
	"encoding/json"
	"io"
)

// Truncater encapsulates everything that responds to Truncate(int64).
type Truncater interface {
	Truncate(int64) error
}

// WriteSeekerTruncater is the interface needed by the Encode first argument.
// It's a narrower type than a file, and can be more easily mocked in testing.
type WriteSeekerTruncater interface {
	io.WriteSeeker
	Truncater
}

// Decode a JSON value from a readable. Can decode JSON object and arrays into
// slices and plain Go values.
//
// If the decoding cannot happen, an error is returned. Otherwise, the value is
// written straight into v.
func Decode(r io.Reader, v interface{}) error {
	decoder := json.NewDecoder(r)

	for {
		if err := decoder.Decode(v); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}

	return nil
}

// Encode a JSON value to a writer. The encoding is tailored towards jump JSON
// file needs. Every write truncates the writable and seeks it to it's
// beginning, so you can safely decode and encode without manipulating the
// writable.
//
// Technically, the writable needs to implement WriterSeekerTruncater, but you'll most
func Encode(w WriteSeekerTruncater, v interface{}) error {
	if _, err := w.Seek(0, io.SeekStart); err != nil {
		return err
	}

	if err := w.Truncate(0); err != nil {
		return err
	}

	encoder := json.NewEncoder(w)

	return encoder.Encode(v)
}
