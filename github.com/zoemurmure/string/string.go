package main

import "fmt"

func main() {
	var firstName = "John"
	lastName := "Doe"
	fmt.Println(firstName, lastName)
	fullName := "John Doe\t(alias \"Foo\")\n"
	fmt.Println(fullName)

	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	fmt.Println(defaultInt, defaultBool, defaultFloat32, defaultFloat64, defaultString)
}
