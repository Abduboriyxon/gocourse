package advanced

import (
	"fmt"
	"time"
)

// "fmt"
// "time"

// func main(){
// 	// BLOCKING ON RECEIVE ONLY IF THE BUFFER IS EMPTY
// ch := make(chan int, 2)
// go func() {
// 	time.Sleep(2*time.Second)
// 	ch <- 1
// 	ch <- 2
// }()
// fmt.Println("Value: ", <- ch)
// fmt.Println("Value: ", <- ch)
// fmt.Println("End of buffered channel")
// }

// func main() {
// 	// ========================= BLOCKING ON SEND ONLY IF THE BUFFER IS FULL =========================
// 	// make(chan type, capacity)
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	fmt.Println("Receiving from buffer")
// 	go func(){
// 		// fmt.Println("Goroutine started")
// 		time.Sleep(2*time.Second)
// 		fmt.Println("Received: ", <-ch) // ends <-- starts
// 	}()
// 	// fmt.Println("Blocking starts")
// 	ch <- 3 // Blocks because the buffer is full
// 	// fmt.Println("Blocking ends")
// 	// fmt.Println("Received: ", <-ch)
// 	// fmt.Println("Received: ", <-ch)
// 	// fmt.Println("Buffered channel")
// }

func main(){
	// BLOCKING ON RECEIVE ONLY IF THE BUFFER IS EMPTY
ch := make(chan int, 2)
time.Sleep(2*time.Second)


fmt.Println("Value: ", <- ch)
fmt.Println("Value: ", <- ch)
fmt.Println("End of buffered channel")
}