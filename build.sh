#!/bin/sh

# Set output directory
OUTPUT_DIR="./build/bin"
OUTPUT_FILE="$OUTPUT_DIR/main"

# Ensure the output directory exists
# 'mkdir -p' is supported by busybox (Alpine's toolset)
mkdir -p "$OUTPUT_DIR"

# Remove the existing build file if it exists
# '[ -f ... ]' is a standard POSIX test command
if [ -f "$OUTPUT_FILE" ]; then
    rm "$OUTPUT_FILE"
    echo "Removed existing build file: $OUTPUT_FILE"
fi

# Build the Go application for Linux
echo "Building Go application for Linux..."
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o "$OUTPUT_FILE" ./cmd/app/main.go

# Check if the build was successful
# '$?' is the standard way to get the exit code in any POSIX shell
if [ $? -ne 0 ]; then
    echo "Go build failed. Exiting."
    exit 1
fi

echo "Build successful: $OUTPUT_FILE"