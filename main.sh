#!/bin/bash

# Build the Go project
go build

# Get the name of the binary (assuming it's the same as the directory name)
binary_name="pokedexgo"

# Check if 'bin' directory exists, create it if it doesn't
if [ ! -d "bin" ]; then
    mkdir bin
fi

# Move the binary to the 'bin' directory
mv "pokedexgo" bin/

echo "Build complete. Binary moved to bin/pokedexgo"

./bin/pokedexgo
