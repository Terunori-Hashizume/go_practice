.PHONY: all fmt vet clean build install

BIN_NAME := prob-init
LOCAL_BIN := $(HOME)/.local/bin/$(BIN_NAME)

all: install

fmt:
	go fmt ./...

vet:
	go vet ./...

clean:
	rm -f $(BIN_NAME)

build:
	go build -o $(BIN_NAME)

install: fmt vet clean build
	mkdir -p $(HOME)/.local/bin
	mv $(BIN_NAME) $(LOCAL_BIN)
