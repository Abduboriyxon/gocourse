package advanced

import (
	"fmt"
	"time"
)

// Goroutines are just functions that leave the main thread and run in the backgrund and come back to join the main thread once the funtions are finishid/ready to return any values
// Gorountines do not stop the programm flow and are non blocking

func main() {

	var err error
	fmt.Println("Begining program")
	go sayHello()
	fmt.Println("After sayHello function")

	go func(){
		err = doWork()
	}()

	// err = go doWork() // this isn not accepted
	go printNumbers()
	go printLetters()

	// time.Sleep(2 * time.Second)
	if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("Work completed successfully")
	}

	time.Sleep(2 * time.Second)

}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from goroutine!")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println("Number:",i,time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println("Letter:",string(letter),time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	// Simulate work
	time.Sleep(1 * time.Second)
	return fmt.Errorf("an error occurred in doWork")
}