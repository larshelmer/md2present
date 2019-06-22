#SHELL := /bin/bash
VERSION := 0.0.1
BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2> /dev/null)
REVISION := $(shell git rev-parse HEAD 2> /dev/null)
CREATED := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

.PHONY: build
build:
	@go build -ldflags "-X main.Branch=$(BRANCH) \
		-X main.Version=$(VERSION) \
		-X main.Revision=$(REVISION) \
		-X main.Created=$(CREATED)" \
		-o md2present cmd/md2present/main.go

.PHONY: clean
clean:
	rm -f md2present
