package basic

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("This will not be printed before exiting.")

	fmt.Println("Exiting the program.")

	// Exit the program with status code 1
	os.Exit(1)

	// This line will not be executed
	fmt.Println("This line will not be printed.")

}