commit = $(shell git describe --always --abbrev=7 --dirty)

.DEFAULT_GOAL := build
test:
	go test ./...
lint:
	golangci-lint run 

build: lint test
	go build -ldflags="-X 'github.com/jigargandhi/to/version.version=v0.2' -X 'github.com/jigargandhi/to/version.commit=$(commit)'"