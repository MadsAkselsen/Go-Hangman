package main

import (
	"fmt"
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

	targetWord := getRandomWord()
	targetWord = "United States of America" // for testing
	guessedLetters := initializeGuessedWords(targetWord)
	printGameState(targetWord, guessedLetters)

	guessedLetters['e'] = true
	printGameState(targetWord, guessedLetters)

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

func initializeGuessedWords(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true

	return guessedLetters
}

func getRandomWord() string {
	return dictionary[rand.Intn(len(dictionary))]
}

func printGameState(targetWord string, guessedLetters map[rune]bool) {
	for _, ch := range targetWord {
		if ch == ' ' {
			fmt.Print(" ")
		} else if guessedLetters[unicode.ToLower(ch)] == true {
			fmt.Printf("%c", ch)
		} else {
			fmt.Print("_")
		}

		fmt.Print(" ")
	}

	fmt.Println()
}