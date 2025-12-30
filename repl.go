package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replLoop(ctx *Context) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for true {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			text := cleanInput(scanner.Text())
			cmd, ok := commands[text[0]]
			if !ok {
				fmt.Println("Unkown command")
				continue
			}

			arg := ""
			if len(text) > 1 {
				arg = text[1]
			}
			err := cmd.callback(ctx, arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %v", err)
			}

		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	}
}

func cleanInput(text string) []string {
	words := []string{}
	for _, word := range strings.Fields(text) {
		words = append(words, strings.ToLower(word))
	}
	return words
}
