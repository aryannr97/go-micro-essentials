GOCMD=$(shell echo go)
GOLINT=$(shell echo golangci-lint)

fmt:
	@echo "+ $@"
	@$(GOCMD) fmt ./...

lint: 
	@echo "+ $@"
	@${GOLINT} run --disable errcheck

test:
	@echo "+ $@"
	@$(GOCMD) test ./... -race -v -coverprofile=coverage.out -covermode=atomic

build:
	@echo "+ $@"
	@$(GOCMD) build ./...

all: fmt lint test build
