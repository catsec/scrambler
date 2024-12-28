// Module: input
// Module for getting input from the user
//
use crate::constants::*;
use crate::utils::*;
use std::io::{self, Write};
use std::process;

// Function to choose an action from a list of choices
pub fn choose(action: &str, choices: &[&str]) -> usize {
    // Display the action and choices to the user

    assert!(!choices.is_empty(), "Choices cannot be empty");

    // show the paction to the user
    println!("\n{}\n", action);

    // customize the prompt based on the number of choices ( 1 or 2 vs 1 to n)
    let prompt = if choices.len() > 2 {
        format!("Enter a number between 1 and {}: ", choices.len())
    } else {
        "Enter 1 or 2: ".to_string()
    };

    // loop until the user enters a valid choice
    loop {
        for (index, choice) in choices.iter().enumerate() {
            // add a space before single digit choices for better alignment
            let space = if index < 9 { " " } else { "" };
            println!("{}{}. {}", space, index + 1, choice);
        }

        // get the user input allowing only numbers
        let input = getinput(&prompt, NUMBERS);

        // parse the input as a number and check if it is a valid choice
        match input.trim().parse::<usize>() {
            Ok(num) if num >= 1 && num <= choices.len() => {
                // return the choice as a 0-based index
                return num - 1;
            }

            // show an error message if the choice is invalid
            _ => println!("\nInvalid choice. Please try again.\n"),
        }
    }
}

// Function to get input from the user and validate it
pub fn getinput(prompt: &str, allowed: &str) -> String {
    // Get input from the user and validate it
    let mut input = String::new();
    loop {
        // Show the prompt and get the input
        print!("\n{}", prompt);
        io::stdout().flush().unwrap();

        // Clear the input buffer
        input.clear();
        io::stdin().read_line(&mut input).unwrap();
        let input = input.trim();

        // Check if the input is valid
        if input.chars().all(|c| allowed.contains(c)) {
            return input.to_string();
        } else {
            println!("\nInvalid input. Please enter a valid input.");
        }
    }
}

// Get a password from the user and validate it
pub fn getpassword(recover: bool) -> Vec<u8> {
    // allowed characters for the password: upper case, lower case, numbers, special characters
    let allowed = format!("{}{}{}{}", UPPER, LOWER, NUMBERS, SPECIAL);
    if !recover {
        // if the user is not recovering a wallet, show a warning about the password strength
        println!(
            "\nIt's extremely important to choose a strong password\n\
        Nothing would help you if your password is cracked or guessed.\n\
        12 chars long and a mix of upper, lower, numbers & special chars is recommended.\n"
        );
    }

    // loop until the user enters a valid password
    loop {
        // get the password from the user two times
        let password = getinput("Enter password: ", &allowed);
        let password2 = getinput("Enter password again: ", &allowed);

        // check if the passwords match
        if password != password2 {
            println!("\nPasswords do not match");
            continue;
        }

        if recover {
            // don't check the password strength if the user is recovering a wallet
            return password.into_bytes();
        }

        // check if the password has at least one upper case, one lower case, one number, one special character
        let has_upper = password.chars().any(|c| c.is_ascii_uppercase());
        let has_lower = password.chars().any(|c| c.is_ascii_lowercase());
        let has_number = password.chars().any(|c| c.is_ascii_digit());
        let has_special = password.chars().any(|c| SPECIAL.contains(c));

        if !has_upper || !has_lower || !has_number || !has_special || password.len() < 12 {
            // show a warning if the password is weak

            // reuqire the user to confirm if they want to continue with a weak password
            let agree=getinput("\nPassword is weak. Are you sure you want to continue? (type \"YES\" to continue): ", &UPPER);
            if agree == "YES" {
                // User confirmed to continue with a weak password
                println!("\nRemember your password, it CANNOT be recovered.\n");
                return password.into_bytes();
            }
        } else {
            // password is strong
            println!("\nRemember your password, it CANNOT be recovered.\n");
            return password.into_bytes();
        }
    }
}

// promot the user to get the number of words in the wallet
pub fn getwalletsize() -> usize {
    loop {
        // get the input from the user allowiung only numbers
        let input = getinput(
            "\nEnter the number of words in your wallet (12-33): ",
            NUMBERS,
        );
        match input.trim().parse::<usize>() {
            Ok(w) if w >= 12 && w <= 33 => return w,
            _ => println!("\nInvalid wallet size. Enter a number between 12 and 33."),
        }
    }
}

// get the words from the user and validate them
pub fn getwords(walletsize: usize, lang: usize) -> Vec<usize> {
    // Ensure wallet size does not exceed the maximum allowed
    if walletsize > MAX_WORDS {
        panic!("Wallet size cannot exceed {}", MAX_WORDS);
    }

    // Create a vector to store the indexes of the words
    let mut indexes: Vec<usize> = vec![0; walletsize];

    for i in 0..walletsize {
        loop {
            // Prompt the user to enter the word
            print!("Enter word number {}: ", i + 1);
            io::stdout().flush().expect("Failed to flush stdout");
            let mut input = String::new();
            io::stdin()
                .read_line(&mut input)
                .expect("Failed to read input");
            let word = input.trim();

            // Check if the word exists in the word list
            if let Some(index) = WORDS[lang].iter().position(|&w| w == word) {
                indexes[i] = index;
                break;
            } else {
                // Find suggestions based on the entered word
                let suggestions = find_suggestions(word, &WORDS[lang]);

                // Show an error message and suggestions
                println!("\nInvalid word. Please enter a valid word from the word list.");
                if !suggestions.is_empty() {
                    println!("\nDid you mean one of these?");
                    for suggestion in suggestions {
                        println!(" -> {}", suggestion);
                    }
                }
            }
        }
    }

    // Return the indexes of the words as a vector
    return indexes;
}

// Save the wallet words to a file
pub fn savewallet(words: &[usize], lang: usize) {
    // Ask the user if they want to save the wallet
    let save = choose(
        "Would you like to save your scrambled wallet words?",
        &["Yes", "No"],
    );
    if save == 1 {
        // User does not want to save the wallet
        return;
    }

    // Get the filename for the wallet
    println!("file will be saved as .txt in the current directory");
    let allowed = format!("{}{}{}", UPPER, LOWER, NUMBERS);
    let filename = getinput(
        "Enter a filename for your wallet (no extension): ",
        &allowed,
    );
    let filename = format!("{}.txt", filename);

    // try to create the file and write the words to it
    let mut file = std::fs::File::create(filename.clone()).expect("Failed to create file");
    for &word in words {
        // write the word to the file
        writeln!(file, "{}", WORDS[lang][word]).expect("Failed to write to file");
    }

    println!("\nWallet saved to {}", filename);
}

// Recover the wallet words from a file
pub fn recoverfromfile() -> (usize, usize, Vec<usize>) {
    // Ask the user if they want to recover from a file
    let choice = choose("Do you want to recover from a file?", &["Yes", "No"]);
    if choice == 1 {
        // User does not want to recover from a file
        return (0, 0, vec![]);
    }
    // assenble the allowed characters for the filename
    let allowed = format!("{}{}{}", UPPER, LOWER, NUMBERS);
    println!("\nFile should be a .txt file in the current directory.");
    let filename = getinput(
        "Enter the filename of your wallet (no extension): ",
        &allowed,
    );
    let filename = format!("{}.txt", filename);

    // try to read the file and recover the wallet words
    let file = std::fs::read_to_string(&filename).expect("Failed to read file");
    let readwords: Vec<&str> = file.lines().collect();
    let walletsize = readwords.len();

    // check if the wallet size is valid
    if walletsize < 12 || walletsize > 33 {
        panic!("Wallet size must be between 12 and 33 words.");
    }

    // try to recover the wallet words for each language
    for (lang_index, words_in_language) in WORDS.iter().enumerate() {
        // create a vector to store the indexes of the words
        let mut indices = Vec::with_capacity(walletsize);

        // check if all the words in the file are in the current language
        let mut all_words_match = true;

        for &word in &readwords {
            // check if the word is in the word list
            if let Some(index) = words_in_language.iter().position(|&w| w == word) {
                // store the index of the word
                indices.push(index);
            } else {
                // the word is not in the word list
                all_words_match = false;
                break;
            }
        }

        // check if all the words in the file are in the current language
        if all_words_match {
            // all words are in the current language
            println!("\nWallet recovered from file: {}\n", filename);
            println!("Language: {}\n", LANG[lang_index]);

            // print the recovered words
            println!("here are the words found in the file (before unscambling)");
            printwords(&indices, lang_index, true);
            println!("\nTo unscramble the words, enter the password");

            // return the wallet size, language index, and word indices
            return (walletsize, lang_index, indices);
        }
    }
    // the words in the file are not in any of the supported languages
    println!("The wallet file contains words not found in any supported language.");
    process::exit(1);
}

// Warn the user if they are connected to the internet and ask if they want to continue
pub fn warnuser() {
    println!("\n******************************************************************");
    println!("*             WARNING: YOU ARE CONNECTED TO THE INTERNET         *");
    println!("*                    THIS A REALLY BAD IDEA                      *");
    println!("*                                                                *");
    println!("* If there is by chance a maleware on your computer your wallet  *");
    println!("* might be exposed and lost. unless you are just testing this    *");
    println!("* utility, please disconnect, and wipe the computer after usage  *");
    println!("******************************************************************\n");

    print!("Are you want to continue? (type \"YES\" to continue):");
    std::io::stdout().flush().unwrap();
    let mut input = String::new();
    std::io::stdin().read_line(&mut input).unwrap();
    if input.trim() != "YES" {
        println!("Exiting...");
        process::exit(0);
    }
}
