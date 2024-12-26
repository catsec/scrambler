//
// Chaning the hashing algorithm parameters or the number of iterations will make the key incompatible with the default one
// and you will not be able to recover your wallet
//
use crate::constants::*;

use argon2::{Argon2, Params};
use sha3::{Digest, Sha3_512};
use std::io::{self, Write};
use std::process;

pub fn sha3(data: &[u8], iterations: u32) -> Vec<u8> {
    
    // Hash the data using SHA3-512 for the specified number of iterations
    let mut hash = data.to_vec();
    for _ in 0..iterations {
        let mut hasher = Sha3_512::new();
        hasher.update(&hash);
        hash = hasher.finalize().to_vec();
    }
    return hash
}

pub fn derivekey(password: Vec<u8>) -> [u8; 64] {
    // Derive a secret key from the password using Argon2
    
    println!("\nCalculating derived secret key, this WILL take a while\n");
    
    // create Argon2 parameters
    // if you change those your key will not be compatible with the default one and you will not be able to recover your wallet
    let params = Params::new(
        A_MEMORY,         
        A_TIME,               
        A_PARALLELISM,               
        Some(64),    
    ).expect("Failed to create Argon2 parameters");
    
    // create an Argon2 instance with Argon2id and version 0x13 (the latest version at the time of writing)
    let argon2 = Argon2::new(argon2::Algorithm::Argon2id, argon2::Version::V0x13, params);
    
    // create a 64-byte buffer to store the secret key
    let mut secretkey = [0u8; 64];
    
    // repeat for a number of iterations (10 for now)

    for i in 1..=ITERATIONS {
        
        // create a new salt for each iteration by hashing the password and the iteration number
        
        let counter: u32 = i as u32+580;
        let salt = sha3(&password, counter);
        
        // hash the password into the secret key using Argon2
        if let Err(e) = argon2.hash_password_into(&password, &salt, &mut secretkey) {
            
            // oops...
            eprintln!("Error hashing password: {}", e);
            process::exit(1);
        }
        
        // show the progress to the user
        print!("..{}%", i * (100 / ITERATIONS));
        io::stdout().flush().expect("Failed to flush stdout");
        }
    
    // retuen the secret key
    return secretkey
}
