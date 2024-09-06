# Include go binaries into path
export PATH := $(GOPATH)/bin:$(PATH)

CURRENT_BRANCH_NAME := $(shell git rev-parse --abbrev-ref HEAD)
DATA_PATH := $(shell pwd)/data
BIN := $(CURDIR)/bin/
SOURCE_PATH := GOBIN=$(BIN) DATA_PATH=$(DATA_PATH) CURDIR=$(shell pwd) CURRENT_BRANCH_NAME=$(CURRENT_BRANCH_NAME)

install: mod

mod-action-%:
	@echo "Run go mod ${*}...."
	GOBIN=$(BIN) GO111MODULE=on go mod $*
	@echo "Done go mod  ${*}"

mod: mod-action-verify mod-action-tidy mod-action-vendor mod-action-download mod-action-verify ## Download all dependencies

test: clean-cache-test ## run all tests
	$(SOURCE_PATH) go test ./postsort/... -race -v -coverprofile coverage.out
	go tool cover -html=coverage.out -o coverage.html
	rm coverage.out

clean-cache-test: ## clean cache
	@echo "Cleaning test cache..."
	$(SOURCE_PATH) go clean -testcache



