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


## Installation

If you don't want to compile, help yourself to the binaries: [here](https://github.com/catsec/scrambler/releases/latest).

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

## Development
### File Structure
- `src/`
  - `constants.rs`: Contains constant definitions such as wordlists and cryptographic parameters.
  - `input.rs`: Handles user input and validation.
  - `crypto.rs`: Implements cryptographic functions like hashing and key derivation.
  - `utils.rs`: Contains utility functions for word scrambling and recovery.
- `Cargo.toml`: Rust package configuration.


## How it works

the program take two inputs:
- password
- list of words

Then the follow process accures:
- Derive a 512bit key from the password and deviding it to 10 or 11 bits chunks (according to the wordlist size).
- Words are then converted to their index according to the wordlist resulting in a 10 or 11 bit value for each word.
- This value is then XORed with the key chunk until all words are XORed.
- Calculate a new word according to the XORed values index.
- the result is a list of new words that if will be xored again with a key derived from the same password will return the original value.

This process will returen a valid list word for any password, eliminating known plaintext attacks (though SLIP39 produces preditable first words that might be used to eliminate some results)

## License
This project is licensed under the Apache 2.0 License. See the `LICENSE` file for details.

## Author
**Ram Prass** - Catsec

## Q&A

### Why do I even need it?

Current wallet backups are not encrypted. This means that if someone gains access to your backup, you're basically F**KED. No safe is impenetrable, and someone will always find a way to access it.
Using a password to protect your wallet gives you some assurance that, even if someone gets hold of your wallet’s recovery words, they won’t be able to use them without knowing your password.

## Why do you create a word list?

I decided to use a word list to maintain compatibility with hardware tools designed to safeguard wallets.
This way, you can etch the scrambled words onto those devices, ensuring they are not the original sequence but still recoverable.

## Why does it take so long (up to 2 minutes) to derive the key?

TL;DR: I don’t trust you.

The short answer is that while I encourage users to create strong passwords, most don’t.
From easily guessable personal details to slight variations of previously used passwords, users tend to choose poor passwords.

By ensuring every password produces a legitimate output and employing a time-intensive key derivation process, the system offers better protection against hackers.

Additionally, while I use state-of-the-art algorithms (SHA-3, Argon2id), they may become vulnerable over time. Wallet backups are meant to last for years, so extra caution is necessary.

## Can I save my generated scrambled word list on the cloud?

I wouldn’t advise it. Always use multiple layers of protection:

Store it on a USB stick, place it in a safe, or both. After all, it’s your money at stake.

## How can I be sure you’re not stealing my wallet?

You can’t! Don’t trust me, and don’t trust anyone else either.

Review the code if you’re able to, and always use an air-gapped machine. Wipe it clean after use!

## All this security advice seems excessive: air-gapped machines, wiping everything clean...

No, it’s not excessive.

When it comes to your money, no precaution is too great.


---
### Disclaimer
This tool is provided "as is" without warranty of any kind. Use at your own risk. Ensure proper backups before scrambling or recovering wallet words.
