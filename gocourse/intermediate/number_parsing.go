package intermediate

import (
	"fmt"
	"strconv"
)

func main() {

	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
	}
	fmt.Println("Converted number:", num)
	fmt.Println("Converted number:", num+1)

	numistr, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
	}
	fmt.Println("Converted number using ParseInt:", numistr)

	floatstr := "123.45"
	floatNum, err := strconv.ParseFloat(floatstr, 64)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
	}
	fmt.Printf("Converted float number: %.2f\n", floatNum)

	binaryStr := "1010"
	binaryNum, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error converting binary string to integer:", err)
	}
	fmt.Println("Converted binary number:", binaryNum)

	HexStr := "FF"
	hex, err := strconv.ParseInt(HexStr, 16, 64)
	if err != nil {
		fmt.Println("Error converting binary string to integer:", err)
	}
	fmt.Println("Converted hex:", hex)

}