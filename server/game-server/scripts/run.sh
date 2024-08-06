#!/bin/bash

set -e

if [ "$1" == "--build" ]; then
    ./build.sh
    if [ $? -ne 0 ]; then
        echo "Build failed, not starting the server."
        exit 1
    fi
fi

echo "Starting the game server..."
./bin/game-server