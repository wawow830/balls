# Package API

Import the package with:

```go
import "github.com/wawow830/balls/sequence"
```

## `New(seed uint64) *Generator`

Creates an independent deterministic generator. All seeds, including zero, are valid. A generator starts at index zero.

## `(*Generator).Next() Record`

Advances the generator and returns the next record. A generator is intentionally state-local; do not call `Next` concurrently on the same instance. Give each goroutine its own generator when generating independent streams.

## `Verify(record Record) bool`

Recomputes the checksum from a record's index and value. It detects accidental record corruption but is not a cryptographic authenticity check.

## `Record`

A record contains an unsigned index, an unsigned generated value, and a lowercase 16-character checksum. Its field tags are stable for JSON encoding.

## Stability

The sequence and checksum are compatibility-sensitive. See [the reproducibility specification](specification.md) for the normative algorithm. New APIs may be added without changing existing output.
