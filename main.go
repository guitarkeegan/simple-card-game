package main

import (
	"fmt"
)

func main() {
	var userInput string

	fmt.Print("Enter your input: ")
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if userInput == "rand" || userInput == "high" {
		fmt.Printf("Great! lets play %s\n", userInput)
	} else {
		fmt.Println("Sorry, we don't have that one.")
	}
}
