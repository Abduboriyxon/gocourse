package intermediate

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Command: ", os.Args[0])

	for i, agr := range os.Args {
		fmt.Println("Argument", i, ":", agr)
	}

	// Difine flags
	var name string
	var age int
	var male bool

	flag.StringVar(&name, "name", "John", "Your full name")
	flag.IntVar(&age, "age", 18, "Age of the user")
	flag.BoolVar(&male, "male", true, "gender of the user")

	flag.Parse()

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Male:", male)


}