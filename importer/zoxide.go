package importer

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"time"

	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

// zoxideVersion is the db.zo format version jump knows how to read. zoxide
// itself refuses to open databases tagged with a different version, and so do
// we.
const zoxideVersion uint32 = 3

// zoxideMaxDBSize mirrors the 32 MiB limit zoxide enforces when deserializing
// db.zo. Anything larger is treated as corrupt rather than parsed.
const zoxideMaxDBSize = 32 << 20

var zoxideDefaultConfigPaths = []string{
	"$_ZO_DATA_DIR/db.zo",
	"$XDG_DATA_HOME/zoxide/db.zo",
	"$HOME/.local/share/zoxide/db.zo",
}

// Zoxide is an importer for the zoxide tool.
func Zoxide(conf config.Config, configPaths ...string) Importer {
	if len(configPaths) == 0 {
		configPaths = zoxideDefaultConfigPaths
	}

	return &zoxide{
		config:      conf,
		configPaths: configPaths,
	}
}

type zoxide struct {
	config      config.Config
	configPaths []string
}

func (i *zoxide) Import(fn Callback) error {
	zoxideEntries, err := i.parseConfig()
	if err != nil {
		return err
	}

	jumpEntries, err := i.config.ReadEntries()
	if err != nil {
		return err
	}

	for _, entry := range zoxideEntries {
		if _, found := jumpEntries.Find(entry.Path); found {
			continue
		}

		fn.Call(entry)

		jumpEntries = append(jumpEntries, entry)
	}

	return i.config.WriteEntries(jumpEntries)
}

func (i *zoxide) parseConfig() (scoring.Entries, error) {
	data, err := readConfigBytes(i.configPaths)
	if err != nil {
		return nil, err
	}

	return parseZoxideDB(data)
}

// parseZoxideDB decodes the binary db.zo database.
//
// zoxide serializes its database with bincode using fixed-width, little-endian
// integers. The layout is a u32 version tag followed by a length-prefixed list
// of directories:
//
//	u32 version (== zoxideVersion)
//	u64 len(dirs)
//	dirs[len]:
//	  u64 len(path)
//	  u8  path[len(path)]  (UTF-8)
//	  f64 rank
//	  u64 last_accessed    (Unix seconds)
//
// Duplicate paths are left in place; the caller de-duplicates while merging.
func parseZoxideDB(data []byte) (scoring.Entries, error) {
	if len(data) > zoxideMaxDBSize {
		return nil, fmt.Errorf("importer: zoxide database is %d bytes, larger than the %d byte limit", len(data), zoxideMaxDBSize)
	}

	r := zoxideDBReader{data: data}

	version, err := r.uint32()
	if err != nil {
		return nil, fmt.Errorf("importer: cannot read zoxide database version: %w", err)
	}
	if version != zoxideVersion {
		return nil, fmt.Errorf("importer: unsupported zoxide database version %d, jump supports %d", version, zoxideVersion)
	}

	count, err := r.uint64()
	if err != nil {
		return nil, fmt.Errorf("importer: cannot read zoxide entry count: %w", err)
	}

	var entries scoring.Entries

	for n := uint64(0); n < count; n++ {
		entry, err := readZoxideEntry(&r)
		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

func readZoxideEntry(r *zoxideDBReader) (*scoring.Entry, error) {
	path, err := r.str()
	if err != nil {
		return nil, fmt.Errorf("importer: cannot read zoxide entry path: %w", err)
	}

	rank, err := r.float64()
	if err != nil {
		return nil, fmt.Errorf("importer: cannot read zoxide entry rank: %w", err)
	}

	lastAccessed, err := r.uint64()
	if err != nil {
		return nil, fmt.Errorf("importer: cannot read zoxide entry timestamp: %w", err)
	}

	return &scoring.Entry{
		Path: path,
		Score: &scoring.Score{
			Weight: int64(math.Round(rank)),
			Age:    time.Unix(int64(lastAccessed), 0),
		},
	}, nil
}

// zoxideDBReader consumes fixed-width little-endian values from a db.zo byte
// slice without copying it. Reads past the end return io.ErrUnexpectedEOF.
type zoxideDBReader struct {
	data []byte
	pos  int
}

func (r *zoxideDBReader) next(n int) ([]byte, error) {
	// n is either a constant width (4/8) or a length already bounds-checked
	// against the remaining data, so it can never be negative here.
	if n > len(r.data)-r.pos {
		return nil, io.ErrUnexpectedEOF
	}

	b := r.data[r.pos : r.pos+n]
	r.pos += n
	return b, nil
}

func (r *zoxideDBReader) uint32() (uint32, error) {
	b, err := r.next(4)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(b), nil
}

func (r *zoxideDBReader) uint64() (uint64, error) {
	b, err := r.next(8)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(b), nil
}

func (r *zoxideDBReader) float64() (float64, error) {
	b, err := r.uint64()
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(b), nil
}

func (r *zoxideDBReader) str() (string, error) {
	length, err := r.uint64()
	if err != nil {
		return "", err
	}

	// Guard against a corrupt length before slicing for it.
	if length > uint64(len(r.data)-r.pos) {
		return "", fmt.Errorf("importer: zoxide entry path length %d exceeds remaining data", length)
	}

	b, err := r.next(int(length))
	if err != nil {
		return "", err
	}
	return string(b), nil
}
