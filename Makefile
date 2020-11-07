.PHONY: help

build: # @HELP build the app
build:
	go get -v -t -d ./...
	go build -o ./snippetbox -v ./cmd/web 

test: # @HELP run tests
test:
	go test -v ./cmd/web

help: # @HELP prints this message
help:
	@echo "VARIABLES:"
	@echo "  BINS = $(BINS)"
	@echo "  OS = $(OS)"
	@echo "  ARCH = $(ARCH)"
	@echo "  REGISTRY = $(REGISTRY)"
	@echo
	@echo "TARGETS:"
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST)    \
	    | awk '                                   \
	        BEGIN {FS = ": *# *@HELP"};           \
	        { printf "  %-30s %s\n", $$1, $$2 };  \
	    '