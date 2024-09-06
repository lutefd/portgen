# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=portgen
BINARY_UNIX=$(BINARY_NAME)_unix

BUILD_DIR=./dist

all: test build

build:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v ./cmd/portgen

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

run:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v ./cmd/portgen
	./$(BUILD_DIR)/$(BINARY_NAME)

deps:
	$(GOGET) -v -t -d ./...

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_UNIX) -v ./cmd/portgen


.PHONY: all build test clean run deps build-linux
