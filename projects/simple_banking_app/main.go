package main

import (
	"fmt"
	"sync"
	"time"
)

/*
open account : initial deposit >0
deposit: updating bank balance, you have to make sure that the race condition is checked
close : make sure that no transactions are happening while the account is getting closed
*/

type Account struct {
	balance int64
	status  bool
	sync.Mutex
}

/* modifying the data
f.Lock()
f+=5
f.Unlock()
*/

func Open(amount int64) (*Account, error) {
	if amount < 0 {
		fmt.Println(fmt.Errorf("Amount %v is less than 0", amount))
		return nil, fmt.Errorf("Amount %v is less than 0", amount)
	}
	acc := &Account{balance: amount, status: true}
	fmt.Println("account opened ", acc)
	return acc, nil
}

func (a *Account) Balance() (int64, error) {
	// if the account is active
	if a.status {
		fmt.Println("Balance: ", a.balance)
		return a.balance, nil
	}
	fmt.Println(fmt.Errorf("account is inactive"))
	return 0, fmt.Errorf("account is inactive")
}

func (a *Account) Deposit(amount int64) (int64, error) {
	if a.status {
		// amount less than 0 means withdrawal
		a.Lock()
		defer a.Unlock()

		if amount < 0 {
			bal, err := a.Balance()
			if err != nil {
				fmt.Println(err)
				return 0, err
			}
			if bal+amount < 0 {
				fmt.Println(fmt.Errorf("not enough balance in the account"))

				return 0, fmt.Errorf("not enough balance in the account")
			}
		}
		// amount > 0 => deposit
		a.balance += amount
	} else {
		fmt.Println(fmt.Errorf("account inactive"))
		return 0, fmt.Errorf("account inactive")
	}

	fmt.Println("deposited: ", a.balance)
	return a.balance, nil
}

func (a *Account) Close() (int64, error) {
	a.Lock()

	defer a.Unlock()

	if a.status {
		a.status = false
		fmt.Println("account closed")
		return a.balance, nil
	}

	return 0, fmt.Errorf("the account is already closed")
}

func main() {
	a, _ := Open(1000)

	defer a.Close()
	go a.Deposit(3000)
	go a.Balance()
	go a.Deposit(-100)
	go a.Balance()

	time.Sleep(5 * time.Second)
}
