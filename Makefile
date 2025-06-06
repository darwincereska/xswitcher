# Project Variables
PROJECT_NAME := xswitcher
VERSION := 0.1.2
BUILD_TIME := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
GOARCH := $(shell go env GOARCH)
GOOS := $(shell go env GOOS)
DESTDIR ?= 
PREFIX ?= /usr
BINDIR ?= $(PREFIX)/bin

# Go Variables
GO ?= go
GOFMT ?= gofmt
GOLINT ?= golint
GOCILINT ?= golangci-lint

# Build Flags
LDFLAGS := -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)"

# Directories
BIN_DIR := bin
DIST_DIR := dist

.PHONY: all build clean test lint fmt vet install uninstall

all: build

# Build the Project
build: fmt vet
	@mkdir -p $(BIN_DIR)
	$(GO) build $(LDFLAGS) -o $(BIN_DIR)/$(PROJECT_NAME) ./main.go

# Run Tests
test: 
	$(GO) test -v ./...

# Run lint
lint:
	$(GOLINT) ./...
	$(GOCILINT) run

# Format Code
fmt:
	$(GOFMT) -w -s .

# Run Vet
vet:
	$(GO) vet ./...

# Clean Build Artifacts
clean:
	rm -rf $(BIN_DIR) $(DIST_DIR) pkg src $(PROJECT_NAME)-$(VERSION).tar.gz $(PROJECT_NAME)-$(VERSION)-x86_64.pkg.tar.zst
	rm -rf $(PROJECT_NAME)-$(VERSION)-1-x86_64.pkg.tar.zst $(PROJECT_NAME)-debug-$(VERSION)-1-x86_64.pkg.tar.zst
	$(GO) clean

# Install system-wide
install: build
	install -Dm755 $(BIN_DIR)/$(PROJECT_NAME) $(DESTDIR)$(BINDIR)/$(PROJECT_NAME)

# Uninstall
uninstall:
	rm -f $(DESTDIR)$(BINDIR)/$(PROJECT_NAME)

# Cross-compilation targets
build-linux:
	GOOS=linux GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(BIN_DIR)/$(PROJECT_NAME)-linux-amd64 ./main.go

build-darwin:
	GOOS=darwin GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(BIN_DIR)/$(PROJECT_NAME)-darwin-amd64 ./main.go

build-windows:
	GOOS=windows GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(BIN_DIR)/$(PROJECT_NAME)-windows-amd64.exe ./main.go

# Package for distribution
package: build-linux build-darwin build-windows
	@mkdir -p $(DIST_DIR)
	tar -czvf $(DIST_DIR)/$(PROJECT_NAME)-$(VERSION)-linux-amd64.tar.gz -C $(BIN_DIR) $(PROJECT_NAME)-linux-amd64
	tar -czvf $(DIST_DIR)/$(PROJECT_NAME)-$(VERSION)-darwin-amd64.tar.gz -C $(BIN_DIR) $(PROJECT_NAME)-darwin-amd64
	zip -j $(DIST_DIR)/$(PROJECT_NAME)-$(VERSION)-windows-amd64.zip $(BIN_DIR)/$(PROJECT_NAME)-windows-amd64.exe

# Dependency Management
deps:
	$(GO) mod download

tidy: 
	$(GO) mod tidy

# For Arch Linux specific (optional)
install-arch-deps: 
	sudo pacman -S --needed go base-devel git

# For Debian specific (optional)
install-debian-deps:
	sudo apt-get install -y golang-go golang-golang-x-tools git make gcc