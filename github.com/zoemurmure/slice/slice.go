package main

import "fmt"

func main() {
	slice := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	var num []int
	for i := 0; i < 10; i++ {
		num = append(num, i)
		fmt.Printf("%d\t%d\t%v\n", i, cap(num), num)
	}

	letters := []string{"A", "B", "C", "D", "E"}
	//remove := 2
	//fmt.Println("Before:", letters, "Remove:", letters[remove])
	//letters = append(letters[:remove], letters[remove+1:]...)
	//fmt.Println("After:", letters)
	fmt.Println("Before:", letters)
	slice1 := letters[0:2]
	slice2 := make([]string, 3)
	copy(slice2, letters[1:4])

	slice1[1] = "Z"

	fmt.Println("After:", letters)
	fmt.Println("slice2:", slice2)
}
