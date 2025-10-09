package intermediate

import "fmt"

type person struct{
	name string
	age int
}

type Employee struct{
	person
	empID string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.age)
}

func (e Employee) introduce() {
	fmt.Printf("Hello, my name is %s and I am %d years old, my epm id: %s\n", e.name, e.age, e.empID)
}

func main() {

	emp := Employee{
		person: person{ name: "Alice", age: 30},
		empID: "E123",
		salary: 75000.00,
	}

	fmt.Println("Employee Name:", emp.name) // Accessing embedded struct field directly
	fmt.Println("Employee Age:", emp.age) // Accessing embedded struct field directly
	fmt.Println("Employee emp ID:", emp.empID) 
	fmt.Println("Employee Salary:", emp.salary) 

	emp.introduce() // Calling method of embedded struct


}
