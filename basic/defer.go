package basic

import "fmt"

func main() {

	process(10)

}

func process(i int) {
	defer fmt.Println("Deferred of i:", i)
	defer fmt.Println("first deferred call executed")
	defer fmt.Println("second deferred call executed")
	defer fmt.Println("third deferred call executed")
	fmt.Println("Normal call executed") 
	i++
	fmt.Println("Value of i:", i)
}
// the one defer earlier will be processed last and the one defer later will be processed first