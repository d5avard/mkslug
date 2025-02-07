.DEFAULT_GOAL := build

build:
	go build -o ./bin/ ./cmd/mkslug/...

test:
	go test -v ./...

install: build
	cp ./bin/mkslug /usr/local/bin/mkslug