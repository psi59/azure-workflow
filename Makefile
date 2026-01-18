.PHONY: build test clean install package version-sync sign

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

install: package
	xattr -d com.apple.quarantine azure-workflow.alfredworkflow 2>/dev/null || true
	open azure-workflow.alfredworkflow

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

# Ad-hoc code signing for macOS Gatekeeper compatibility
sign:
	codesign --sign - --force $(BINARY_NAME)

# Alfred workflow package
package: build-universal version-sync sign
	zip -r azure-workflow.alfredworkflow $(BINARY_NAME) services.yaml info.plist icons/ icon.png

run:
	./$(BINARY_NAME) "$(QUERY)"
