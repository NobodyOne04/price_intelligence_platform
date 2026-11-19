# Changelog

## Unreleased

### Added
- README, CLAUDE.md, changelog and ADRs
- `api/schema.sql` for the keywords table; a handler test; `make check`, `build`, `test`, `images`

### Fixed
- The API didn't compile (empty `go.mod`, typos, undefined names); it builds, uses a connection pool, and returns 201/400 properly
- Crawlers ran one after another (`wg.Wait()` inside the loop); now in parallel
- Dockerfiles: the API image copied a missing `config.env`, and the crawler image copied from outside its build context
