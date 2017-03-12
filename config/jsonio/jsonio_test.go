package jsonio

import (
	"bytes"
	"io"
	"reflect"
	"testing"
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
	if err := Decode(r, &value); err != nil {
		t.Errorf("Decoding unsuccessful: %v", err)
	}

	if !value.Ok {
		t.Errorf("Expected value.Ok to be true, got %v", value)
	}
}

func TestBadDecode(t *testing.T) {
	var value struct{ Ok bool }

	r := newTestIO(`{"Ok":true`)
	if err := Decode(r, &value); err == nil {
		t.Errorf("Expected error, got %v", err)
	}
}

func TestEncode(t *testing.T) {
	var value struct{ Ok bool }
	value.Ok = true

	w := newTestIO("")
	if err := Encode(w, value); err != nil {
		t.Errorf("Encoding unsuccessful: %v", err)
	}

	want := `{"Ok":true}`
	if got, err := w.ReadString('\n'); reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, got %v %v", want, got, err)
	}
}
