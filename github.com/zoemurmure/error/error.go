package main

import (
	"errors"
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

var ErrNotFound = errors.New("Employee not found!")

func main() {
	employee, err := getInformation(1001)
	if errors.Is(err, ErrNotFound) {
		fmt.Printf("Not Found: %v\n", err)
	} else {
		fmt.Print(employee)
	}
}

func getInformation(id int) (*Employee, error) {
	/*
		employee, err := apiCallEmployee(id)
		if err != nil {
			return nil, fmt.Errorf("Got an error when get the employee information: %v", err)
		} else {
			return employee, nil
		}
	*/
	for tries := 0; tries < 3; tries++ {
		employee, err := apiCallEmployee(id)
		if err == nil {
			return employee, nil
		}
		fmt.Println("Server is not responding, retrying...")
		time.Sleep(time.Second * 2)
	}
	//return nil, fmt.Errorf("server has failed to respond to get the employee information")
	return nil, ErrNotFound
}

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{FirstName: "John", LastName: "Doe"}
	return &employee, nil
}
