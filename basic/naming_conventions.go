package baisc

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}
type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}
func main() {
	// PascalCase for exported names
	// Ex: MyVariable, MyFunction, MyStruct
	// Structs, Interfaces, enums

	// snake_case like for files names
	// Ex: my_variable, my_function

	// UPPERCASE for constants
	// Ex: const PI = 3.14

	// mixedCase for variables and functions
	// Ex: myVariable, myFunction

	const PI = 3.14
	var employeeID = 12345
	fmt.Println(PI, employeeID)
}
