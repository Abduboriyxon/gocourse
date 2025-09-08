package basic

import "fmt"

func main() {

	// i := 1
	// for i<=5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for as while with break
	// sum := 0
	// for {
	// 	sum+=10
	// 	fmt.Println(sum)
	// 	if sum>=50 {
	// 		break
	// 	}
	// }

	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("Odd numbers:",num)
		num++ // Increment num to avoid infinite loop and -- decrement to go backwards
	}

}