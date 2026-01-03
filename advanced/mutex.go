package advanced

import (
	"fmt"
	"sync"
)

// type counter struct{
// 	mu sync.Mutex
// 	count int
// }

// func (c *counter) increment() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.count++
// }

// func (c *counter) getValue() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.count
// }

// func main() {

// 	var wg sync.WaitGroup
// 	counter := &counter{}
// 	numGoroutines := 10

// 	// wg.Add(numGoroutines)
// 	for range numGoroutines{
// 		wg.Add(1)
// 		go func(){
// 			defer wg.Done()
// 			for range 1000{
// 				counter.increment()
// 				// counter.count++
// 			}
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Printf("Final Counter Value: %d\n", counter.getValue())
// }

func main() {
	var wg sync.WaitGroup
	var counter int
	var mu sync.Mutex

	numGoroutines := 5
	wg.Add(numGoroutines)

	increment := func() {
		defer wg.Done()
		for range 1000 {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}
	for range numGoroutines{
		go increment() 
	}
	wg.Wait()
	fmt.Printf("Final Counter Value: %d\n", counter)
}