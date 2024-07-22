#!/usr/bin/env bash
set -e

# Temporary file to capture Tailwind CLI output
TEMP_FILE=$(mktemp)

# Start Tailwind CSS generation with --watch in the background and redirect output to temp file
tailwindcss -i ui/src/css/main.css -o ui/static/css/tailwind.css --watch=always > "$TEMP_FILE" 2>&1 &

# Capture the PID of the Tailwind process
TAILWIND_PID=$!

# Function to clean up on exit
cleanup() {
	kill $TAILWIND_PID
	rm -f "$TEMP_FILE"
}

# Register cleanup function to be called on the script's exit
trap cleanup EXIT

# Wait for Tailwind to finish the first build
echo "Waiting for Tailwind CSS build to complete..."
while ! grep -q "Done in" "$TEMP_FILE"; do
	sleep 1
done

echo "Tailwind CSS build completed. Task output:"
cat "$TEMP_FILE"
echo ""

# Start the Go server
go run cmd/server.go -loglevel=info