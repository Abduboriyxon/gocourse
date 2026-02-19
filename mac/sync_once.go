package mac

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize(){
	fmt.Println("This will not be repeated no matter how many times you call this functionusing once.Do()")
}

func main() {

	var wg sync.WaitGroup
	for i := range 5{
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("goroutine #", i)
			once.Do(initialize)
			// initialize()
		}()
	}
	wg.Wait()


}