package main

import "fmt"

func findprimes(num int) bool {
	//if num == 1 {
	//	return false
	//}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	if num > 1 {
		return true
	} else {
		return false
	}
	//return true
}

func main() {
	for i := 1; i <= 20; i++ {
		if findprimes(i) {
			//fmt.Println("Find prime: ", i)
			fmt.Printf("%v ", i)
		}
	}
}
