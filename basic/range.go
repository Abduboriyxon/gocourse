package basic

import "fmt"
func main() {

	message := "Hello World"
	for i, v := range message {
		fmt.Printf("Index: %d, Character: %c\n", i, v)
	}

// array slice string range iterat in order from 0 to the last index 
// for maps range iterat over key value but in random order 
// and for channel range iterat over values received from the channel until it is closed 
}