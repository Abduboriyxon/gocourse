package advanced

import (
	"fmt"
	"time"
)

// import "fmt"

func unbuffered() {

	ch := make(chan int)
	go func() {
		// ch <- 1
		fmt.Println("Sent 111 to channel")
		time.Sleep(1 * time.Second)
		fmt.Println("Sent 1 to channel")
	}()
	go func() {
		fmt.Println("Sent 22 to channel")
		ch <- 1
		fmt.Println("Sent 2 to channel")
	}()
	reciver := <- ch
	fmt.Println(reciver)
	fmt.Println("End of main")

}