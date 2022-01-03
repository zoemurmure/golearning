package main

import "fmt"

func main() {
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807
	fmt.Println(integer8, integer16, integer32, integer64)

	rune := 't'
	fmt.Println(rune)

	// Chanllenge 1
	//var integer int = 2147483647
	var integer int = 9223372036854775807
	fmt.Println(integer)

	// Chanllenge 2
	// constant -10 overflows uint
	//var uinteger uint = -10
	//fmt.Println(uinteger)
}
