package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// ==== DIFFERENCES BETWEEN context.TODO() AND context.Background() ====
// func main() {
// 	todoContext := context.TODO()
// 	contextBkg := context.Background()

// 	ctx := context.WithValue(todoContext, "name", "John")
// 	fmt.Println(ctx)
// 	fmt.Println(ctx.Value("name"))

// 	ctx1 := context.WithValue(contextBkg, "city", "New York")
// 	fmt.Println(ctx1)
// 	fmt.Println(ctx1.Value("city"))
// }

// func checkEvenOdd(ctx context.Context, number int) string {
// 	select{
// 	case <- ctx.Done():
// 		return "Operation cancelled"
// 	default:
// 		if number%2 == 0 {
// 			return fmt.Sprintf("%d is Even", number)
// 		} else {
// 			return fmt.Sprintf("%d is Odd", number)
// 		}
// 	}
// }

// func main() {
// 	ctx := context.TODO()
// 	result := checkEvenOdd(ctx, 5)
// 	fmt.Println("Result with context.TODO():", result)

// 	ctx = context.Background()
// 	ctx, cancel := context.WithTimeout(ctx, 1 * time.Second)
// 	defer cancel()

// 	result = checkEvenOdd(ctx, 10)
// 	fmt.Println("Result with context.Background():", result)

// 	time.Sleep(2 * time.Second)
// 	result = checkEvenOdd(ctx, 15)
// 	fmt.Println("Result after timeout:", result)
// }

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work cancelled:", ctx.Err())
			return
		default:
			fmt.Println("Doing work...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2 * time.Second) // timer of context starts here. You have this specified time duration to use this context. After this time duration, the context will send a cancelation signal
	defer cancel()
	// ctx, cancel := context.WithCancel(ctx)
	// go func(){
	// 	time.Sleep(2 * time.Second) // simulating a heavy task. Consider this heavy time consuming operation
	// 	cancel() // manually canceling only after the task is finished
	// }()
	ctx = context.WithValue(ctx, "requestID", "cjkkjd33fe")
	go doWork(ctx)
	time.Sleep(3 * time.Second)
	
	requestID := ctx.Value("requestID")
	if requestID != nil {
		fmt.Println("Request ID:",requestID)
	} else {
		fmt.Println("No Request ID found")
	}
	logWithContext(ctx, "This is a log message with context")
}
func logWithContext(ctx context.Context, message string) {
	requestIDVal := ctx.Value("requestID")
	log.Printf("RequestID: %v - %v\n", requestIDVal, message)
}