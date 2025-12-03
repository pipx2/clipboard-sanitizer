#!/bin/bash

# Path to the main package
MAIN_PKG="./cmd/sanitizer/"

# Output directory
OUTPUT_DIR="./build"
if [ -z "$OUTPUT_DIR" ]; then
    rm -rf "$OUTPUT_DIR"
fi
mkdir -p "$OUTPUT_DIR"

echo "Building clipboard-sanitizer..."
go build -o "$OUTPUT_DIR"/clipboard-sanitizer "$MAIN_PKG"
echo "Binaries are found in $OUTPUT_DIR"
