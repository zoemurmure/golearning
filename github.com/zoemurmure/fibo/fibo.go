package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		panic("nees an integer parameter")
	}
	num, _ := strconv.Atoi(os.Args[1])

	fmt.Println(fibo(num))
}

func fibo(num int) []int {
	if num < 2 {
		return make([]int, 0)
	}
	//result := []int{1, 1}
	result := make([]int, num)
	result[0], result[1] = 1, 1
	for i := 2; i < num; i++ {
		//result = append(result, result[i-1]+result[i-2])
		result[i] = result[i-1] + result[i-2]
	}

	return result
}
