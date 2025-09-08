package basic

import (
	"fmt"
	"math"
)
func main(){
	// Variable declaration
	var a,b = 10, 3
	var result int

	result = a + b // Addition
	fmt.Println("Addition:", result)

	result = a - b // Subtraction
	fmt.Println("Subtraction:", result)

	result = a * b // Multiplication
	fmt.Println("Multiplication:", result)

	result = a / b // Di vision
	fmt.Println("Division:", result)

	result = a % b // Remainder
	fmt.Println("Remainder:", result)

	const p float64 = 22.0 / 7
	fmt.Println("Value of p:", p)

	// Overflow with singed integers
	var maxInt int64 = 9223372036854775807 // Max value for int64
	fmt.Println("Max Int64:", maxInt)

	maxInt += 1
	fmt.Println("Overflowed Max Int64:", maxInt)

	// Overflow with unsigned integers
	var uMaxInt uint64 = 18446744073709551615 // Max value for uint64
	fmt.Println("Max Uint64:", uMaxInt) 

	uMaxInt += 1
	fmt.Println("Overflowed Max Uint64:", uMaxInt)

	// Underflow with float64
	var smallFloat float64 = 1.0e-323
	fmt.Println("Small Float64:", smallFloat)

	smallFloat /= math.MaxFloat64
	fmt.Println("Underflowed Small Float64:", smallFloat) 

}
