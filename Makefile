#SHELL := /bin/bash
VERSION := 0.0.1
#BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2> /dev/null)
REVISION := $(shell git rev-parse HEAD 2> /dev/null)
CREATED := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

GOFILES = $(shell find . -type f -name '*.go')

.PHONY: build
build:
	@go build -ldflags "-X main.version=$(VERSION) \
		-X main.revision=$(REVISION) \
		-X main.created=$(CREATED)" \
		-o md2present cmd/md2present/main.go

.PHONY: clean
clean:
	rm -f md2present

.PHONY: lint
lint:
	@gofmt -s -w $(GOFILES)
	@goimports -w $(GOFILES)
	@golangci-lint run --enable-all
