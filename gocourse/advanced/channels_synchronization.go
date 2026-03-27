package advanced

import (
	"fmt"
	"time"
)

// func main() {
// 	done := make(chan int)
// 	go func(){
// 		fmt.Println("working...")
// 		time.Sleep(2*time.Second)
// 		done <- 0
// 	}()
// 	<- done
// 	fmt.Println("Finished.")
// }

func main() {
	ch := make(chan int)
	go func(){
		fmt.Println("Sending...")
		ch <- 9
		time.Sleep(2*time.Second)
		fmt.Println("Sent.")
	}()
	value := <- ch // blocking until value is sent
	fmt.Println("Finished: ", value)
}


// =========== Synchronizing Multiple Goroutines and ensuring that all gorutines are complete ===========
// func main() {
// 	numGoroutines := 3
// 	done := make(chan int, 3)
// 	for i := range numGoroutines{
// 		go func(id int){
// 			fmt.Printf("Goroutine %d working...\n", id)
// 			time.Sleep(time.Second)
// 			done <- id // sending signal of completion
// 		}(i)
// 	}
// 	for range numGoroutines{
// 		<- done // Wait for each goroutine to finish, wait all goroutines to signal completion
// 	}
// 	fmt.Println("All goroutines finished.")
// }

// ========== synchronizing data exchange ==========
// func main(){
// 	data := make(chan string)
// 	go func(){
// 		for i := range 5{
// 			data <- "hello " + string('0'+i)
// 			time.Sleep(100 * time.Millisecond)
// 		}
// 		close(data)
// 	}()
// 	// close(data) // channel closed before gorutine could send a value to channel
// 	for value := range data {
// 		fmt.Println("received value: ", value, ":", time.Now())
// 	} // loops over only on active channel, create receiver each time and stops creating receiver (looping) once channel is closed

// }