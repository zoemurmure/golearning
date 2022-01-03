package store

import (
	"errors"
	"fmt"
)

type Account struct {
	firstName, lastName string
}

func (a *Account) ChangeName(f, l string) {
	a.firstName = f
	a.lastName = l
}

type Employee struct {
	Account
	credit float64
}

func (e Employee) String() string {
	return fmt.Sprintf("Name: %s %s, Credits: %.2f", e.firstName, e.lastName, e.credit)
}

// AddCredits, RemoveCredits, and CheckCredits
func (e *Employee) AddCredits(n float64) (float64, error) {
	if n > 0.0 {
		e.credit += n
		return n, nil
	} else {
		return 0.0, errors.New("Invalid credit value")
	}

}

func (e *Employee) RemoveCredits(n float64) (float64, error) {
	if n > 0.0 {
		if n <= e.credit {
			e.credit -= n
			return n, nil
		} else {
			return 0.0, errors.New("You can't remove more credits than the account has.")
		}
	} else {
		return 0.0, errors.New("You can't remove negative numbers.")
	}
}

func (e *Employee) CheckCredits() float64 {
	return e.credit
}
