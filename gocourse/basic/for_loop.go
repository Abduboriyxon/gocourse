package basic

import "fmt"

func main() {

	// simple iteration over range
	// for i:=0; i<10; i++ {
	// 	fmt.Println(i)
	// }

	// iteration over collection
	// numbrs := []int{1,2,3,4,5,6}
	// for index, value := range numbrs {
	// 	fmt.Printf("index: %d, value: %d\n", index, value)
	// }

	// for i:=0; i<10; i++{
	// 	if i%2 == 0 {
	// 		continue // continue to next iteration
	// 	}
	// 	fmt.Println("odd numbers: ",i)
	// 	if i == 7 {
	// 		break // exit the loop
	// 	}
	// }

	// Asterisk pattern

	// rows := 5
	// // outer loop
	// for i := 1; i <= rows; i++ {
	// 	// inner loop for spaces efore stars
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	// inner loop for stars
	// 	for k := 1; k <= (2*i - 1); k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println() // move to next line after each row
	// }


}