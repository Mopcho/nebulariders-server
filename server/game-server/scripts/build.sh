#!/bin/bash

set -e

echo "Building the game server..."
go build -o bin/game-server ./cmd/game-server

echo "Build completed successfully."