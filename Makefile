### General
SHELL := /bin/bash
ROOT := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))

include $(ROOT)/tools.mk

### Commands
.PHONY: initialize
initialize: initialize-once install-tools

.PHONY: 
initialize-once:
	brew install mise
	mise i
	bun i

### Development Commands
.PHONY: build
build:
	bunx nx affected -t build

.PHONY: fmt
fmt:
	bunx nx affected -t fmt

.PHONY: lint
lint:
	bunx nx affected -t lint

.PHONY: test
test:
	bunx nx affected -t test

.PHONY: graph
graph:
	bunx nx graph

### Go Commands
.PHONY: fmt-go
fmt-go:
	$(GO_ENV) $(GO) mod tidy
	mise exec -- golangci-lint fmt ./...

.PHONY: lint-go
lint-go: $(BIN)/govulncheck-$(GOVULNCHECK_VERSION)
	mise exec -- golangci-lint run ./... -v -c $(ROOT)/go/.golangci.yaml ./...
	$(BIN)/govulncheck ./...

.PHONY: test-go
test-go:
	$(GO_ENV) $(GO) test -v ./...
