package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// cliCommand represents a command in the CLI application.
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// commands is a list of available CLI commands.
var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex application",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show a helpful message",
			callback:    commandHelp,
		},
	}
}

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

// commandHelp prints the list of available commands.
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	for _, cmd := range commands {
		fmt.Printf("  %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

// commandExit exits the program.
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
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
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			// Empty input, just reprompt
			continue
		}
		cmdName := words[0]
		cmd, found := commands[cmdName]
		if !found {
			fmt.Printf("Unknown command: %s\n", cmdName)
			continue
		}
		if err := cmd.callback(); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing command %s: %v\n", cmdName, err)
		}

	}
}
