# price_intelligence_platform

A June 2025 Go prototype: product-price crawlers plus a keyword API. Work in progress.

## Commands
```sh
cd crawlers && go build ./... && go run .   # works
cd api && go build ./...                    # fails: empty go.mod and compile errors (see README)
```
The root `Makefile` is empty. No tests.

## Layout
- `crawlers/` (module `crawlers`, Go 1.23): `common/` interfaces and types, `registry/` the
  parser list, `amazon/` the colly scraper, `main.go` the runner.
- `api/` (module `api`, unfinished): `main.go`, `db/` (pgx, `PG_URL`), `handler/`, `router/` (gin), `models/`.
- `docker/`: one Dockerfile per service.
- `docs/adr/`: decision records (`.adr`), index in `docs/adr/README.md`.

## Conventions
- Read the ADR before changing a recorded decision; commit subjects carry its number.
- Two separate Go modules; build each from its own folder.
- A new store = a new package implementing `common.Parser`, added to `registry.All()`.
- Scraping Amazon may break their terms and gets blocked quickly; keep request rates low.
