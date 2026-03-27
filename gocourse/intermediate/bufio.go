package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(strings.NewReader("Hello, bufio packageeee!\nHow are you?"))

	// Reading byte slices
	data := make([]byte, 20)
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println("Error reading bytes:", err)
		return
	}
	fmt.Printf("Read %d bytes: %s\n", n, string(data[:n]))

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading string:", err)
		return
	}
	fmt.Println("Read line:", line)

	writer := bufio.NewWriter(os.Stdout)

	// Writing byte slices
	data1 := []byte("Hello bufio packcage!\n")
	n, err = writer.Write(data1)
	if err != nil {
		fmt.Println("Error writing bytes:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	// Flush the buffer to ensure all data is written to os.Stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	// Writing strings
	str := "This is a string.\n"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	// Flush the buffer again
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}



} 