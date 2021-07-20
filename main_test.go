package main

import (
	"testing"
	"time"
)

func TestInitialiseAccount(t *testing.T){
	want := time.Now().Day()
	got := newAccount("Initialise_Account_test").openingDate.Day()

	if want != got {
		t.Errorf("Want %v, Got %v", want, got)
	}
}

func TestDeposit(t *testing.T) {
	testCases := []struct {
		desc        string
		deposit     float64
		expected    float64
		accountName string
	}{
		{
			desc:        "deposit 100",
			deposit:     100,
			expected:    100,
			accountName: "test_deposit_1",
		},
		{
			desc:        "deposit 20000",
			deposit:     20000,
			expected:    20000,
			accountName: "test_deposit_2",
		},
		{
			desc:        "deposit 1",
			deposit:     1,
			expected:    1,
			accountName: "test_deposit_3",
		},
	}

	for _, tC := range testCases {
		account := newAccount(tC.accountName)
		account.deposit(tC.deposit)
		if tC.expected != account.getBalance() {
			t.Errorf("expected %v. got %v", tC.expected, account.getBalance())
		}
		if account.name != tC.accountName {
			t.Errorf("expected account name %v. got %v", tC.accountName, account.name)
		}
	}

}

func TestWithdraw(t *testing.T) {
	account := newAccount("test_withdraw")
	account.deposit(1000)
	account.withdraw(200)
	want := float64(800)
	got := account.getBalance()

	if want != got {
		t.Errorf("Want %v. Got %v", want, got)
	}

	err := account.withdraw(1000)

	if err == nil {
		t.Errorf("Expected error. Got %v", err)
	}
}

func TestGetStatement(t *testing.T) {
	account := newAccount("test_withdraw")
	account.deposit(1000)
	account.withdraw(200)
	account.deposit(1000)
	account.deposit(1000)
	account.withdraw(50)

	want := 5
	got := len(account.transactions)
	account.getStatement()

	if want != got {
		t.Errorf("Want %v, Got %v", want, got)
	}
}
