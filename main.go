package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/catsec/scrambler/wordlists"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/sha3"

	"github.com/agnivade/levenshtein"
)

var words []string

func selectWordList() {

	printStyled("\n{cyan}{bold}Select a word list:\n\n")
	printStyled(" 1. {bold}English (slip39, 1024 words, used by Trezor)\n")
	printStyled(" 2. {bold}English (bip39, 2048 words)\n")
	printStyled(" 3. {bold}Czech (bip39, 2048 words)\n")
	printStyled(" 4. {bold}Chinese simplified (bip39, 2048 words)\n")
	printStyled(" 5. {bold}Chinese traditional (bip39, 2048 words)\n")
	printStyled(" 6. {bold}French (bip39, 2048 words)\n")
	printStyled(" 7. {bold}Italian (bip39, 2048 words)\n")
	printStyled(" 8. {bold}Japanese (bip39, 2048 words)\n")
	printStyled(" 9. {bold}Korean (bip39, 2048 words)\n")
	printStyled("10. {bold}Spanish (bip39, 2048 words)\n")
	printStyled("11. {bold}Portuguese (bip39, 2048 words)\n")
	reader := bufio.NewReader(os.Stdin)
	var selection int
	for {
		printStyled("\n{cyan}Enter the number of the word list to use: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		selection, err = strconv.Atoi(input)
		if err == nil && selection >= 1 && selection <= 3 {
			break
		}
		printStyled("{red}Invalid selection. Please try again.\n")
	}
	switch selection {
	case 1:
		words = wordlists.Slip39
	case 2:
		words = wordlists.English
	case 3:
		words = wordlists.Czech
	case 4:
		words = wordlists.Chinesesimplified
	case 5:
		words = wordlists.Chinesetraditional
	case 6:
		words = wordlists.French
	case 7:
		words = wordlists.Italian
	case 8:
		words = wordlists.Japanese
	case 9:
		words = wordlists.Korean
	case 10:
		words = wordlists.Spanish
	case 11:
		words = wordlists.Portuguese
	}
	printStyled("\n{green}Word list successfully loaded!\n\n")
}

func findWord(word string) int {
	for i, w := range words {
		if w == word {
			return i
		}
	}
	return -1
}

func suggestWords(word string) []string {
	type suggestion struct {
		word  string
		score int
	}
	var suggestions []suggestion

	for _, w := range words {
		distance := levenshtein.ComputeDistance(word, w)
		suggestions = append(suggestions, suggestion{word: w, score: distance})
	}

	sort.Slice(suggestions, func(i, j int) bool {
		return suggestions[i].score < suggestions[j].score
	})

	prefix := ""
	if len(word) >= 4 {
		prefix = word[:4]
	}

	prioritySuggestions := []string{}
	regularSuggestions := []string{}

	for _, s := range suggestions {
		if len(prioritySuggestions) < 3 && len(prefix) >= 4 && len(s.word) >= 4 && s.word[:4] == prefix {
			prioritySuggestions = append(prioritySuggestions, s.word)
		} else if len(regularSuggestions) < 3 && len(prioritySuggestions) < 3 {
			regularSuggestions = append(regularSuggestions, s.word)
		}
	}

	return append(prioritySuggestions, regularSuggestions...)
}

func isWeakPassword(password string) bool {
	if len(password) < 8 {
		return true
	}
	hasLower := strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz")
	hasUpper := strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	hasNumber := strings.ContainsAny(password, "0123456789")
	hasSpecial := strings.ContainsAny(password, "!@#$%^&*()-_=+[]{}|;:',.<>?/")

	return !(hasLower && hasUpper && hasNumber && hasSpecial)
}

func hashRepeatedly(data []byte, iterations int) []byte {
	hash := data
	for i := 0; i < iterations; i++ {
		digest := sha3.Sum512(hash)
		hash = digest[:]
	}
	return hash
}

func bytesToBitString(data []byte) string {
	var bitString strings.Builder
	for _, b := range data {
		for i := 7; i >= 0; i-- {
			if (b & (1 << i)) != 0 {
				bitString.WriteByte('1')
			} else {
				bitString.WriteByte('0')
			}
		}
	}
	return bitString.String()
}

func splitString(input string, length int) []string {
	if length <= 0 {
		return []string{}
	}
	var result []string
	for i := 0; i < len(input); i += length {
		end := i + length
		if end > len(input) {
			end = len(input)
		}
		result = append(result, input[i:end])
	}
	return result
}

func xorBitStrings(bits1, bits2 string) string {
	if len(bits1) < len(bits2) {
		bits1 = strings.Repeat("0", len(bits2)-len(bits1)) + bits1
	} else if len(bits2) < len(bits1) {
		bits2 = strings.Repeat("0", len(bits1)-len(bits2)) + bits2
	}
	var result strings.Builder
	for i := 0; i < len(bits1); i++ {
		if bits1[i] == bits2[i] {
			result.WriteByte('0')
		} else {
			result.WriteByte('1')
		}
	}
	return result.String()
}

func bitsToInt(bits string) int {
	result := new(big.Int)
	result.SetString(bits, 2)
	return int(result.Int64())
}

func intToBits(value int, bitLength int) string {
	bitString := strconv.FormatInt(int64(value), 2)
	if len(bitString) < bitLength {
		padding := bitLength - len(bitString)
		bitString = strings.Repeat("0", padding) + bitString
	}
	return bitString
}

func printStyled(text string) {
	reset := "\033[0m"
	styles := map[string]string{
		"red":       "\033[31m",
		"green":     "\033[32m",
		"yellow":    "\033[33m",
		"cyan":      "\033[36m",
		"bold":      "\033[1m",
		"underline": "\033[4m",
	}
	for style, code := range styles {
		placeholder := fmt.Sprintf("{%s}", style)
		text = strings.ReplaceAll(text, placeholder, code)
	}
	text = strings.ReplaceAll(text, "{reset}", reset)
	fmt.Print(text + reset)
}

func pressAnyKey() {
	printStyled("\n{bold}{cyan}Press any key to continue...\n")
	bufio.NewReader(os.Stdin).ReadByte()
}

func choice(message string, first string, second string, letter1 string, letter2 string) bool {
	var userInput string
	letter1 = strings.ToUpper(letter1)
	letter2 = strings.ToUpper(letter2)
	reader := bufio.NewReader(os.Stdin)
	for {
		printStyled("{bold}{cyan}" + message)
		printStyled("\nPlease Choose {bold}{cyan}(" + letter1 + ") {reset}" + first + ", or {bold}{cyan}(" + letter2 + ") {reset}" + second + ": ")
		userInput, _ = reader.ReadString('\n')
		userInput = strings.ToUpper(strings.TrimSpace(userInput))
		if userInput == letter1 || userInput == letter2 {
			break
		}
		printStyled("\n{red}Invalid choice!\n")
	}
	return userInput == letter1
}
func getPassword(recover bool) string {
	reader := bufio.NewReader(os.Stdin)
	var password1, password2 string
	for {
		printStyled("\n{cyan}Enter password: ")
		password1, _ = reader.ReadString('\n')
		password1 = strings.TrimSpace(password1)

		printStyled("{cyan}Confirm the password: ")
		password2, _ = reader.ReadString('\n')
		password2 = strings.TrimSpace(password2)

		if password1 != password2 {
			printStyled("{red}{bold}\nError: Passwords do not match. Try again.")
			continue
		}

		if !recover && isWeakPassword(password1) {
			printStyled("\n{yellow}Warning: Your password is weak. It should be at least 8 characters long\n")
			printStyled("{yellow}and include a mix of uppercase, lowercase, numbers, and special characters.\n")
			printStyled("{yellow}Type {red}'YES'{yellow} if you want to continue with this password: ")
			confirmation, _ := reader.ReadString('\n')
			confirmation = strings.TrimSpace(confirmation)
			if confirmation == "YES" {
				printStyled("\n{green}Ok, weak password accepted.")
				break
			} else {
				printStyled("\n{green}Please enter a stronger password.")
				continue
			}
		} else {
			printStyled("\n{green}Password accepted.")
			break
		}
	}

	if !recover {
		printStyled("\n\n{yellow}Don't forget your password - there is {underline}NO WAY{reset}{yellow} to recover it!\n\n")
	}
	return password1
}
func main() {

	printStyled("\n\n{cyan}{bold}{underline}Welcome to the wallet word scrambler\n\n")
	printStyled("A password will be used to scramble your backup words\n")
	printStyled("In order for the program to work a wordlist must be present at the current folder\n")
	printStyled("The wordlist must match your wallet backup\n\n")
	printStyled("{red}Warning:\n")
	printStyled("{yellow}This program is meant to run on a fresh formatted and air-gapped machine\n")
	printStyled("{yellow}It is not safe to run it on a machine connected to any kind of network\n")
	printStyled("{yellow}{bold}SECURE WIPE{reset}{yellow} your machine after use\n\n")

	pressAnyKey()

	selectWordList()

	recover := choice("Do you want to recover a wallet or create (scramble) a new one?", "Recover", "Create", "R", "C")
	reader := bufio.NewReader(os.Stdin)
	var password = getPassword(recover)
	printStyled("\n{cyan}Gererating a key from your password.\nThis is an intentionaly slow process, please wait.\n\n")
	argon2Hash := hashRepeatedly([]byte("this is just something to season the dish"), 4847868)
	memory := uint32(1024 * 1024)
	time := uint32(16)
	threads := uint8(1)
	keyLen := uint32(64)
	rounds := 10
	progress := 0
	for i := 1; i <= rounds; i++ {
		argon2Hash = hashRepeatedly(argon2Hash, i)
		argon2Hash = argon2.IDKey([]byte(password), argon2Hash, time, memory, threads, keyLen)
		progress = (i * 100) / rounds
		showProgress := ".." + strconv.Itoa(progress) + "%"
		printStyled(showProgress)
	}

	keyBits := bytesToBitString(argon2Hash)
	wordBitSize := int(math.Log2(float64(len(words))))
	keyBitsWords := splitString(keyBits, wordBitSize)

	printStyled("\n\n{green}Key generated.\n")

	var walletWordCount int
	for {
		printStyled("\n{cyan}Enter the number of words in your wallet (12-33): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		walletWordCount, err = strconv.Atoi(input)
		if err == nil && walletWordCount >= 12 && walletWordCount <= 33 {
			break
		}
		fmt.Println("Invalid input. Please enter a number between 12 and 33.")
	}
	newWords := make([]string, walletWordCount)
	for i := 0; i < walletWordCount; i++ {
		var word string
		for {
			fmt.Printf("Enter word %d: ", i+1)
			word, _ = reader.ReadString('\n')
			word = strings.TrimSpace(word)
			wordIndex := findWord(word)
			if wordIndex == -1 {
				suggestions := suggestWords(word)
				printStyled("{red}Invalid word. The word must exist in the wordlist.\n")
				if len(suggestions) > 0 {
					printStyled("{yellow}Did you mean:\n")
					for _, suggestion := range suggestions {
						printStyled(suggestion + "\n")
					}
				}
			} else {
				wordBits := intToBits(wordIndex, wordBitSize)
				xorResult := xorBitStrings(keyBitsWords[i], wordBits)
				newWordIndex := bitsToInt(xorResult)
				newWords[i] = words[newWordIndex]
				break
			}
		}
	}
	if !recover {

		printStyled("\n{bold}{underline}{cyan}Here are your new wallet words\n\n")
	} else {
		printStyled("\n{bold}{underline}{cyan}Here are your recovered wallet words\n\n")
	}
	numberPadding := len(fmt.Sprintf("%d", len(newWords)))
	for i, word := range newWords {
		fmt.Printf("%*d. %s\n", numberPadding, i+1, word)
	}

	if !recover {
		printStyled("\n\n{yellow}Write the words down and store them in a safe place.\n\n")
	}
	pressAnyKey()
}
