package main

import (
	"fmt"

	"github.com/zoemurmure/calculator"
	"github.com/zoemurmure/geometry"
	"github.com/zoemurmure/store"
	"rsc.io/quote"
)

func main() {
	num := calculator.Sum(1, 2)
	fmt.Println(num)
	fmt.Println("Version: ", calculator.Version)
	fmt.Println(quote.Hello())

	// Challenge 1
	// cannot refer to unexported name calculator.internalSum
	//total := calculator.internalSum(5)
	//fmt.Println(total)

	t := geometry.Triangle{}
	t.SetSize(3)
	fmt.Println("Perimeter:", t.Perimeter())

	var s geometry.Shape = geometry.Square{}
	//s.SetSize(5)
	printInformation(s)

	c := geometry.Circle{}
	c.SetRadius(6)
	printInformation(c)

	e := store.Employee{Account: store.Account{}}
	e.ChangeName("John", "Doe")
	fmt.Println(e)
	fmt.Println("Credit: ", e.CheckCredits())
	if credits, err := e.AddCredits(325); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New Credits Balance = ", credits)
	}
	fmt.Println("Credit: ", e.CheckCredits())

	if credits, err := e.RemoveCredits(443); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New Credits Balance = ", credits)
	}
	fmt.Println("Credit: ", e.CheckCredits())
}

func printInformation(s geometry.Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}
