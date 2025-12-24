package main

import (
	"fmt"
	"time"
)

// func main() {
// 	ticker := time.NewTicker(2*time.Second)
// 	defer ticker.Stop()
// 	// for tick := range ticker.C {
// 	// 	fmt.Println("Tick a:", tick)
// 	// 	fmt.Println("Tick a:", tick)
// 	// }
// 	i := 1
// 	for range 5 {
// 		i *= 2
// 		fmt.Println("Tick:", i)
// 	}
// }

// ======= scheduling logging, periodic tasks, polling for updates =======
// func periodicTask(){
// 	fmt.Println("Performing periodic task at", time.Now())
// }
// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	for {
// 		select{
// 			case <-ticker.C:
// 				periodicTask()
// 		}
// 	}
// }

func main(){
	ticker := time.NewTicker(time.Second)
	stop := time.After(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case tick := <-ticker.C:
			fmt.Println("Tick at:", tick)
		case <-stop:
			fmt.Println("Stopping ticker")
			return
		}
	}
}