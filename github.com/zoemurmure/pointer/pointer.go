package main

import "fmt"

func main() {
	name := "John"
	updateName(&name)
	fmt.Println(name)
}

func updateName(name *string) {
	*name = "David"
}
