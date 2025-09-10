package basic

import "fmt"

func main() {

	// func <name>(<parameters>) <returnType> {
	// code block
	// return <value>
	// }

	// function name with Captilal letter - public, example fmt.Println
	// function name with small letter - private

	// sum := add(1,2)
	// fmt.Println(sum)
	// fmt.Println(add(2,3))

	// func(){
	// 	fmt.Println("Anonymous function")
	// }() 

	// greet := func(){
	// 	fmt.Println("Anonymous function2")
	// }
	// greet()

	// operation := add
	// result := operation(3,5)
	// fmt.Println(result)

	result := applyOperation(4, 5, add)
	fmt.Println(result)

	multiplyBy2 := createMultiplier(2)
	fmt.Println(multiplyBy2(3)) // Output: 6
}

func add(a, b int) int {
	return a+b
}

// function that takes another function as parameter
func applyOperation(x int, y int, operation func(int, int)int) int {
	return operation(x, y)
}

// function that returns another function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

