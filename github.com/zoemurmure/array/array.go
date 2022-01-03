package main

import "fmt"

func main() {
	var a1 [3]int
	a2 := [10]string{"first", "second", "third"}
	a3 := [...]float64{25: 9.9}

	fmt.Println("a1:", a1, len(a1))
	fmt.Println("a2:", a2, len(a2))
	fmt.Println("a3:", a3, len(a3))
	fmt.Println(a2[1:2])

	a4 := []float64{1, 2, 3}
	//append(a3, 1.2)
	a4 = append(a4, 1.2)
	fmt.Println(a4)
}
