package importer

import (
	"bytes"
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
	content, err := readConfig(i.configPaths)
	if err != nil {
		return nil, err
	}

	return parseZoxideDB([]byte(content))
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
func parseZoxideDB(data []byte) (scoring.Entries, error) {
	r := bytes.NewReader(data)

	var version uint32
	if err := binary.Read(r, binary.LittleEndian, &version); err != nil {
		return nil, fmt.Errorf("importer: cannot read zoxide database version: %w", err)
	}
	if version != zoxideVersion {
		return nil, fmt.Errorf("importer: unsupported zoxide database version %d, jump supports %d", version, zoxideVersion)
	}

	var count uint64
	if err := binary.Read(r, binary.LittleEndian, &count); err != nil {
		return nil, fmt.Errorf("importer: cannot read zoxide entry count: %w", err)
	}

	var entries scoring.Entries

	for n := uint64(0); n < count; n++ {
		entry, err := readZoxideEntry(r)
		if err != nil {
			return nil, err
		}

		if _, found := entries.Find(entry.Path); found {
			continue
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

func readZoxideEntry(r *bytes.Reader) (*scoring.Entry, error) {
	var pathLen uint64
	if err := binary.Read(r, binary.LittleEndian, &pathLen); err != nil {
		return nil, fmt.Errorf("importer: cannot read zoxide entry path length: %w", err)
	}

	// Guard against a corrupt length before allocating for it.
	if pathLen > uint64(r.Len()) {
		return nil, fmt.Errorf("importer: zoxide entry path length %d exceeds remaining data", pathLen)
	}

	path := make([]byte, pathLen)
	if _, err := io.ReadFull(r, path); err != nil {
		return nil, fmt.Errorf("importer: cannot read zoxide entry path: %w", err)
	}

	var rank float64
	if err := binary.Read(r, binary.LittleEndian, &rank); err != nil {
		return nil, fmt.Errorf("importer: cannot read zoxide entry rank: %w", err)
	}

	var lastAccessed uint64
	if err := binary.Read(r, binary.LittleEndian, &lastAccessed); err != nil {
		return nil, fmt.Errorf("importer: cannot read zoxide entry timestamp: %w", err)
	}

	return &scoring.Entry{
		Path: string(path),
		Score: &scoring.Score{
			Weight: int64(math.Round(rank)),
			Age:    time.Unix(int64(lastAccessed), 0),
		},
	}, nil
}
