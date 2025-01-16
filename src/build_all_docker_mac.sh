#!/bin/bash
targets=(
    "x86_64-pc-windows-gnu"
    "x86_64-unknown-linux-gnu"
    "aarch64-unknown-linux-gnu"
    "x86_64-apple-darwin"
    "aarch64-apple-darwin"
)
output_dir="~/dev/rust/scrambler/target"
output_dir=$(eval echo "$output_dir") 
mkdir -p "$output_dir"
echo "Pulling the Rust Docker image..."
docker pull rust:latest || { echo "Failed to pull Docker image"; exit 1; }
build_target() {
    local target="$1"
    local output_name=""
    local binary_name="scrambler"
    echo "Building for target: $target"
    if [[ "$target" == "x86_64-unknown-linux-gnu" ]]; then
        output_name="scrambler-linux-x86_64"
        compiler_pkg="gcc-x86-64-linux-gnu"
        linker="x86_64-linux-gnu-gcc"
        strip_cmd="x86_64-linux-gnu-strip"
    elif [[ "$target" == "aarch64-unknown-linux-gnu" ]]; then
        output_name="scrambler-linux-arm64"
        compiler_pkg="gcc-aarch64-linux-gnu"
        linker="aarch64-linux-gnu-gcc"
        strip_cmd="aarch64-linux-gnu-strip"
    elif [[ "$target" == "x86_64-pc-windows-gnu" ]]; then
        output_name="scrambler-windows-x86_64.exe"
        compiler_pkg="gcc-mingw-w64-x86-64"
        linker="x86_64-w64-mingw32-gcc"
        strip_cmd="x86_64-w64-mingw32-strip"
        binary_name="scrambler.exe"
    elif [[ "$target" == "x86_64-apple-darwin" ]]; then
        output_name="scrambler-macos-intel"
        strip_cmd="strip"
    elif [[ "$target" == "aarch64-apple-darwin" ]]; then
        output_name="scrambler-macos-Mx"
        strip_cmd="strip"
    else
        echo "Unsupported target: $target"
        return
    fi
    if [[ "$target" == *"darwin"* ]]; then
        cargo build --release --target "$target" || { echo "Build failed for $target"; exit 1; }
        $strip_cmd "target/$target/release/$binary_name" || { echo "Failed to strip macOS binary for $target"; exit 1; }
    else
        docker run --rm -v $(pwd):/project -w /project rust bash -c "
            rustup target add $target &&
            apt-get update &&
            apt-get install -y $compiler_pkg &&
            CARGO_TARGET_$(echo "$target" | tr '[:lower:]' '[:upper:]' | tr '-' '_')_LINKER=$linker cargo build --target $target --release &&
            $strip_cmd target/$target/release/$binary_name
        " || { echo "Build failed for $target"; exit 1; }
    fi
    mv "target/$target/release/$binary_name" "$output_dir/$output_name" || { echo "Failed to move executable for $target"; exit 1; }
    echo "Executable for $target moved to $output_dir/$output_name"
    echo "Cleaning up build folder for $target"
    rm -rf "target/$target"
}
sign_executable() {
    local file="$1"
    local sig_file="$file.pgp"
    echo "Signing $file with PGP..."
    gpg --armor --output "$sig_file" --yes --detach-sign --local-user "ram@catsec.com" "$file" || { echo "Failed to sign $file"; exit 1; }
    echo "Signed $file and saved signature to $sig_file"
}
for target in "${targets[@]}"; do
    build_target "$target"
done
for executable in "$output_dir"/*; do
    if [[ -f "$executable" && ! "$executable" == *.pgp ]]; then
        sign_executable "$executable"
    fi
done
echo "All builds, stripping, and signing completed successfully."