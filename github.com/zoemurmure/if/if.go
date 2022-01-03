package main

import "fmt"

func givenumber() int {
	return -1
}

func main() {
	if num := givenumber(); num < 0 {
		fmt.Println("smaller than zero")
	} else if num == 0 {
		fmt.Println("equal to zero")
	} else {
		fmt.Println("bigger than zero")
	}
}
