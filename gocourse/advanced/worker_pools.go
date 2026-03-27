package advanced

import (
	"fmt"
	"time"
)

// ========== Basic Worker Pool Pattern ==========
// func worker(id int, tasks <- chan int, results chan <- int) {
// 	for task := range tasks {
// 		fmt.Printf("Worker %d processing task %d\n", id, task)
// 		// Simulate doing some work
// 		time.Sleep(time.Second)
// 		results <- task * 2
// 	}
// }
// func main() {
// 	numWorkers := 3
// 	numJob := 10
// 	tasks := make(chan int, numJob)
// 	results := make(chan int, numJob)
// 	// Create workers
// 	for i := range numWorkers{
// 		go worker(i, tasks, results)
// 	}
// 	// Send values to tasks channel
// 	for i := range numJob{
// 		tasks <- i
// 	}
// 	close(tasks)
// 	// Collect results
// 	for range numJob{
// 		result := <- results
// 		fmt.Printf("Result: %d\n", result)
// 	}
// }

// ========== Ticket Processing Worker Pool Example ==========
type ticketRequest struct {
	personID int
	numTickets int
	cost int
}
func ticketProcessor(requests <- chan ticketRequest, results chan <- int) {
	for req := range requests{
		fmt.Printf("Processing %d ticket(s) of person %d total cost $%d\n", req.numTickets, req.personID, req.cost)
		// Simulate processing time
		time.Sleep(time.Second)
		results <- req.personID
	}
}
func main() {
	numRequests := 5
	price := 5
	ticketRequests := make(chan ticketRequest, numRequests)
	tickectResults := make(chan int, numRequests)
	// start ticket processors
	for range 3{
		go ticketProcessor(ticketRequests, tickectResults)
	}
	// send ticket requests
	for i := range numRequests{
		ticketRequests <- ticketRequest{personID: i+1, numTickets: (i + 1)*2, cost: (i + 1) * price}
	}
	close(ticketRequests)
	for range numRequests{
		fmt.Printf("Ticket for personID %d processed successfully\n", <- tickectResults)
	}
}