package importer

import (
	"bytes"
	"encoding/binary"
	"os"
	p "path"
	"testing"
	"time"

	"github.com/gsamokovarov/assert"
	"github.com/gsamokovarov/jump/config"
)

func TestZoxide(t *testing.T) {
	conf := &config.InMemory{}
	configPath := p.Join(td, "zoxide", "db.zo")

	imp := Zoxide(conf, configPath)

	err := imp.Import(nil)
	assert.Nil(t, err)

	// The fixture holds 5 records, but one is a duplicate path, so only 4
	// unique entries are imported.
	assert.
		Len(t, 4, conf.Entries).
		// 0
		Equal(t, "/Users/genadi/dev/masse", conf.Entries[0].Path).
		Equal(t, 1, conf.Entries[0].Score.Weight).
		Equal(t, time.Unix(1536272502, 0), conf.Entries[0].Score.Age).
		// 1
		Equal(t, "/Users/genadi/dev/gloat", conf.Entries[1].Path).
		Equal(t, 2, conf.Entries[1].Score.Weight).
		Equal(t, time.Unix(1536272492, 0), conf.Entries[1].Score.Age).
		// 2
		Equal(t, "/Users/genadi/dev", conf.Entries[2].Path).
		Equal(t, 4, conf.Entries[2].Score.Weight).
		Equal(t, time.Unix(1536272506, 0), conf.Entries[2].Score.Age).
		// 3
		Equal(t, "/Users/genadi/dev/jump", conf.Entries[3].Path).
		Equal(t, 10, conf.Entries[3].Score.Weight).
		Equal(t, time.Unix(1536272816, 0), conf.Entries[3].Score.Age)

	for i, j := 0, 1; i < len(conf.Entries)-1; i, j = i+1, j+1 {
		assert.True(t, conf.Entries[i].CalculateScore() <= conf.Entries[j].CalculateScore())
	}
}

func TestZoxideCustomEnv(t *testing.T) {
	conf := &config.InMemory{}

	os.Setenv("_ZO_DATA_DIR", p.Join(td, "zoxide"))
	defer func() {
		os.Unsetenv("_ZO_DATA_DIR")
	}()

	imp := Zoxide(conf)

	err := imp.Import(nil)
	assert.Nil(t, err)

	assert.Len(t, 4, conf.Entries)
}

func TestZoxideUnsupportedVersion(t *testing.T) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, uint32(2))
	binary.Write(&buf, binary.LittleEndian, uint64(0))

	_, err := parseZoxideDB(buf.Bytes())
	assert.NotNil(t, err)
}

func TestZoxideTruncated(t *testing.T) {
	// A valid header promising one entry, but no entry data follows.
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, zoxideVersion)
	binary.Write(&buf, binary.LittleEndian, uint64(1))

	_, err := parseZoxideDB(buf.Bytes())
	assert.NotNil(t, err)
}

func TestZoxideEmpty(t *testing.T) {
	// A valid header with no directories yields no entries and no error.
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, zoxideVersion)
	binary.Write(&buf, binary.LittleEndian, uint64(0))

	entries, err := parseZoxideDB(buf.Bytes())
	assert.Nil(t, err)
	assert.Len(t, 0, entries)
}

func TestZoxideTruncatedHeader(t *testing.T) {
	// The version is present, but the entry count is cut short.
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, zoxideVersion)
	buf.Write([]byte{0x01, 0x00}) // 2 of the 8 count bytes

	_, err := parseZoxideDB(buf.Bytes())
	assert.NotNil(t, err)
}

func TestZoxideTruncatedPath(t *testing.T) {
	// An entry declares a 23-byte path but supplies fewer bytes.
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, zoxideVersion)
	binary.Write(&buf, binary.LittleEndian, uint64(1))
	binary.Write(&buf, binary.LittleEndian, uint64(23))
	buf.WriteString("/short")

	_, err := parseZoxideDB(buf.Bytes())
	assert.NotNil(t, err)
}

func TestZoxideTooLarge(t *testing.T) {
	// A structurally valid but oversized database must be rejected up front,
	// mirroring zoxide's own 32 MiB deserialization limit. The header is valid
	// (version + count=0) and the trailing bytes would be ignored, so without
	// the size guard this would parse successfully — this isolates the guard.
	data := make([]byte, zoxideMaxDBSize+1)
	binary.LittleEndian.PutUint32(data[0:], zoxideVersion)
	binary.LittleEndian.PutUint64(data[4:], 0) // count = 0

	_, err := parseZoxideDB(data)
	assert.NotNil(t, err)
}

func TestZoxideAtSizeLimit(t *testing.T) {
	// A database exactly at the limit is allowed (the guard rejects only what
	// is strictly larger). Valid header + count=0 + padding parses to nothing.
	data := make([]byte, zoxideMaxDBSize)
	binary.LittleEndian.PutUint32(data[0:], zoxideVersion)
	binary.LittleEndian.PutUint64(data[4:], 0) // count = 0

	entries, err := parseZoxideDB(data)
	assert.Nil(t, err)
	assert.Len(t, 0, entries)
}

func TestZoxideCorruptPathLength(t *testing.T) {
	// A wildly out-of-range path length must be rejected by the guard rather
	// than panicking on a slice expression or exhausting memory.
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, zoxideVersion)
	binary.Write(&buf, binary.LittleEndian, uint64(1))
	binary.Write(&buf, binary.LittleEndian, uint64(1)<<63)

	_, err := parseZoxideDB(buf.Bytes())
	assert.NotNil(t, err)
}
