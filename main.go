package main

import (
	"errors"
	"fmt"
	"time"
)


type account struct {
	name string
	openingDate time.Time
	balance float64
	transactions []transaction
}

type transaction struct {
	date time.Time
	amount float64
	balance float64
}

func main() {

	return
}

func newAccount(name string) *account {
	a := account{name: name}
	return &a
}

func (a *account) getBalance() float64 {
	return a.balance
}

func (a *account) deposit(amount float64) {
	t := transaction{
		date: time.Now(),
		amount: amount,
		balance: a.balance + amount,
	}
	a.balance += amount
	a.transactions = append(a.transactions, t)
}

func (a *account) withdraw(amount float64) error {

	if (a.balance - amount) > 0 {
		return errors.New("insuficient funds")
	} 

	t := transaction {
		date: time.Now(),
		amount: -amount,
		balance: a.balance - amount,
	}
	a.balance -= amount
	a.transactions = append(a.transactions, t)

	return nil
}

func (a *account) getStatement() {
	fmt.Printf("Date || Amount || Balance\n")

	for _, v := range a.transactions {
		day := v.date.Day()
		month := v.date.Month()
		year := v.date.Year()
		fmt.Printf("%v/%v/%v || %v || %v\n", day, month, year, v.amount, v.balance)
	}
}