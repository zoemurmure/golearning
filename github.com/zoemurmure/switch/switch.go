package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sec := time.Now().Unix()
	rand.Seed(sec)
	i := rand.Int31n(10)
	switch i {
	case 0:
		fmt.Println("Zero ...")
	case 1:
		fmt.Println("One ...")
	case 2:
		fmt.Println("Two ...")
	default:
		fmt.Println("No match")
	}

	switch time.Now().Weekday().String() {
	case "Satursday", "Sunday":
		fmt.Println("Holiday!!!")
	default:
		fmt.Println("Hard work...")
	}

	r := rand.Float32()
	switch {
	case r > 0.1:
		fmt.Println(">0.1", r)
	default:
		fmt.Println("<=0.1", r)
	}

	switch num := 15; {
	case num < 50:
		fmt.Println("< 50")
		fallthrough
	case num > 100:
		fmt.Println("> 100")
		fallthrough
	case num > 200:
		fmt.Println("> 200")
	}
}
