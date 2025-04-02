#!/bin/bash
set -e

DIST_DIR="dist"

# Clean and recreate dist/
rm -rf "$DIST_DIR"
mkdir -p "$DIST_DIR"

echo "🔧 Building binaries into $DIST_DIR..."

# Build Linux binary
GOOS=linux GOARCH=amd64 go build -o "$DIST_DIR/indiandata-linux-amd64"

# Build macOS Intel binary
GOOS=darwin GOARCH=amd64 go build -o "$DIST_DIR/indiandata-darwin-amd64"

# Build macOS Apple Silicon binary
GOOS=darwin GOARCH=arm64 go build -o "$DIST_DIR/indiandata-darwin-arm64"

# Zip them
echo "📦 Zipping binaries..."
cd "$DIST_DIR"
zip indiandata-linux-amd64.zip indiandata-linux-amd64
zip indiandata-darwin-amd64.zip indiandata-darwin-amd64
zip indiandata-darwin-arm64.zip indiandata-darwin-arm64
cd ..

echo "✅ Done. Artifacts are in ./$DIST_DIR"
