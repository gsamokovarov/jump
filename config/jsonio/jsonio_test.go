package jsonio

import (
	"bytes"
	"io"
	"testing"

	"github.com/gsamokovarov/assert"
)

type testIO struct {
	*bytes.Buffer
}

func (w *testIO) Seek(offset int64, whence int) (ret int64, err error) {
	if offset != 0 {
		panic("expected offset to be 0")
	}

	if whence != io.SeekStart {
		panic("expected whence to be io.SeekStart")
	}

	return
}

func (w *testIO) Truncate(size int64) error {
	if size != 0 {
		panic("expected size to be 0")
	}

	return nil
}

func newTestIO(s string) *testIO {
	return &testIO{Buffer: bytes.NewBufferString(s)}
}

func TestDecode(t *testing.T) {
	var value struct{ Ok bool }

	r := newTestIO(`{"Ok":true}`)
	assert.Nil(t, Decode(r, &value))

	assert.True(t, value.Ok)
}

func TestBadDecode(t *testing.T) {
	var value struct{ Ok bool }

	r := newTestIO(`{"Ok":true`)
	assert.NotNil(t, Decode(r, &value))
}

func TestEncode(t *testing.T) {
	var value struct{ Ok bool }
	value.Ok = true

	w := newTestIO("")
	assert.Nil(t, Encode(w, value))

	json, err := w.ReadString('\n')
	assert.Nil(t, err)

	assert.Equal(t, "{\"Ok\":true}\n", json)
}
