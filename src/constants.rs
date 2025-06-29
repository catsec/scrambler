// module: constants
// constants used by the program
//

// Import wordlists from separate modules
pub use crate::wordlists::WORDS;

// do not change this values:
pub const A_TIME: u32 = 5;
pub const A_MEMORY: u32 = 2 * 1024 * 1024;
pub const A_PARALLELISM: u32 = 4;
pub const ITERATIONS: usize = 10;
pub const BAR_SIZE: usize = 40;

// you may change this on your own risk
pub const MAX_WORDS: usize = 33;
pub const UPPER: &str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
pub const LOWER: &str = "abcdefghijklmnopqrstuvwxyz";
pub const NUMBERS: &str = "0123456789";
pub const SPECIAL: &str = "!@#$%^&*()-_=+[]{}|;:'\",.<>?/";

// if wordlists are added you can add them here (but don't overide)
pub const LANG: [&str; 11] = [
    "SLIP39 (English, 1024 words, used by Trezor)",
    "English (BIP 39, 2048 words)",
    "Czech (BIP 39, 2048 words)",
    "French (BIP 39, 2048 words)",
    "Italian (BIP 39, 2048 words)",
    "Portuguese (BIP 39, 2048 words)",
    "Spanish (BIP 39, 2048 words)",
    "Japanese (BIP 39, 2048 words)",
    "korean (BIP 39, 2048 words)",
    "Chinese simplified (BIP 39, 2048 symbols)",
    "chinese traditional (BIP 39, 2048 symblos)",
];