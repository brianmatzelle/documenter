#!/bin/bash

# Exit on error
set -e

# Function to cleanup on exit
cleanup() {
    echo "Cleaning up..."
    docker compose -f docker-compose.dev.yml down
}

# Register cleanup function
trap cleanup EXIT

# Check if docker is running
if ! docker info > /dev/null 2>&1; then
    echo "Docker is not running. Please start Docker first."
    exit 1
fi

# Build and start the development environment
echo "Starting development environment..."
docker compose -f docker-compose.dev.yml up --build
