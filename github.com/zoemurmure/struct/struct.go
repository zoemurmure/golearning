package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string
	Address   string `json:"address,omitempty"`
}

type Employee struct {
	// Information Person
	Person
	ManagerID int
}

type Contractor struct {
	Person
	CompanyID int
}

func main() {
	employees := []Employee{
		{
			Person:    Person{FirstName: "larry", LastName: "jin"},
			ManagerID: 0,
		},
		{
			Person:    Person{ID: 0, FirstName: "bob", LastName: "doe", Address: ""},
			ManagerID: 0,
		},
	}
	//fmt.Println(employees)

	data, _ := json.Marshal(employees)
	fmt.Printf("%s\n", data)

	var decoded []Employee
	json.Unmarshal(data, &decoded)
	fmt.Printf("%v\n", decoded)
}
