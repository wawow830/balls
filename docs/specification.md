# Reproducibility specification

## State

A generator contains two unsigned 64-bit integers: `state`, initialized from the seed, and `index`, initialized to zero. Arithmetic wraps modulo 2^64.

## Transition

For each record:

1. Add `0x9e3779b97f4a7c15` to `state`.
2. Set `value` to `state`.
3. Replace `value` with `(value XOR (value >> 30)) * 0xbf58476d1ce4e5b9`.
4. Replace `value` with `(value XOR (value >> 27)) * 0x94d049bb133111eb`.
5. Replace `value` with `value XOR (value >> 31)`.
6. Emit the current `index`, the resulting `value`, and its checksum.
7. Increment `index`.

Right shifts are logical. Multiplication and addition wrap to 64 bits.

## Checksum

The checksum input is 16 bytes: the index followed by the value, both encoded as little-endian unsigned 64-bit integers. Apply FNV-1a-64 and encode the result as 16 lowercase hexadecimal characters.

## Compatibility

For a fixed seed, implementations conforming to this document must emit identical records in identical order. Changes to the state transition, byte order, checksum, or JSON field meanings require a new algorithm version.
