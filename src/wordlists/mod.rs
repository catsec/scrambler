// Wordlist modules for different languages
pub mod slip39;
pub mod english;
pub mod czech;
pub mod french;
pub mod italian;
pub mod portuguese;
pub mod spanish;
pub mod japanese;
pub mod korean;
pub mod chinese_simplified;
pub mod chinese_traditional;

use slip39::SLIP39_WORDS;
use english::ENGLISH_WORDS;
use czech::CZECH_WORDS;
use french::FRENCH_WORDS;
use italian::ITALIAN_WORDS;
use portuguese::PORTUGUESE_WORDS;
use spanish::SPANISH_WORDS;
use japanese::JAPANESE_WORDS;
use korean::KOREAN_WORDS;
use chinese_simplified::CHINESE_SIMPLIFIED_WORDS;
use chinese_traditional::CHINESE_TRADITIONAL_WORDS;

// Re-export the words array for compatibility with existing code
pub const WORDS: [[&str; 2048]; 11] = [
    SLIP39_WORDS,
    ENGLISH_WORDS,
    CZECH_WORDS,
    FRENCH_WORDS,
    ITALIAN_WORDS,
    PORTUGUESE_WORDS,
    SPANISH_WORDS,
    JAPANESE_WORDS,
    KOREAN_WORDS,
    CHINESE_SIMPLIFIED_WORDS,
    CHINESE_TRADITIONAL_WORDS,
];