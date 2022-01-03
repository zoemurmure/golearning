package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int16 = 127
	var b int32 = 23456
	fmt.Println(int32(a) + b)

	//var s string = "11"
	//var i int = 32
	//fmt.Println(int(s) + i) // cannot convert s (type string) to type int

	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, s)

}
