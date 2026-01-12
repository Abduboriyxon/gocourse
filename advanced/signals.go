package advanced

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	pid := os.Getpid()
	fmt.Println("Process ID:", pid)
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	// Notify channel on interrupt or terminate signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGUSR1)

	go func(){
		sig := <- sigs
		fmt.Println("recieved signal:", sig)
		done <- true
	}()

	go func(){
		for{
			select{
				case <-done:
					fmt.Println("Stoppig work due to signal")
					// os.Exit(0)
					return
				default:
					fmt.Println("Working...")
					time.Sleep(time.Second)
			}
		}
		// // sig := <-sigs
		// for sig := range sigs{
		// switch sig {
		// case syscall.SIGINT:
		// 	fmt.Println("Recieved SIGINT (Interrupt)")
		// case syscall.SIGTERM:
		// 	fmt.Println("Recieved SIGTERM (Terminate)")
		// case syscall.SIGHUP:
		// 	fmt.Println("Recieved SIGHUP (Hangup)")
		// case syscall.SIGUSR1:
		// 	fmt.Println("Recieved SIGUSR1 (User-defined signal 1)")
		// 	fmt.Println("User define function executed")
		// 	// continue
		// }
	// }
		// fmt.Println("Graceful exit")
		// os.Exit(0)
	}()
	// simulate some work
	fmt.Println("Working...")
	for {
		time.Sleep(time.Second)
	}
}

// tasklist - List of all processes on Windows
// taskkill /F /PID <PID> : Kill process by PID on Windows(taskkill /F /PID 12345)
// Stop-Processes -Id 12345 -Force