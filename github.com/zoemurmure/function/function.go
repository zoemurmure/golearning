package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := sum(os.Args[1], os.Args[2])
	sum, mul := calc(os.Args[1], os.Args[2])

	fmt.Println(sum, mul)
}

func sum(num1 string, num2 string) (result int) {
	n1, _ := strconv.Atoi(num1)
	n2, _ := strconv.Atoi(num2)
	result = n1 + n2
	return
}

func calc(num1 string, num2 string) (sum int, mul int) {
	n1, _ := strconv.Atoi(num1)
	n2, _ := strconv.Atoi(num2)
	sum = n1 + n2
	mul = n1 * n2
	return
}
