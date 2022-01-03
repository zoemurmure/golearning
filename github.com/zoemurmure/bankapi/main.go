package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/zoemurmure/bank"
)

type CustomAccount struct {
	*bank.Account
}

func (c *CustomAccount) Statement() string {
	data, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

var accounts = map[float64]*CustomAccount{}

func main() {
	accounts[1001] = &CustomAccount{
		Account: &bank.Account{
			Customer: bank.Customer{
				Name:    "John",
				Address: "Los Angeles, California",
				Phone:   "(213) 555 0147",
			},
			Number: 1001,
		},
	}
	accounts[1002] = &CustomAccount{
		Account: &bank.Account{
			Customer: bank.Customer{
				Name:    "Mark",
				Address: "Irvine, California",
				Phone:   "(949) 555 0198",
			},
			Number: 1002,
		},
	}
	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/transfer", transfer)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func statement(w http.ResponseWriter, r *http.Request) {
	numberqs := r.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}
	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			fmt.Fprint(w, bank.Statement(account))
		}
	}
}

func deposit(w http.ResponseWriter, r *http.Request) {
	numberqs := r.URL.Query().Get("number")
	amountqs := r.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}
	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			if err = account.Account.Deposit(amount); err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprint(w, account.Statement())
			}
		}
	}
}

func withdraw(w http.ResponseWriter, r *http.Request) {
	numberqs := r.URL.Query().Get("number")
	amountqs := r.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}
	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			if err = account.Account.Withdraw(amount); err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprint(w, account.Statement())
			}
		}
	}
}

func transfer(w http.ResponseWriter, r *http.Request) {
	fromqs := r.URL.Query().Get("from")
	amountqs := r.URL.Query().Get("amount")
	toqs := r.URL.Query().Get("to")

	if fromqs == "" {
		fmt.Fprintf(w, "Account number is missing")
		return
	}

	if from, err := strconv.ParseFloat(fromqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number")
	} else if to, err := strconv.ParseFloat(toqs, 64); err != nil {
		fmt.Fprintf(w, "Invaid destination account number")
	} else {
		if accountA, ok := accounts[from]; !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", from)
		} else if accountB, ok := accounts[to]; !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", to)
		} else {
			err := accountA.Account.Transfer(accountB.Account, amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprint(w, accountA.Statement())
			}
		}
	}
}
