#!/bin/bash

# Set output directory
OUTPUT_DIR="./.dist/bin"
OUTPUT_FILE="$OUTPUT_DIR/main"

# Ensure the output directory exists
mkdir -p "$OUTPUT_DIR"

# Remove the existing build file if it exists
if [ -f "$OUTPUT_FILE" ]; then
    rm "$OUTPUT_FILE"
    echo "Removed existing build file: $OUTPUT_FILE"
fi

# Build the Go application for Linux
GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -ldflags="-s -w" -o "$OUTPUT_FILE" ./cmd/app/main.go

# Check if the build was successful
if [ $? -ne 0 ]; then
    echo "Go build failed. Exiting."
    exit 1
fi

echo "Build successful: $OUTPUT_FILE"