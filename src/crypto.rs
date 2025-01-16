// Module: crypto
// Chaning the hashing algorithm parameters or the number of iterations will make the key incompatible with the default one
// and you will not be able to recover your wallet
//
use crate::constants::*;

use argon2::{Argon2, Params};
use sha3::{Digest, Sha3_512};
use std::io::{self, Write};
use std::process;
use std::time::Instant;

// Hash the data using SHA3-512 for the specified number of iterations
pub fn sha3(data: &[u8], iterations: u32) -> Vec<u8> {
    let mut hash = data.to_vec();
    for _ in 0..iterations {
        let mut hasher = Sha3_512::new();
        hasher.update(&hash);
        hash = hasher.finalize().to_vec();
    }
    hash
}

// Derive a secret key from the password using Argon2
pub fn derive_key(password: Vec<u8>) -> [u8; 64] {
    println!("Deriving secret key, this WILL take a while (have some tea and relax)\n");

    // Create Argon2 parameters
    let params = Params::new(A_MEMORY, A_TIME, A_PARALLELISM, Some(64))
        .expect("Failed to create Argon2 parameters");

    // Create an Argon2 instance with Argon2id and version 0x13 (the latest version at the time of writing)
    let argon2 = Argon2::new(argon2::Algorithm::Argon2id, argon2::Version::V0x13, params);

    // Create a 64-byte buffer to store the secret key
    let mut secret_key = [0u8; 64];

    // Track the start time
    let start_time = Instant::now();
    
    print!("[{}] 0% (Time left: calculating)", " ".repeat(BAR_SIZE));
    io::stdout().flush().expect("Failed to flush stdout");
    
    // Iterate for the specified number of iterations
    for i in 1..=ITERATIONS {
        // Create a new salt for each iteration by hashing the password and the iteration number
        let counter: u32 = i as u32 + 580;
        let salt = sha3(&password, counter);

        // Hash the password into the secret key using Argon2
        if let Err(e) = argon2.hash_password_into(&password, &salt, &mut secret_key) {
            eprintln!("Error hashing password: {}", e);
            process::exit(1);
        }

        // Calculate progress and estimated remaining time
        let elapsed_time = start_time.elapsed();
        let avg_time_per_step = elapsed_time / i as u32;
        let remaining_steps = ITERATIONS - i;
        let estimated_remaining_time = format!("{} seconds", (avg_time_per_step * remaining_steps as u32).as_secs());

        // Generate progress bar
        let progress_dots = "=".repeat((i * BAR_SIZE / ITERATIONS) as usize);
        let remaining_dots = " ".repeat((BAR_SIZE - progress_dots.len()) as usize);
        
        // Update the progress line
        print!(
            "\r[{}{}] {}% (Time left: {})",
            progress_dots,
            remaining_dots,
            (i * 100) / ITERATIONS,
            estimated_remaining_time
        );
        io::stdout().flush().expect("Failed to flush stdout");
    }

    println!("\n\nKey derivation completed successfully.");

    // Return the secret key
    secret_key
}
