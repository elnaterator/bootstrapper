.DEFAULT_GOAL := help
.PHONY: help dev

build: ## build
	go build -o bin/bootstrap cmd/bootstrap/main.go

install: ## install the binary to ~/.local/bin/
	cp ./bin/bootstrap ~/.local/bin/bootstrap

test: ## run unit tests
	go test ./... -v

run: ## run
	go run cmd/bootstrap/main.go

help: ## print help (default)
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
