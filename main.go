package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/sha3"

	"github.com/agnivade/levenshtein"
)

var words []string

func loadWordList() {
	var validFiles []string
	var fileLines []string

	files, err := filepath.Glob("*.txt")
	if err != nil {
		printStyled("\n{red}Error scanning directory for word files\n")
		os.Exit(1)
	}

	if len(files) == 0 {
		printStyled("\n{red}No wordlist txt files found in the current directory!\n")
		os.Exit(1)
	}

	for _, file := range files {
		fileLines, err = readLines(file)
		if err == nil && (len(fileLines) == 1024 || len(fileLines) == 2048) {
			validFiles = append(validFiles, file)
		}
	}

	if len(validFiles) == 0 {
		printStyled("\n{red}No valid word list files found! Each file must contain 1024 or 2048 lines.\n")
		os.Exit(1)
	}

	var selectedFile string

	if len(validFiles) == 1 {
		selectedFile = validFiles[0]
		printStyled("\n{green}Using word list: {bold}" + strings.TrimSuffix(filepath.Base(selectedFile), ".txt") + "\n")
	} else {
		printStyled("\n{cyan}{bold}Select a word list:\n")
		for i, file := range validFiles {
			fmt.Printf("%d. %s\n", i+1, strings.TrimSuffix(filepath.Base(file), ".txt"))
		}
		selectedFile = validFiles[promptFileSelection(len(validFiles))-1]
		printStyled("\n{green}Using word list: {bold}" + strings.TrimSuffix(filepath.Base(selectedFile), ".txt") + "\n")
	}

	words, err = readLines(selectedFile)
	if err != nil {
		printStyled("\n{red}Error loading word list.\n")
		os.Exit(1)
	}

	printStyled("\n{green}Word list successfully loaded!\n\n")
}

func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func promptFileSelection(max int) int {
	reader := bufio.NewReader(os.Stdin)
	var selection int

	for {
		printStyled("{cyan}Enter the number of the file to use: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		selection, err = strconv.Atoi(input)
		if err == nil && selection >= 1 && selection <= max {
			break
		}
		printStyled("{red}Invalid selection. Please try again.\n")
	}

	return selection
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
	result := []string{}
	for i, s := range suggestions {
		if i >= 3 {
			break
		}
		result = append(result, s.word)
	}
	return result
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

	loadWordList()

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
						printStyled("{white}" + suggestion + "\n")
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
