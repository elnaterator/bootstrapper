.DEFAULT_GOAL := help
.PHONY: help dev

build: ## build
	go build -o bin/command cmd/command/main.go

test: ## run unit tests
	go test ./... -v

help: ## print help (default)
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
