package basic

import "fmt"

func main() {

	process()
	fmt.Println("Program continues after recovery")


}

func process() {
	defer func() {
		//if r := recover(); r != nil {
		r := recover()
		if r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("Stared processing")
	panic("Something went wrong!")
}