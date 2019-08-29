//+build wasm,js

package main

import (
	"fmt"
	"time"
)

func exampleFunc() {
	num := 0
	for {
		fmt.Println("Hello from Web Assembly. Number of times this goroutine has ran: ", num)
		time.Sleep(1 * time.Second)
		num++
	}
}

func main() {
	// This exitChan blocks the program from returning the main function.
	// Without it, your program will exit before the goroutine can complete its execution.
	// This is useful if you want your goroutines to continually run.
	exitChan := make(chan string)
	go exampleFunc()
	<-exitChan
}
