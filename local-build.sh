#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Delete the existing Go application build if it exists
if [ -f my-application ]; then
    echo "Deleting existing Go application build..."
    rm my-application
fi

# Build the Go application
echo "Building the Go application..."
go build -o my-application ./src
