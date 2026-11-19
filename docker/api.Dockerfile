# build from the repo root: docker build -f docker/api.Dockerfile -t price-api .
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY api/go.mod api/go.sum ./
RUN go mod download

COPY api/ .

RUN go build -o api .


FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/api .

# configuration comes from the environment: PG_URL, PORT (default 8080)
EXPOSE 8080

CMD ["./api"]
