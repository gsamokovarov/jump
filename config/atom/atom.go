package atom

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Truncater encapsulates everything that responds to Truncate(int64).
type Truncater interface {
	Truncate(int64) error
}

// File is a minimal file interface for atomic writes.
type File interface {
	io.Reader
	io.Writer
	io.Seeker
	io.Closer

	Truncater
}

// Open opens an atomic file for writing.
func Open(name string) (File, error) {
	tmp, err := ioutil.TempFile(filepath.Dir(name), ".jump")
	if err != nil {
		return nil, err
	}

	// Copy the contents of the input file, so we can read it from the temporary.
	if _, err := os.Stat(name); !os.IsNotExist(err) {
		file, err := os.OpenFile(name, os.O_RDWR, 0644)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		_, err = io.Copy(tmp, file)
		if err != nil {
			return nil, err
		}

		_, err = tmp.Seek(0, 0)
		if err != nil {
			return nil, err
		}
	}

	return &file{to: name, tmp: tmp}, nil
}

type file struct {
	to  string
	tmp *os.File
	// dirty indicates whether this tmpfile has experienced errors which mean it
	// should not be used to update the original.
	dirty bool
}

func (f *file) Read(p []byte) (int, error) {
	return f.tmp.Read(p)
}

func (f *file) Write(p []byte) (int, error) {
	n, err := f.tmp.Write(p)
	if err != nil {
		f.dirty = true
	}

	return n, err
}

func (f *file) Seek(offset int64, whence int) (int64, error) {
	return f.tmp.Seek(offset, whence)
}

func (f *file) Truncate(n int64) error {
	return f.tmp.Truncate(n)
}

func (f *file) Close() error {
	if err := f.tmp.Close(); err != nil {
		return err
	}

	if f.dirty {
		return nil
	}

	return os.Rename(f.tmp.Name(), f.to)
}
