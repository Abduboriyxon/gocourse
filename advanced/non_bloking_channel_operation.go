package advanced

import (
	"fmt"
	"time"
)

func main() {

	// ch := make(chan int)
	// ==== Non-blocking receive Operations ====
	
	// select {
	// case val := <- ch:
	// 	fmt.Println("Received:", val)
	// default:
	// 	fmt.Println("No value received")
	// }

	// ==== Non-blocking send Operations ====

	// select{
	// case ch <- 42:
	// 	fmt.Println("Sent message")
	// default:
	// 	fmt.Println("No message sent")
	// }

	// ==== non-blocking operation in real time system ====

	data := make(chan int)
	quit := make(chan bool)

	go func(){
		for {
			select{
			case d := <- data:
				fmt.Println("Received data:", d)
			case <- quit:
				fmt.Println("Quitting...")
				return
			default:
				fmt.Println("Waiting for data...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(time.Second)
	}

	quit <- true
}