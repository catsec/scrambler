// Module: utils
// Various utility functions used by the program
//

use crate::constants::*;
use levenshtein::levenshtein;
use std::net::TcpStream;

// Check if the user is connected to the internet
pub fn internetconnection() -> bool {
    TcpStream::connect("8.8.8.8:53").is_ok() // Google's public DNS
}

// Divide the key into chunks of the specified size
pub fn dividekey(data: [u8; 64], parts: usize, chunksize: usize) -> Vec<u16> {
    // Calculate the total number of bits required for the chunks
    let totalbits = parts * chunksize;

    // Calculate the number of bits available in the key
    let availablebits = data.len() * 8;
    if availablebits < totalbits {
        // show an error message if the key is too short
        panic!("Insufficient binary data for the requested chunks and size");
    }

    // Convert the key into a vector of u16 chunks (wordlist is max 11 bits)
    let mut chunks = Vec::with_capacity(parts);
    let mut chunk: u16 = 0;
    let mut bitcounter = 0;

    // Iterate over the key data and extract the chunks
    for &value in &data {
        let mut temp = value;
        for _ in 0..8 {
            // Extract the bits from the byte
            if bitcounter < chunksize {
                chunk |= ((temp & 1) as u16) << bitcounter;
                temp >>= 1;
                bitcounter += 1;
            }

            // we have a full chunk
            if bitcounter == chunksize {
                chunks.push(chunk);
                chunk = 0;
                bitcounter = 0;

                // we have all the chunks
                if chunks.len() == parts {
                    break;
                }
            }
        }

        // we have all the chunks
        if chunks.len() == parts {
            break;
        }
    }

    // show an error message if we could not generate all the chunks
    if chunks.len() != parts {
        panic!("Failed to generate all the chunks");
    }

    return chunks;
}

// suggest words based on the user input
pub fn find_suggestions(word: &str, wordlist: &[&str]) -> Vec<String> {
    let mut suggestions = Vec::new();

    // Words that start with the same first 4 letters
    if word.len() >= 4 {
        let prefix = &word[..4];
        suggestions.extend(
            wordlist
                .iter()
                .filter(|&&w| w.starts_with(prefix))
                .take(3) // Limit to 3 suggestions
                .cloned()
                .map(String::from),
        );
    }

    // Words with the smallest Levenshtein distance
    if suggestions.len() < 3 {
        let mut distances: Vec<(usize, &str)> = wordlist
            .iter()
            .map(|&w| (levenshtein(word, w), w))
            .filter(|&(dist, _)| dist <= 3) // Limit to a maximum distance of 3
            .collect();

        // Sort by distance, then alphabetically
        distances.sort_by(|a, b| a.0.cmp(&b.0).then_with(|| a.1.cmp(b.1)));

        // Add up to 3 additional suggestions
        suggestions.extend(
            distances
                .into_iter()
                .take(3 - suggestions.len()) // Fill up to 3 suggestions
                .map(|(_, w)| w.to_string()),
        );
    }

    suggestions
}

// Calculate the number of bits required to represent the word list
pub fn getwordlistbitsize(lang: usize) -> usize {
    // all words are non-empty, so we can count them directly
    let words = &WORDS[lang];

    // count only non-empty words
    let word_count = words.iter().filter(|&&word| !word.is_empty()).count();

    // calculate the number of bits required to represent the word list
    f64::from(word_count as u32).log2().ceil() as usize
}

// Scramble the wallet words using the secret key
pub fn scramblewords(words: Vec<usize>, secretkey: [u8; 64], lang: usize) -> Vec<usize> {
    // Get the number of bits required to represent the word list
    let wordlistbitsize = getwordlistbitsize(lang);

    // Divide the key into chunks of the required size
    let keychunks = dividekey(secretkey, words.len(), wordlistbitsize);

    // prepare a vector to store the new words
    let mut newwords = Vec::with_capacity(words.len());

    // Scramble each word using the key chunks
    for (i, &word) in words.iter().enumerate() {
        // XOR the word index with the key chunk
        let scrambled_word = word ^ keychunks[i] as usize;

        // get the word index
        let valid_word = scrambled_word % WORDS[lang].len();

        // store the new word
        newwords.push(valid_word);
    }
    newwords
}

// Print the wallet words to the user
pub fn printwords(words: &[usize], lang: usize, recover: bool) {
    // change the message based on the action
    if recover {
        println!("\nRecovered words:\n");
    } else {
        println!("\nNew words:\n");
    }

    // print the words with their indexes
    for (i, &word) in words.iter().enumerate() {
        // add a space before single digit indexes for better alignment
        let space = if i < 9 { " " } else { "" };
        println!("{}{}: {}", space, i + 1, WORDS[lang][word]);
    }
}
