# price_intelligence_platform

A 2025 Go prototype: product-price crawlers plus a keyword API on Postgres.

## Commands
```sh
make check     # gofmt check, go vet and go test for api/ and crawlers/ (run before every PR)
make build     # both modules
make images    # Docker images, built from the repo root
psql "$PG_URL" -f api/schema.sql   # the keywords table
```

## Layout
- `api/` (module `api`, Go 1.25 for gin): `main.go`, `db/` (pgxpool from `PG_URL`, keyword
  queries), `handler/` (gin handlers + test), `router/`, `models/`, `schema.sql`.
- `crawlers/` (module `crawlers`, Go 1.23): `common/` interfaces and types, `registry/` the
  parser list, `amazon/` the colly scraper, `main.go` the runner.
- `docker/`: one Dockerfile per service; build context is the repo root.
- `docs/adr/`: decision records (`.adr`), index in `docs/adr/README.md`.

## Conventions
- Read the ADR before changing a recorded decision; commit subjects carry its number.
- Two separate Go modules; `make` targets run in each.
- Use the pgx pool, never one shared `pgx.Conn` (gin serves requests concurrently).
- A new store = a new package implementing `common.Parser`, added to `registry.All()`.
- Scraping Amazon may break their terms and gets blocked quickly; keep request rates low.
