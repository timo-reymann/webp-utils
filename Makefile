SHELL := /bin/bash
BUILD_ARGS=-ldflags "-X github.com/timo-reymann/webp-utils/pkg/buildinfo.GitSha=$(shell git rev-parse --short HEAD) -X github.com/timo-reymann/webp-utils/pkg/buildinfo.Version=$(shell git describe --abbrev=0)"
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

create-dist: ## Create dist folder if not already existent
	@mkdir -p dist/

build-linux: ## Build for linux (amd64, arm64)
	GOOS=linux GOARCH=amd64 go build -o dist/webp-utils-linux-amd64 $(BUILD_ARGS)
	@GOOS=linux GOARCH=arm64 go build -o dist/webp-utils-linux-arm64 $(BUILD_ARGS)

build-windows: ## Build for windows (amd64)
	@GOOS=windows GOARCH=amd64 go build -o dist/webp-utils-windows-amd64.exe $(BUILD_ARGS)

build-darwin: ## Build for macOS (amd64, arm64)
	@GOOS=darwin GOARCH=amd64 go build -o dist/webp-utils-darwin-amd64 $(BUILD_ARGS)
	@GOOS=darwin GOARCH=amd64 go build -o dist/webp-utils-darwin-arm64 $(BUILD_ARGS)

build: create-dist build-linux build-darwin build-windows ## Build for all platform
