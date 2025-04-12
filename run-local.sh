#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Build the Go application
echo "Building the Go application..."
go build -o my-go-app ./cmd

# Run the built application
echo "Running the Go application..."
./my-go-app 2 2