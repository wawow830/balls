# Releasing

## Prepare

- [ ] Confirm `main` is clean and current.
- [ ] Run `gofmt` and verify that it produces no changes.
- [ ] Run `go test -race ./...`.
- [ ] Run the benchmark suite and review material regressions.
- [ ] Compare generated records with the committed compatibility fixtures.
- [ ] Validate `schema/record.schema.json`.
- [ ] Review public API and specification changes for compatibility.
- [ ] Update README examples and package documentation.
- [ ] Summarize user-visible changes in the release notes.

## Publish

- [ ] Create an annotated semantic-version tag.
- [ ] Push the tag to the canonical repository.
- [ ] Build release artifacts from the tagged commit.
- [ ] Record artifact checksums in the release.
- [ ] Verify a clean consumer can install the package and command.

## Verify

- [ ] Run the documented example against the released version.
- [ ] Confirm generated output matches the tagged specification.
- [ ] Open follow-up issues for deferred or discovered work.
