package advanced

import "fmt"

func main() {
	ch := make(chan int)
	producer(ch)
	consumer(ch)
}
// Send oly channel
func producer(ch chan <- int){
	go func(){
		for i := range 5{
			ch <- i
		}
		close(ch)
	}()
}
// Receive only channel
func consumer(ch <- chan int){
	for value := range ch {
		fmt.Println("Received:", value)
	}
}

// go func(ch chan <- int){
// 	for i := range 5{
// 		ch <- i
// 	}
// 	close(ch)
// }(ch)
// 	for value := range ch {
// 		fmt.Println("Received:", value)
// 	}