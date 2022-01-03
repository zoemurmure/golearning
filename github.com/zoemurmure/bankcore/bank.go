package bank

import (
	"encoding/json"
	"errors"
)

type Bank interface {
	Statement() string
}

type Customer struct {
	Name    string
	Address string
	Phone   string
}

type Account struct {
	Customer
	Number  int32
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}
	if a.Balance < amount {
		return errors.New("the amount to withdraw should be smaller than the account's balance")
	}
	a.Balance -= amount
	return nil
}

func (a *Account) Statement() string {
	//return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
	data, err := json.Marshal(a)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func (a *Account) Transfer(account *Account, amount float64) error {
	if err := a.Withdraw(amount); err != nil {
		return err
	}
	if err := account.Deposit(amount); err != nil {
		return err
	}
	return nil
}

func Statement(b Bank) string {
	return b.Statement()
}
