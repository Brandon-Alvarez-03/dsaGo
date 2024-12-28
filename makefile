.PHONY: init build run clean test benchmark run-sort run-search run-ds docs

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
BINARY_NAME=dsaApp

# Initialize Go module
init:
	$(GOCMD) mod init dsaGo
	$(GOCMD) mod tidy

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME) -v main.go

# Run all demonstrations
run:
	$(GORUN) main.go

# Run specific demonstrations
run-sort:
	$(GORUN) main.go -demo sort

run-search:
	$(GORUN) main.go -demo search

run-ds:
	$(GORUN) main.go -demo ds

# Run tests
test:
	$(GOTEST) -v ./...

# Run benchmarks
benchmark:
	$(GOTEST) -bench=. ./...

# Generate documentation
docs:
	godoc -http=:6060

# Clean up binaries and temporary files
clean:
	rm -f $(BINARY_NAME)
	rm -f *.out
	find . -type f -name '*.test' -delete

# Install required tools
tools:
	$(GOCMD) install golang.org/x/tools/cmd/godoc@latest

# Check code formatting
fmt:
	$(GOCMD) fmt ./...

# Run all static analysis tools
lint: fmt
	golangci-lint run

# Full development setup
setup: tools init lint build test

# Help command
help:
	@echo "Available commands:"
	@echo "  make init       - Initialize Go module"
	@echo "  make build      - Build the project"
	@echo "  make run        - Run all demonstrations"
	@echo "  make run-sort   - Run sorting demonstrations"
	@echo "  make run-search - Run searching demonstrations"
	@echo "  make run-ds     - Run data structure demonstrations"
	@echo "  make test       - Run tests"
	@echo "  make benchmark  - Run benchmarks"
	@echo "  make clean      - Clean up binaries"
	@echo "  make docs       - Start documentation server"
	@echo "  make fmt        - Format code"
	@echo "  make lint       - Run linters"
	@echo "  make setup      - Complete development setup"