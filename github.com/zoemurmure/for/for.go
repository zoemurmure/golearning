package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	rand.Seed(time.Now().Unix())
	var num int64
	//fmt.Println(num)
	for num != 5 {
		num = rand.Int63n(10)
		fmt.Println(num)
	}

	var n int32
	for {
		fmt.Println("In loop...")
		if n = rand.Int31n(10); n == 5 {
			fmt.Println("Finish!")
			break
		}
		fmt.Println(n)
	}
}
