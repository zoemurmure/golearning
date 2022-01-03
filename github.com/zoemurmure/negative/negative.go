package main

import "fmt"

func main() {
	var val int = 0
	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d\n", &val)
		switch {
		case val == 0:
			fmt.Println("0 is neither negative nor positive")
		case val > 0:
			fmt.Println("You entered ", val)
		default:
			panic("NEGATIVE!")
		}
	}
}
