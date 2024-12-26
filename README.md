# Catsec Wallet Word Scrambler

Ever wondered what would happen if your wallet backup got stolen? 

Catsec Wallet Word Scrambler is a Rust-based tool designed to securely scramble wallet backup words using a password of your choice. It provides a way to enhance the security of your wallet by obfuscating your recovery words. The program supports scrambling and recovery workflows.

## Features
- **Scramble Wallet Words**: Securely scramble your wallet recovery words using a user-provided password.
- **Recover Wallet Words**: Recover scrambled wallet words using the correct password.
- **Wordlist Support**: Includes support for multiple languages and wordlists.
- **Password Security**: Utilizes strong cryptographic algorithms like Argon2 and SHA3 for password hashing and scrambling.
- **Platform Independence**: Designed to run on air-gapped machines without requiring a network connection.

## Security Warning
This program is intended to run on:
- Freshly formatted and air-gapped machines.
- Machines with no active network connections.

Although nothing is saved locally, it is strongly recommended to securely wipe the machine after use.

## Requirements
- **Rust**: The program is written in Rust and requires the Rust toolchain to build.
- **Dependencies**: The following Rust crates are used:
  - `argon2`: For password hashing.
  - `sha3`: For cryptographic hashing.
  - `levenshtein`: For word suggestion based on similarity.

## Installation

### Prerequisites
1. Install Rust: [Rust Installation Guide](https://www.rust-lang.org/tools/install)
2. Clone the repository:
   ```bash
   git clone https://github.com/catsec/scrambler.git
   cd scrambler
   ```

### Build
To build the project, run:
```bash
cargo build --release
```
The compiled binary will be available in the `target/release/` directory.

## Usage

### Running the Program
To run the program, use:
```bash
cargo run --release
```

### Workflow
1. **Scramble Wallet Words**:
   - Select the "Scramble a new wallet" option.
   - Choose the desired language and wordlist.
   - Provide a secure password.
   - Enter your wallet words.
   - View or save the scrambled words.

2. **Recover Wallet Words**:
   - Select the "Recover an existing wallet" option.
   - Provide the scrambled words and password.
   - The program will unscramble and display the original words.

### File Support
You can optionally save or load wallet words from `.txt` files in the current directory.

### Example
```bash
Welcome to Catsec's wallet word scrambler

What would you like to do?
1. Scramble a new wallet
2. Recover an existing wallet

Enter your choice: 1
```

## Development
### File Structure
- `src/`
  - `constants.rs`: Contains constant definitions such as wordlists and cryptographic parameters.
  - `input.rs`: Handles user input and validation.
  - `crypto.rs`: Implements cryptographic functions like hashing and key derivation.
  - `utils.rs`: Contains utility functions for word scrambling and recovery.
- `Cargo.toml`: Rust package configuration.

### Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

## License
This project is licensed under the Apache 2.0 License. See the `LICENSE` file for details.

## Author
**Ram Prass** - Catsec

---

### Disclaimer
This tool is provided "as is" without warranty of any kind. Use at your own risk. Ensure proper backups before scrambling or recovering wallet words.
