package advanced

import (
	"fmt"
	"sync"
	"time"
)
// ======= BASIC EXAMPLE WITHOUT USING CHANNELS =======
// func worker(id int, wg *sync.WaitGroup){
// 	defer wg.Done()
// 	// wg.Add(1) // WRONG PRACTICE: This should be called before starting the goroutine
// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(time.Second) // simulate some time spent on processing the task
// 	fmt.Printf("Worker %d finished\n", id)
// }
// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	wg.Add(numWorkers) // Correct place to add the number of goroutines to wait for
// 	for i := range numWorkers{
// 		go worker(i, &wg)
// 	}
// 	wg.Wait() // blocking mechanism
// 	fmt.Println("All workers completed")
// }

// ======= EXAMPLE WITH CHANNELS =======
// func worker(id int, tasks <-chan int ,results chan<- int, wg *sync.WaitGroup){
// 	defer wg.Done()
// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(time.Second) // simulate some work
// 	for task := range tasks{
// 		results <- task * 2
// 	}
// 	fmt.Printf("Worker %d finished\n", id)
// }
// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	numJobs := 5
// 	results := make(chan int, numJobs)
// 	tasks := make(chan int, numJobs)

// 	wg.Add(numWorkers)
// 	for i := range numWorkers{
// 		go worker(i+1, tasks ,results, &wg)
// 	}
// 	for i := range numJobs{
// 		tasks <- i + 1
// 	}
// 	close(tasks)
// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()
// 	for result := range results{
// 		fmt.Println("Result:", result)
// 	}
// }

// ======= CONSTRUCTION EXAMPLE =======
type Worker struct{
	ID int
	Task string
}
func (w *Worker) PerformTask(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("WorkerID %d started %s\n", w.ID, w.Task)
	time.Sleep(2 * time.Second) // simulate some work
	fmt.Printf("WorkerID %d completed %s\n", w.ID, w.Task)
}

func main() {
	var wg sync.WaitGroup
	tasks := []string{"digging", "laying brick", "painting"}
	for i, task := range tasks{
		worker := Worker{ID: i+1, Task: task}
		wg.Add(1)
		go worker.PerformTask(&wg)
	}
	// Wait for all workers to complete
	wg.Wait()
	// construction completed
	fmt.Println("All construction tasks completed")
}