# price_intelligence_platform

An early-stage (2025) **price intelligence** prototype in Go: crawlers collect product prices for
tracked keywords, and an API manages those keywords in Postgres. Both parts build and pass
`make check`; the crawlers aren't yet wired to the API's keywords (they search `laptop`).

## Quick start

```sh
make check                                       # gofmt, vet and tests for both modules
psql "$PG_URL" -f api/schema.sql                 # once: the keywords table
cd api && PG_URL=postgres://… go run .           # API on :8080 (PORT to change)
cd crawlers && go run .                          # every registered parser for "laptop"
make images                                      # Docker images price-api, price-crawlers
```

## API (`api/`)

| Method | Path | Body | Answer |
|---|---|---|---|
| GET | `/keywords` | | enabled keywords: `id`, `value`, `enabled`, `created_at`, `updated_at` |
| POST | `/keywords` | `{"value": "laptop"}` | `201`, or `400` for empty/invalid input |

gin + pgx connection pool (`PG_URL`), `.env` loaded when present, schema in `api/schema.sql`.

## Crawlers (`crawlers/`)

- `common/interfaces.go`: the `Parser` interface (`Name()`, `Parse(keyword)`) and the
  `Result` / `Variant` types (title, prices, stock, properties).
- `registry/registry.go`: the list of parsers; `main.go` runs them in parallel and logs the
  result counts.
- `amazon/crawler.go`: an Amazon search + product-page scraper built on colly. Scraping Amazon
  may break their terms and gets blocked quickly; keep request rates low.

## Docker

`docker/api.Dockerfile` and `docker/crawlers.Dockerfile`, built from the repo root
(`make images`): multi-stage Go → Alpine. Configuration comes from the environment.

Licensed under Apache-2.0. Decisions: [docs/adr](docs/adr/README.md).
