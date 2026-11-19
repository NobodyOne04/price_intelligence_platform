# build from the repo root: docker build -f docker/crawlers.Dockerfile -t price-crawlers .
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY crawlers/go.mod crawlers/go.sum ./
RUN go mod download

COPY crawlers/ .

RUN go build -o crawler .

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/crawler .

CMD ["./crawler"]
