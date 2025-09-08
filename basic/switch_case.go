package basic

import "fmt"

func main() {

	// Switch statement
	// switch expression {
	// 	case value1:
	// 		// Code to be executed if expression equals value1
	// 	case value2:
	// 		// Code to be executed if expression equals value2
	// 	default:
	// 		// Code to be executed if expression doesn't match any case
	// }

	// fruit := "apple"
	// switch fruit {
	// case "apple":
	// 	fmt.Println("This is an apple.")
	// case "banana":
	// 	fmt.Println("This is a banana.")
	// default:
	// 	fmt.Println("Unknown fruit.")
	// }

	// Switch with multiple cases
	// day := "monday"
	// switch day {
	// case "monday", "tuesday", "wednesday", "thursday", "friday":
	// 	fmt.Println("It's a weekday.")
	// case "saturday", "sunday":
	// 	fmt.Println("It's the weekend!")
	// default:
	// 	fmt.Println("Unknown day.")
	// }

	// number := 15
	// switch {
	// case number < 10:
	// 	fmt.Println("The number is less than 10")
	// case number >= 10 && number <= 20:
	// 	fmt.Println("The number is between 10 and 20") 
	// 	default:
	// 	fmt.Println("The number is greater than 20")
	// }

	// num := 2 
	// switch {
	// case num> 1:
	// 	fmt.Println("The number is greater than 1")
	// 	fallthrough
	// case num == 2:
	// 	fmt.Println("The number is equal to 2")
	// default:
	// 	fmt.Println("not two")
	// }

	checkType(42)
	checkType("hello")
	checkType(3.14)
	checkType(true)

}
func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("it's an integer")
	case string:
		fmt.Println("it's a string")
	case float64:
		fmt.Println("it's a float")
	default:
		fmt.Println("unknown type")
	}
}