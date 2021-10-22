commit = $(shell git describe --always --abbrev=7 --dirty)

.DEFAULT_GOAL := build
test:
	go test ./...
build: test
	go build -ldflags="-X 'github.com/jigargandhi/to/version.version=v0.1' -X 'github.com/jigargandhi/to/version.commit=$(commit)'"