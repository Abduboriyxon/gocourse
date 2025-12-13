package main

import (
	"fmt"
	"time"
)

func channel() {

	// variable := make(chan type) '<-' operator
	greeting := make(chan string)
	greetString := "Hello"

	// greeting <- greetString // blocking cuz it is continiously tyring to recieve values, this is incorrect
	go func(){
		greeting <- greetString
		greeting <- "World"
		for _, e := range "abcde"{
			greeting <- "Alphabet: " + string(e)
		}
	}()

	// go func(){
	// 	receiver := <- greeting
	// 	fmt.Println(receiver)
	// 	receiver = <- greeting
	// 	fmt.Println(receiver)
	// }()

	receiver := <- greeting
	fmt.Println(receiver)
	receiver = <- greeting
	fmt.Println(receiver)

	for range 5{
		rcvr := <- greeting
		fmt.Println(rcvr)
	}

	time.Sleep(1*time.Second)
	fmt.Println("End of the program")
}