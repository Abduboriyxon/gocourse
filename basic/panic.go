package basic

import "fmt"

func main() {

	// panic(interface{})

	// Example of valid input
	process(10)

	// Example of invalid input that causes panic
	process(-5)

}

func process(input int) {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0 {
		fmt.Println("Before panic")
		panic("Negative input not allowed")
	}
	fmt.Println("Processing input:", input)
}