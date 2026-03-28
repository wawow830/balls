# balls

A small deterministic data generator for fixtures, demos, reproducible tests, and lightweight workload simulation.

## Run

Requires Go 1.22 or newer.

```sh
go run ./cmd/balls -count 5 -seed 42
```

Each line is an independent JSON record:

```json
{"index":0,"value":13679457532755275413,"checksum":"fc9b03ef7e01ba84"}
{"index":1,"value":2949826092126892291,"checksum":"ca74bae04722c51d"}
```

The same seed and count always produce the same stream. Output can be redirected directly into a fixture:

```sh
go run ./cmd/balls -count 1000 -seed 42 > testdata.jsonl
```

## Library

The generator is also available as a small Go package:

```go
import "github.com/wawow830/balls/sequence"

generator := sequence.New(42)
record := generator.Next()
```

Records include a monotonic index, a deterministic 64-bit value, and a compact checksum. The package is state-local and safe to use concurrently when each goroutine owns its own generator.

## Test

```sh
go test ./...
```

## License

MIT
