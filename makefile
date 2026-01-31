GO := go
GOFMT := gofmt
DIRECTORIES_TO_TEST := $$(go list ./... | grep -v /internal/ | grep -v /types/)

MAKEFLAGS += --silent

.PHONY: help

help: # Show available make commands
	@awk 'BEGIN {FS = ":.*#"} /^[a-zA-Z0-9_-]+:.*#/ { printf "  %-15s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

test: # Run all tests
	@echo -e "\e[0;32mRunning tests...\e[0m"
	$(GO) test -cover $(DIRECTORIES_TO_TEST)

lint: # Format source files
	$(GOFMT) -d -w .
