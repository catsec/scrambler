# Wallet Word Scrambler

## Overview
Have you ever been worried that your wallet words backup be exposed and used to steal your wallet?

The **Scrambler** is a secure program designed to scramble or recover wallet backup words using a password. This ensures an added layer of security for your wallet's seed phrase. It is intended to be used on an air-gapped and freshly formatted machine to maximize security.

## Features

- Supports multiple wordlists (BIP-39 and SLIP-39).
- Uses Argon2 and SHA-3 for secure key derivation and hashing.
- Provides word suggestions for misspelled wallet words.
- Scrambles or recovers wallet words with user-provided passwords.
- Warns users about password strength and lack of recovery options.

## Security Warning

> **This program is meant to run on a freshly formatted and air-gapped machine.**
>
> - Do not run this program on a machine connected to any kind of network.
> - Securely wipe the machine after use to ensure no sensitive data remains.

## Usage

1. **Setup the Environment:**
   - Compile the source code or download pre-compiled single executable from my github
   - Check the PGP signature if you downloaded a pre-complied file 

2. **Run the Program:**
   ```bash
   go run main.go
   ```

3. **Follow the Prompts:**
   - Select the desired wordlist.
   - Choose to either scramble or recover wallet words.
   - Enter a password (securely).
   - Provide your wallet's seed phrase.

4. **Output:**
   - Scrambled or recovered wallet words will be displayed.

## Supported Wordlists

- **SLIP-39:** English (1024 words, used by Trezor).
- **BIP-39:**
  - English (2048 words)
  - Czech
  - Chinese Simplified
  - Chinese Traditional
  - French
  - Italian
  - Japanese
  - Korean
  - Spanish
  - Portuguese

## Technical Details

- **Key Derivation:** Uses Argon2 ID with high memory and computational cost to ensure strong resistance against brute-force attacks.
- **Hashing:** SHA-3 (512-bit) for repeated secure hashing.
- **Word Matching:** Utilizes the Levenshtein distance algorithm for word suggestions.
- **Bit Manipulation:** Encodes and manipulates wallet words as bit strings for scrambling.

## Password Policy

- A strong password should be at least 8 characters long and include:
  - Uppercase letters
  - Lowercase letters
  - Numbers
  - Special characters
- Weak passwords are discouraged but can be used if explicitly confirmed by the user.

## Example Usage

1. **Scrambling Wallet Words:**
   - Select a wordlist.
   - Enter a secure password.
   - Input your wallet's seed phrase.
   - The program will output a scrambled version of your seed phrase.

2. **Recovering Wallet Words:**
   - Select a wordlist.
   - Provide the password used for scrambling.
   - Enter the scrambled wallet words.
   - The program will output the original wallet words.

## Limitations

- Do not use this program on connected machine and wipe it clean after usage
- Password recovery is not possible; if the password is lost, the scrambled words cannot be recovered.


## Disclaimer

The **Scrambler** is provided "as is" without any warranty. The user assumes all responsibility for ensuring the security of their machine and environment when using this program. The authors are not liable for any loss of funds or data.

---

### Author

Written by **Ram Prass** (2024).

