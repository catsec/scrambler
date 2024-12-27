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

nope - just run the binary

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


## License
This project is licensed under the Apache 2.0 License. See the `LICENSE` file for details.

## Author
**Ram Prass** - Catsec

## Q&A

### Why do I even need it?

Current wallet backups are not encrypted, that means that if someone gets their hand on it, you are basicly F**ked
no safe is strong enough, and someone always finds the key.
Using a password to protect it gives you some assurance that even if someone got your wallet words, they can't
use them without knowing your password

### Why do you create a word list

I decided to use a word list to be "compatible" with hardware instruments that are being sold to safe keep your
wallet. in that way you can "etch" the words on those devices only its a scarmbled one - and not the "original"


### Why does it take so long (might be even 2 minutes) to derive the key?

TLDR; I don't trust you 

The simple answer is that even though I encourage users to use a real strong password, most users don't do it.
from easily guessable personal details to previusly used passwords (variations) - users tend to choose a bad password
the fact the every password will give a legitmate output and the usage of long key derivation gives you a better chance 
to withstand a hacker attack

Also while I use top of the class algorithms (sha-3, Argon2id) they might becode vulnerable as time passess,
Wallet backup should be there for years...

### Can I save my generated scrambled word list on the cloud?

I wouldn't advise it - always use several defense mechanisms, store it on a USB stick. put it in a safe. after all - its your money...

### How can I be sure you are not stealing my wallet?

You can't! don't trust me, don't trust anyone!
Check the code if you can, and always use air-gaped machine and wipe it clean after!

### All those security advices, are a bit exsesive aren't they? air-gaped machine... wiping clear... 

Nope they are not. 


---
### Disclaimer
This tool is provided "as is" without warranty of any kind. Use at your own risk. Ensure proper backups before scrambling or recovering wallet words.
