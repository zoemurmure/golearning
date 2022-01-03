package main

import (
	"fmt"
	"time"
)

func processing(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Finish processing"
}

func replicating(ch chan string) {
	time.Sleep(time.Second)
	ch <- "Finish replicating"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go processing(ch1)
	go replicating(ch2)
	for i := 0; i < 2; i++ {
		select {
		case process := <-ch1:
			fmt.Println(process)
		case replicate := <-ch2:
			fmt.Println(replicate)
		}
	}
}
