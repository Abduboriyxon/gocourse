package main

import "fmt"

func producer(ch chan <- int) {
	for i := range 6 {
		ch <- i
	}
	close(ch)
}
func filter(in <- chan int, out chan <- int){
	for val := range in{
		if val % 2 == 0{
			out <- val
		}
	}
	close(out)
}
func main(){
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch1)
	go filter(ch1, ch2)

	for val := range ch2{
		fmt.Println("Received: ", val)
	}
}



// func main(){
// 	ch := make(chan int)
// 	go func(){
// 		close(ch)
// 		// close(ch)
// 	}()
// 	time.Sleep(time.Second)
// }

// === range over closed channel ===
// func main(){
// 	ch := make(chan int)
// 	go func(){
// 		for i := range 5{
// 			ch <- i
// 		}
// 		close(ch)
// 	}()
// 	for val := range ch{
// 		fmt.Println("Received: ",val)
// 	}
// }

// === receiving from a closed channel example ===
// func main() {
// 	ch := make(chan int)
// 	close(ch)
// 	val, ok := <- ch
// 	if !ok {
// 		fmt.Println("Channel is closed")
// 		return
// 	} 
// 	fmt.Println(val)
// }

// === simple closing channel example ===
// func main() {
// ch := make(chan int)
// go func(){
// 	for i := range 5{
// 		ch <- i
// 	}
// 	close(ch)
// }()
// for val := range ch{
// 	fmt.Println(val)
// }
// }