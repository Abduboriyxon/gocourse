package intermediate

import "fmt"

func main() {

	// var ptr *int
	var ptr *int
	var a int = 10
	ptr = &a // refrence operator 

	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(*ptr) // dereferencinga a pointer 


	// if ptr == nil {
	// 	fmt.Println("Pointer is nil")
	// }

	modifyValue(ptr)
	fmt.Println(a)

}

func modifyValue(num *int) {
	*num++
}