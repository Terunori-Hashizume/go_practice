.DEFAULT_GOAL := run

.PHONY: fmt vet clean build run

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

clean: vet
	rm -f hello

build: clean
	go build -o hello .

run: build
	./hello