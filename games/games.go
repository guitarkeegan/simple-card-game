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

func playRand() {

	var choice int
	fmt.Println("Guess a number between 1 and 10")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("problem gettig input")
	}

	num := rand.Intn(10) + 1

	fmt.Printf("the random number is: %d, and you chose %d\n", num, choice)
	if num == choice {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost. Sorry!")
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

	fmt.Printf("You have: %d\n", randIdx1)
	fmt.Println("Stay with this or choose another card? True or False?")
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
	} else {
		fmt.Println("Sorry, you lost.")
	}
}
