package main

import "fmt"

func highlow(high int, low int) {
	if high < low {
		fmt.Println("Panic!")
		panic("function highlow, high < low")
	}
	defer fmt.Println("Defered! highlow(", high, ", ", low, ")")
	fmt.Println("Call! highlow(", high, ", ", low, ")")

	highlow(high, low+1)
}

func main() {
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("main(): recover", handler)
		}
	}()
	highlow(2, 0)
	fmt.Println("Finish!")
}
