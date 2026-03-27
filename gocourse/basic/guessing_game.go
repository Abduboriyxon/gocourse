package basic

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	// Generate a random number between 1 and 100
	target := r.Intn(100) + 1

	// welcom message
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have selected a number between 1 and 100.")
	fmt.Println("Can you guess what it is?")

	var guess int
	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)

		// check the guess
		if guess == target {
			fmt.Println("Congratulations! You guessed the correct number:", target)
			break
		} else if guess < target {
			fmt.Println("Too low! Try again.")
		} else {
			fmt.Println("Too high! Try again.")
		}
	}
}