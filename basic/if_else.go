package basic

import "fmt"

func main() {

	// age:= 20
	// if age >= 18 {
	// 	println("You are a adult.")
	// }

	// temprature := 25
	// if temprature >= 30 {
	// 	fmt.Println("It's a hot day.")
	// } else {
	// 	fmt.Println("It's a cold day.")
	// }

	// score := 85
	// if score >= 90 {
	// 	fmt.Println("Grade: A")
	// } else if score >= 80 {
	// 	fmt.Println("Grade: B")
	// } else if score >= 70 {
	// 	fmt.Println("Grade: C")
	// } else {
	// 	fmt.Println("Grade: F")
	// }

	// num := 18
	// if num%2 == 0 {
	// 	if num%3 == 0 {
	// 		fmt.Println(num, "is an even number and divisible by 3.")
	// 	} else {
	// 		fmt.Println(num, "is an even number but not divisible by 3.")
	// 	}
	// } else {
	// 	fmt.Println(num, "is an odd number.")
	// }
	
	// || OR
	// && AND

	if 10%2 == 0 || 5%2 == 0 {
		fmt.Println("At least one of the numbers is even.")
	} else {
		fmt.Println("Both numbers are odd.")
	}

	if 11%2 == 0 && 4%2 == 0 {
		fmt.Println("Both numbers are even.")
	} else {
		fmt.Println("At least one of the numbers is odd.")
	}
}