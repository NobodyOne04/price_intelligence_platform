# price_intelligence_platform

An early-stage (June 2025) **price intelligence** prototype in Go: crawlers collect product
prices for tracked keywords, and an API manages the keywords in Postgres.

**Status: work in progress.** The crawlers build and run; the API does not compile yet (see
below), and the Makefile is empty.

## Crawlers (`crawlers/`, builds)

```sh
cd crawlers && go run .        # runs every registered parser for the keyword "laptop"
```

- `common/interfaces.go`: the `Parser` interface (`Name()`, `Parse(keyword)`) and the
  `Result` / `Variant` types (title, prices, stock, properties).
- `registry/registry.go`: the list of parsers; `main.go` runs each and logs the result count.
- `amazon/crawler.go`: an Amazon search + product-page scraper built on colly.

## API (`api/`, doesn't compile yet)

Planned: a gin server with `GET` / `POST /keywords` over a Postgres `keywords` table
(`id`, `value`, `enabled`, timestamps), connection from `PG_URL`, port from `PORT` (default
8080). To get it building: fill `api/go.mod` (it's empty), and fix the typos and undefined
names in `main.go`, `db/`, `handler/` and `router/` (for example `db.ConnWrapper(Conn: conn)`,
`fmt.Error`, `nill`, `Keywards`).

## Docker

`docker/api.Dockerfile` and `docker/crawlers.Dockerfile`: multi-stage Go 1.23 → Alpine builds.

Licensed under Apache-2.0. Decisions: [docs/adr](docs/adr/README.md).
