-- keywords the crawlers track; apply once: psql "$PG_URL" -f api/schema.sql
CREATE TABLE IF NOT EXISTS keywords (
    id         SERIAL PRIMARY KEY,
    value      TEXT        NOT NULL UNIQUE,
    enabled    BOOLEAN     NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
