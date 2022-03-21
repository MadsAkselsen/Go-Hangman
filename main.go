package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
	"unicode"
)

var dictionary = []string{
	"Zombie",
	"Gopher",
	"United States of America",
	"Indonesia",
	"Nazism",
	"Apple",
	"Programming",
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// ### STATE ###
	targetWord := getRandomWord()
	guessedLetters := initializeGuessedLetters(targetWord)
	hangmanState := 6
	printGameState(targetWord, guessedLetters, hangmanState)

	// Printing game state
	//	* Print word you're guessing
	//	* Print hangman state
	// H _ _ _ m _ n

	// Read user input
	// * validate it (e.g. only letters)
	// Determine if the letter is a correct guess or not
	//	* if correct, update guessed letters
	//	* if incorrect, update the hangman state
	// If word is guessed -> game over, you win
	// If hangman is complete -> game over, you lose
}

func initializeGuessedLetters(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true

	return guessedLetters
}

func getRandomWord() string {
	return dictionary[rand.Intn(len(dictionary))]
}

func printGameState(
	targetWord string,
	guessedLetters map[rune]bool,
	hangmanState int,
) {
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	fmt.Println()
	fmt.Println(getHangmanDrawing(hangmanState))
}

func getWordGuessingProgress(targetWord string, guessedLetters map[rune]bool) string {
	result := ""
	for _, ch := range targetWord {
		if ch == ' ' {
			result += " "
		} else if guessedLetters[unicode.ToLower(ch)] == true {
			result += fmt.Sprintf("%c", ch)
		} else {
			result += "_"
		}

		result += " "
	}

	return result
}

func getHangmanDrawing(hangmanState int) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("./states/hangman%d", hangmanState))
	if err != nil {
		panic(err)
	}

	return string(data)
}