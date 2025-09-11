package baisc

import "fmt"

func main() {

	// ... Ellipsis
	// func functionName(param1 type1, param2 type2, params ...typeN) returnType {
	//      function body
	// }  

	// fmt.Println(sum("The total is:", 1, 2, 3, 4, 5))
	// statement, total := sum("Total of 1 2 3 is", 1,2,3)
	// fmt.Println(statement, total)

	statement, total := sum(1,20,30,40,50)
	fmt.Println(statement, total)

	numbers := []int{1,2,3,4,5,10}
	sequence2, total2 := sum(2, numbers...)
	fmt.Println(sequence2, total2)
}

func sum(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}