package mac

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Second * 1)
	quit := make(chan string)

	go func(){
		time.Sleep(time.Second * 5)
		close(quit)
	}()

	for {
		select{
		case <- ticker.C:
			fmt.Println("tick")
		case <- quit:
			fmt.Println("quit...")
			return
		}
	}

}