#!/bin/bash

# List of targets to build for
targets=(
    "x86_64-pc-windows-gnu"
    "x86_64-unknown-linux-gnu"
    "aarch64-unknown-linux-gnu"
    "x86_64-apple-darwin"
    "aarch64-apple-darwin"
)

# Output directory
output_dir="~/dev/rust/scrambler/target"
output_dir=$(eval echo "$output_dir") # Expand ~ to absolute path

# Ensure output directory exists
mkdir -p "$output_dir"

# Pull the Rust Docker image (ensure you have Docker installed and running)
echo "Pulling the Rust Docker image..."
docker pull rust:latest || { echo "Failed to pull Docker image"; exit 1; }

# Function to build for a specific target
build_target() {
    local target="$1"
    local output_name=""

    echo "Building for target: $target"

    # Determine the correct compiler package and linker for non-macOS builds
    if [[ "$target" == "x86_64-unknown-linux-gnu" ]]; then
        output_name="scrambler-linux-x86_64"
        compiler_pkg="gcc-x86-64-linux-gnu"
        linker="x86_64-linux-gnu-gcc"
    elif [[ "$target" == "aarch64-unknown-linux-gnu" ]]; then
        output_name="scrambler-linux-arm64"
        compiler_pkg="gcc-aarch64-linux-gnu"
        linker="aarch64-linux-gnu-gcc"
    elif [[ "$target" == "x86_64-pc-windows-gnu" ]]; then
        output_name="scrambler-windows-x86_64"
        compiler_pkg="gcc-mingw-w64-x86-64"
        linker="x86_64-w64-mingw32-gcc"
    elif [[ "$target" == "x86_64-apple-darwin" ]]; then
        output_name="scrambler-macos-intel"
    elif [[ "$target" == "aarch64-apple-darwin" ]]; then
        output_name="scrambler-macos-Mx"
    else
        echo "Unsupported target: $target"
        return
    fi

    # Build command
    if [[ "$target" == *"darwin"* ]]; then
        # Native build for macOS
        cargo build --release --target "$target" || { echo "Build failed for $target"; exit 1; }
    else
        # Cross-compilation using Docker
        docker run --rm -v $(pwd):/project -w /project rust bash -c "
            rustup target add $target &&
            apt-get update &&
            apt-get install -y $compiler_pkg &&
            CARGO_TARGET_$(echo "$target" | tr '[:lower:]' '[:upper:]' | tr '-' '_')_LINKER=$linker cargo build --target $target --release
        " || { echo "Build failed for $target"; exit 1; }
    fi

    # Move the executable to the target directory
    if [[ "$target" == *"windows"* ]]; then
        mv "target/$target/release/scrambler.exe" "$output_dir/$output_name" || { echo "Failed to move executable for $target"; exit 1; }
    else
        mv "target/$target/release/scrambler" "$output_dir/$output_name" || { echo "Failed to move executable for $target"; exit 1; }
    fi

    echo "Executable for $target moved to $output_dir/$output_name"

    # Clean up build artifacts
    echo "Cleaning up build folder for $target"
    rm -rf "target/$target"
}

# Function to sign an executable
sign_executable() {
    local file="$1"
    local sig_file="$file.pgp"

    echo "Signing $file with PGP..."
    gpg --armor --output "$sig_file" --detach-sign "$file" || { echo "Failed to sign $file"; exit 1; }
    echo "Signed $file and saved signature to $sig_file"
}

# Iterate over each target and build
for target in "${targets[@]}"; do
    build_target "$target"

done

# Sign all built executables
for executable in "$output_dir"/*; do
    if [[ -f "$executable" && ! "$executable" == *.pgp ]]; then
        sign_executable "$executable"
    fi
done

echo "All builds and signing completed successfully."
