// Package sequence provides a deterministic stream of self-checking records.
package sequence

import (
	"encoding/binary"
	"encoding/hex"
	"hash/fnv"
)

// Record is one item in a generated stream.
type Record struct {
	Index    uint64 `json:"index"`
	Value    uint64 `json:"value"`
	Checksum string `json:"checksum"`
}

// Generator produces deterministic records from a seed.
type Generator struct {
	state uint64
	index uint64
}

// New creates a generator. Every seed, including zero, is valid.
func New(seed uint64) *Generator {
	return &Generator{state: seed}
}

// Next returns the next record in the stream.
func (g *Generator) Next() Record {
	g.state += 0x9e3779b97f4a7c15
	value := g.state
	value = (value ^ (value >> 30)) * 0xbf58476d1ce4e5b9
	value = (value ^ (value >> 27)) * 0x94d049bb133111eb
	value ^= value >> 31

	record := Record{
		Index: g.index,
		Value: value,
	}
	record.Checksum = checksum(record.Index, record.Value)
	g.index++
	return record
}

// Verify reports whether a record's checksum matches its index and value.
func Verify(record Record) bool {
	return record.Checksum == checksum(record.Index, record.Value)
}

func checksum(index, value uint64) string {
	var payload [16]byte
	binary.LittleEndian.PutUint64(payload[:8], index)
	binary.LittleEndian.PutUint64(payload[8:], value)
	hash := fnv.New64a()
	_, _ = hash.Write(payload[:])
	return hex.EncodeToString(hash.Sum(nil))
}
