package main

import (
	"fmt"

	"github.com/guitarkeegan/simple-card-game/games"
)

func main() {
	var userInput string

	fmt.Print("Play high or rand? ")
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if userInput == "rand" || userInput == "high" {
		fmt.Printf("Great! lets play %s\n", userInput)
		games.PlayGame(&userInput)
	} else {
		fmt.Println("Sorry, we don't have that one.")
	}
}
