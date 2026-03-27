package intermediate

import (
	"errors"
	"fmt"
)

// func sqrt(x float64) (float64, error) {
// 	if x < 0 {
// 		return 0, errors.New("cannot compute square root of a negative number")
// 	}
// 	// Dummy implementation for the sake of example
// 	return 1, nil
// }

func process (data []byte) error {
	if len(data) == 0 {
		return errors.New("data cannot be empty")
	}
	// Process the data...
	return nil
}

func main() {

	// result, err := sqrt(16)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Square root:", result)

	// result1, err1 := sqrt(-16)
	// if err1 != nil {
	// 	fmt.Println("Error:", err1)
	// 	return
	// }
	// fmt.Println("Square root:", result1)

	// data := []byte{}
	// err := process(data)
	// // if err := process(data); err != nil {}
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Data processed successfully")

	// --- error interface implementation
	// err1 := eprocess()
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }
	
	err := readData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data read successfully")
}

type myError struct {
	msg string
}

func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.msg)
}

func eprocess() error{
	return &myError{"something went wrong"}
}

func readData() error{
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData failed: %w", err)
	}
	return nil
}

func readConfig() error{
	return errors.New("config error")
}