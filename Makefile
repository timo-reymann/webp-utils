SHELL := /bin/bash
.PHONY: help

help: ## Display this help page
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[33m%-30s\033[0m %s\n", $$1, $$2}'

install-packr: ## Install packr tool
	@go get -u github.com/gobuffalo/packr/packr

bundle-schemas: ## Bundle schemas using packr
	@packr

test-coverage-report: ## Run test and display coverage report in browser
	@go test -covermode=count -coverprofile=/tmp/count.out -v ./...
	@go tool cover -html=/tmp/count.out

