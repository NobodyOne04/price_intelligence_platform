# price_intelligence_platform: run `make help`.
.DEFAULT_GOAL := help
.PHONY: help build test check images

help: ## Show this help
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z_-]+:.*##/ {printf "  \033[36m%-8s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build both modules
	cd api && go build ./...
	cd crawlers && go build ./...

test: ## Run the tests
	cd api && go test ./...
	cd crawlers && go test ./...

check: ## gofmt check, vet and tests for both modules
	@test -z "$$(gofmt -l api crawlers)" || { echo "gofmt needed:"; gofmt -l api crawlers; exit 1; }
	cd api && go vet ./... && go test ./...
	cd crawlers && go vet ./... && go test ./...

images: ## Build the Docker images (from the repo root)
	docker build -f docker/api.Dockerfile -t price-api .
	docker build -f docker/crawlers.Dockerfile -t price-crawlers .
