# Wallet Word Scrambler

**Wallet Word Scrambler** is a command-line tool that securely scrambles or recovers wallet backup words using a user-provided password. It ensures additional protection for mnemonic backups, making it ideal for air-gapped systems.

## Features

- **Secure Mnemonic Scrambling**: Use a password to scramble your wallet's backup words.
- **Password Strength Validation**: Warns users about weak passwords but allows overrides.
- **Wallet Recovery**: Recovers scrambled wallet words using the same password.
- **Compatible Wordlists**: Supports BIP39-compatible wordlists containing **1024** or **2048 words**.
- **Air-Gapped Usage**: Designed for secure environments with no network access.

## Requirements

- **Go** (Golang) installed
- A BIP39-compatible wordlist file (e.g., `wordlist.txt`) with **1024** or **2048** words.
- An air-gapped machine for maximum security.

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/wallet-word-scrambler.git
   cd wallet-word-scrambler
   ```
2. **Build the Application**:
   ```bash
   go build -o wallet-scrambler main.go
   ```
3. Place your BIP39 wordlist file(s) in the same directory as the executable. The wordlist file must:
   - Have **1024** or **2048** lines.
   - Be a plain text `.txt` file.

## Usage

1. **Run the Application**:
   ```bash
   ./wallet-scrambler
   ```

2. **Follow the Instructions**:
   - Select a wordlist file.
   - Choose to either **scramble** or **recover** wallet words.
   - Provide and confirm a password.
   - Enter your wallet backup words (12-33 words).

3. **Output**:
   - If scrambling, the tool will provide a set of new scrambled words.
   - If recovering, the tool will output the original wallet words.

## Important Notes

- **Password Requirement**:
  - The password must be at least **8 characters** long and include a mix of:
    - Uppercase letters
    - Lowercase letters
    - Numbers
    - Special characters (`!@#$%^&*()-_=+[]{}` etc.)
  - Weak passwords can be accepted with user confirmation.

- **Air-Gapped Security**:
   - **Run this tool ONLY on a fresh air-gapped machine**.
   - **Do not connect the machine to any network** during use.
   - After use, **securely wipe the machine** to prevent data leaks.

- **No Password Recovery**:
   - If you lose your password, **there is NO way to recover your wallet words**.

## Example

1. Run the tool:
   ```bash
   ./wallet-scrambler
   ```
2. Select a wordlist:
   ```
   1. bip39-english.txt
   Select the wordlist file: 1
   ```
3. Follow prompts to scramble or recover wallet words.

## Dependencies

- `golang.org/x/crypto/argon2` - Password-based key derivation function.
- `golang.org/x/crypto/sha3` - Cryptographic hash functions.
- `github.com/agnivade/levenshtein` - Word suggestion for typo corrections.

## Security Disclaimer

This program is designed to improve the security of mnemonic backups. However:
- Use it **at your own risk**.
- Always test the tool in a secure, isolated environment before using it with real wallet backups.

## License

This project is licensed under the Apache2 License. See the `LICENSE` file for details.

## Contact

For questions or contributions, contact:
- Ram Prass: ram@catsec.com
- GitHub: [https://github.com/catsec](https://github.com/catsec)
