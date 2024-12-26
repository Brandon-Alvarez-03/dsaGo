.PHONY: init build run clean

# Initialize Go module
init:
	go mod init dsaGo
	go mod tidy

# Build the project (creates a binary)
build:
	go build -o dsaApp main.go

# Run the project
run:
	go run main.go

# Clean up binaries and temporary files
clean:
	rm -f dsaApp
