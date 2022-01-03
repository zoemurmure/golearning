package main

import "fmt"

type Person struct {
	Name, Country string
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s\nCountry: %s\n", p.Name, p.Country)
}

func main() {
	p := Person{Name: "John Doe", Country: "U.S"}
	fmt.Println(p)
}
