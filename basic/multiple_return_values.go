package basic

import (
	"errors"
	"fmt"
)

func main() {

	// func functionName(parametr1 type1, parametr2 type2) (returnType1, returnType2) {
	// code block
	// return value1, value2
	// }

	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)

	result, err := compare(5, 10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}


}

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func compare(a,b int) (string, error) {
	if a>b {
		return "a is greater than b", nil
	} else if a<b {
		return "a is less than b", nil
	} else {
		return "", errors.New("Unable to compare which is greater")
	}
}