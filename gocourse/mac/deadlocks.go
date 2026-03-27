package mac

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var mu1, mu2 sync.Mutex

	go func ()  {
		mu1.Lock()
		fmt.Println("Goroutine 1 Locked mu1")
		time.Sleep(1 * time.Second)
		mu2.Lock()
		fmt.Println("Goroutine 1 Locked mu2")
		mu1.Unlock()
		mu2.Unlock()
		fmt.Println("Goroutine 1 Completed")
	}()

	go func ()  {
		mu1.Lock()
		fmt.Println("Goroutine 2 Locked mu1")
		time.Sleep(1 * time.Second)
		mu2.Lock()
		fmt.Println("Goroutine 2 Locked mu2")
		mu1.Unlock()
		mu2.Unlock()
		fmt.Println("Goroutine 2 Completed")
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Main function completed")
	// select {}
}