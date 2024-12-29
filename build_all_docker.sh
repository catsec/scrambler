#!/bin/bash

# List of targets to build for
targets=(
    "x86_64-pc-windows-gnu"
    "x86_64-unknown-linux-gnu"
    "aarch64-unknown-linux-gnu"
)

# Pull the Rust Docker image (ensure you have Docker installed and running)
echo "Pulling the Rust Docker image..."
docker pull rust:latest || { echo "Failed to pull Docker image"; exit 1; }

# Function to build for a specific target
build_target() {
    local target="$1"
    echo "Building for target: $target"

    # Determine the correct compiler package and linker
    if [[ "$target" == "x86_64-unknown-linux-gnu" ]]; then
        compiler_pkg="gcc-x86-64-linux-gnu"
        linker="x86_64-linux-gnu-gcc"
    elif [[ "$target" == "aarch64-unknown-linux-gnu" ]]; then
        compiler_pkg="gcc-aarch64-linux-gnu"
        linker="aarch64-linux-gnu-gcc"
    elif [[ "$target" == "x86_64-pc-windows-gnu" ]]; then
        compiler_pkg="gcc-mingw-w64-x86-64"
        linker="x86_64-w64-mingw32-gcc"
    else
        echo "Unsupported target: $target"
        return
    fi

    # Convert target to uppercase for the environment variable
    upper_target=$(echo "$target" | tr '[:lower:]' '[:upper:]' | tr '-' '_')

    docker run --rm -v $(pwd):/project -w /project rust bash -c "
        rustup target add $target &&
        apt-get update &&
        apt-get install -y $compiler_pkg &&
        CARGO_TARGET_${upper_target}_LINKER=$linker cargo build --target $target --release
    "

    if [ $? -eq 0 ]; then
        echo "Build succeeded for $target"
    else
        echo "Build failed for $target"
        exit 1
    fi
}

# Iterate over each target and build
for target in "${targets[@]}"; do
    build_target "$target"
done

echo "All builds completed successfully."
