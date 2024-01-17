package games

import (
	"fmt"
	"math/rand"
)

func PlayGame(userInput *string) {
	if *userInput == "rand" {
		playRand()
	} else if *userInput == "high" {
		playHigh()
	} else {
		fmt.Println("exiting because there is no game of that name...")
	}
}

func getUserInput() int {
	var choice int
	fmt.Println("Guess a number between 1 and 10")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("problem gettig input")
	}
	return choice
}

func playAgain() bool {
	var choice bool
	fmt.Println("Play again? 1 for yes, 0 for no.")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("problem gettig input")
	}
	return choice
}

func playRand() {

	gameOn := true
	for gameOn {

		num := rand.Intn(10) + 1

		var choice int

		for choice != num {

			choice = getUserInput()

			fmt.Printf("You chose %d\n", choice)
			if num > choice {
				fmt.Println("Too small")
			} else if num < choice {
				fmt.Println("Too big")
			} else {
				fmt.Printf("You got it! Random number way %v\n", num)
			}
		}
		gameOn = playAgain()
	}
}

func playHigh() {
	cards := []int{}
	for i := 2; i < 11; i++ {
		cards = append(cards, i)
	}
	randIdx1 := rand.Intn(11-2) + 2
	randIdx2 := rand.Intn(11-2) + 2

	var draw bool

	fmt.Printf("You have: %d.\n", randIdx1)
	fmt.Println("Stay with this (0) or choose another card (1)?")
	_, err := fmt.Scanln(&draw)
	if err != nil {
		fmt.Println(err)
	}

	if draw {
		randIdx1 = rand.Intn(11-2) + 2
	}

	fmt.Printf("You have %v, and your opponent has %v\n", randIdx1, randIdx2)

	if randIdx1 > randIdx2 {
		fmt.Println("You won!")
	} else if randIdx1 == randIdx2 {
		fmt.Println("Tie!")
	} else {
		fmt.Println("Sorry, you lost.")
	}
}
