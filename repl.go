package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// start the infinite Repl
func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" >")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		command := cleaned[0]
		switch command {
			case "exit":
				os.Exit(0)
		}

		fmt.Println("echoing: ", cleaned)
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}