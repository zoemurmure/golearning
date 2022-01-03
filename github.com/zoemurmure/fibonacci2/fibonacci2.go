package main

import (
	"fmt"
	"time"
)

var quitCh = make(chan bool)

func fib(ch chan int) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quitCh:
			fmt.Println("Done calculating Fibonacci!")
			return
		}
	}
}

func main() {
	start := time.Now()

	fiboCh := make(chan int)
	go fib(fiboCh)

	command := ""

	for {
		fmt.Scanf("%s", &command)
		if command == "quit" {
			quitCh <- true
			break
		}
		fmt.Println(<-fiboCh)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
