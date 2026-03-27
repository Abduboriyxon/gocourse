package intermediate

import "fmt"

func main() {

	// Printing functions
	// fmt.Print("Hello")
	// fmt.Print("World!")
	// fmt.Print(12, 456)

	// fmt.Println("Hello")
	// fmt.Println("World!")
	// fmt.Println(12, 456)

	// name := "Alice"
	// age := 30
	// fmt.Printf("%s is %d years old.\n", name, age)
	// fmt.Printf("Hexadecimal: %X\nbinary %b\n", age, age)

	// formatting functions
	// s := fmt.Sprint("hello", "world", 123, 456)
	// fmt.Println(s)
	// // fmt.Print(s)
	// // fmt.Print(s)

	// s = fmt.Sprintln("hello", "world", 123, 456)
	// fmt.Println(s)

	// sf := fmt.Sprintf("%s is %d years old.\n", name, age)
	// fmt.Println(sf)

	// Scanning functions
	// var name string
	// var age int
	// fmt.Print("Enter your name and age: ")
	// // fmt.Scan(&name, &age) 
	// // fmt.Scanln(&name, &age)
	// fmt.Scanf("%s %d",&name, &age) 
	// fmt.Printf("Hello %s, you are %d years old.\n", name, age)

	// Error formatting functions
	// err := checkAge(19)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	var name string
var age int
fmt.Print("Enter your name and age: ")
// fmt.Scan(&name, &age) 
// fmt.Scanln(&name, &age)
fmt.Scanf("%s %d",&name, &age) 
fmt.Printf("Hello %s, you are %d years old.\n", name, age)

}

// func checkAge(age int) error {
// 	if age < 18 {
// 		return fmt.Errorf("invalid age to drive : %d", age)
// 	}
// 	return nil
// }