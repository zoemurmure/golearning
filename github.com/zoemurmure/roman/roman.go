package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("need a parameter")
	}
	romanString := os.Args[1]
	fmt.Println(convert(romanString))
}

func convert(roman string) int {
	result := 0
	prev := 0
	table := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	for i := len(roman) - 1; i >= 0; i-- {
		if curr, exist := table[rune(roman[i])]; exist {
			if curr < prev {
				result -= curr
			} else {
				result += curr
			}
			prev = curr
		} else {
			panic("wrong parameter")
		}

	}

	return result
}
