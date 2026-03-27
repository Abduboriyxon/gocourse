package intermediate

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(){ 
		fmt.Println("Closing file")
		file.Close()
	}()
	fmt.Println("File opened successfully")

	// read the contest of the file
// 	data := make([]byte, 1024)
// 	_, err = file.Read(data) // buffer to read data into
// 	if err != nil {
// 		fmt.Println("Error reading file:", err)
// 		return
// 	}
// 	fmt.Println("File content:", string(data))


	// read line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
	}
	err = scanner.Err() 
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}