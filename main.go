package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// cleanInput normalizes the input string by lowercasing, removing
// punctuation (replaced with spaces), and splitting on whitespace.
func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	var b strings.Builder
	for _, r := range lower {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			b.WriteRune(r)
		} else {
			// replace punctuation/other with a space so Fields will split
			b.WriteRune(' ')
		}
	}
	return strings.Fields(b.String())
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Print prompt
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			// If Scan returned false, check for error or EOF
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Error reading input:", err)
			}
			// EOF (Ctrl+D) or other termination: exit loop
			fmt.Println() // newline after prompt if EOF
			break
		}

		line := scanner.Text()
		words := cleanInput(line)
		fmt.Printf("Your command was: %v\n", words[0])
	}
}
