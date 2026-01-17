.PHONY: build test clean install package version-sync

BINARY_NAME=azure-workflow
VERSION=$(shell cat VERSION)
WORKFLOW_DIR=~/Library/Application\ Support/Alfred/Alfred.alfredpreferences/workflows/azure-workflow

build:
	go build -o $(BINARY_NAME) ./cmd/main.go

test:
	go test -v ./...

test-coverage:
	go test -v -cover ./...

clean:
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME)-arm64
	rm -f $(BINARY_NAME)-amd64
	rm -f azure-workflow.alfredworkflow

install: build
	mkdir -p $(WORKFLOW_DIR)
	cp $(BINARY_NAME) $(WORKFLOW_DIR)/
	cp services.yaml $(WORKFLOW_DIR)/
	cp info.plist $(WORKFLOW_DIR)/
	cp -r icons $(WORKFLOW_DIR)/

# Build flags for release
LDFLAGS=-s -w
BUILDFLAGS=-trimpath -ldflags "$(LDFLAGS)"

# Universal binary for Apple Silicon & Intel
build-universal:
	GOOS=darwin GOARCH=arm64 go build $(BUILDFLAGS) -o $(BINARY_NAME)-arm64 ./cmd/main.go
	GOOS=darwin GOARCH=amd64 go build $(BUILDFLAGS) -o $(BINARY_NAME)-amd64 ./cmd/main.go
	lipo -create -output $(BINARY_NAME) $(BINARY_NAME)-arm64 $(BINARY_NAME)-amd64
	rm $(BINARY_NAME)-arm64 $(BINARY_NAME)-amd64

# Sync VERSION to info.plist
version-sync:
	@plutil -replace version -string "$(VERSION)" info.plist

# Alfred workflow package
package: build-universal version-sync
	zip -r azure-workflow.alfredworkflow $(BINARY_NAME) services.yaml info.plist icons/ icon.png

run:
	./$(BINARY_NAME) "$(QUERY)"
