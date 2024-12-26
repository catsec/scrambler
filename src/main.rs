//
// *** Catsec wallet word scrambler ***
//
// This program will help you scramble your wallet backup words using a password of your choice
// It is meant to run on a fresh formatted and air-gapped machine
// It is not safe to run it on a machine connected to any kind of network
// Though nothing is saved - secure wipe your machine immediately after use
// The program is written in Rust and uses the following crates:
// - argon2: for password hashing
// - sha3: for hashing
//
// This program is released under apache 2.0 license - copyright (2024) Ram Prass - Catsec
// 
mod constants;
mod input;
mod crypto;
mod utils;

use input::*;
use constants::*;
use crypto::*;
use utils::*;

fn main() {
    // Main function to scramble wallet words

    println!("\nWelcome to Catsec's wallet word scrambler");
    
    println!("This program will help you scramble your wallet backup words");
    println!("using a password of your choice\n");
    
    println!("Warning:\n");
    println!("This program is meant to run on a fresh formatted and air-gapped machine");
    println!("It is not safe to run it on a machine connected to any kind of network");
    println!("Though nothing is saved - secure wipe your machine immediately after use");
    
    // Ask the user if they want to scramble a new wallet or recover an existing one
    let recover = choose("What would you like to do?", &["Scramble a new wallet", "Recover an existing wallet"]) == 1;
    
    // init the wallet size, language, and words vector based on the action
    let (mut walletsize,mut lang, mut words) = if recover {
        // ask if to recover the wallet from a file
        recoverfromfile()
    } else {
        // scramble a new wallet
        (0, 0, vec![])
    };

    // get the language if not recovering from a file using walletsize=0 to see if it was recovered from file, lang 0 is valid)
    if walletsize == 0 {
        
        // get the language from the user
        lang = choose("What wordlist would you like to use?", &LANG);
    }
    
    // get the password from the user
    let password = getpassword(recover);
    
    // derive the secret key from the password
    let secretkey = derivekey(password);
    if walletsize == 0 {

        // get the wallet size if not recovering from a file
        walletsize = getwalletsize();
    }
    if words.is_empty() {

        // get the wallet words if not recovering from a file
        words = getwords(walletsize, lang);
    }

    // scramble the wallet words using the secret key
    let newwords: Vec<usize> = scramblewords(words, secretkey, lang);
    
    // print the new words to the user
    println!("");
    printwords(&newwords, lang, recover);
    if !recover {

        // if not recovering from a file, ask the user if they want to save the wallet
        savewallet(&newwords, lang);
    }
}