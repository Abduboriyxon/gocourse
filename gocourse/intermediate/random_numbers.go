package intermediate

import (
	"fmt"
	"math/rand"
	// "time"
)

func main() {

	// val := rand.New(rand.NewSource(42))
	// val1 := rand.New(rand.NewSource(time.Now().Unix()))
	// fmt.Println(val.Intn(101) + 2)
	// fmt.Println(val1.Intn(101) + 2)

	// fmt.Println(rand.Intn(101) + 2)
	// fmt.Println(rand.Float64()) // 0.0 <= n < 1.0
	

	for {
		fmt.Println("Welcom dice game!")
		fmt.Println("1. Rolling the dice...")
		fmt.Println("2. Exit")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		if choice == 2 {
			fmt.Println("Exiting the game. Goodbye!")
			break
		}

		diceRoll := rand.Intn(6) + 1
		diceRoll2 := rand.Intn(6) + 1
		fmt.Printf("You rolled a %d and a %d. Total: %d\n", diceRoll, diceRoll2, diceRoll+diceRoll2)

		fmt.Println("Do you want to roll again? (y/n)")
		var again string
		_, err = fmt.Scan(&again)
		if err != nil || (again != "y" && again != "n") {
			fmt.Println("Invalid input. Please enter 'y' or 'n'.")
			continue
		}
		if again == "n" {
			fmt.Println("Thanks for playing! Goodbye!")
			break
		}
	}
}