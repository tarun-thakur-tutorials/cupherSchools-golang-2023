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

func (a *Account) Balance(wg *sync.WaitGroup) (int64, error) {
	defer wg.Done()
	// if the account is active
	if a.status {
		fmt.Println("Balance: ", a.balance)
		return a.balance, nil
	}
	fmt.Println(fmt.Errorf("account is inactive"))
	return 0, fmt.Errorf("account is inactive")
}

func (a *Account) Deposit(amount int64, wg *sync.WaitGroup) (int64, error) {
	if a.status {
		// amount less than 0 means withdrawal
		a.Lock()
		defer a.Unlock()

		if amount < 0 {
			wg.Add(1)
			bal, err := a.Balance(wg)
			if err != nil {
				fmt.Println(err)
				wg.Done()
				return 0, err
			}
			if bal+amount < 0 {
				fmt.Println(fmt.Errorf("not enough balance in the account"))
				wg.Done()
				return 0, fmt.Errorf("not enough balance in the account")
			}
		}
		// amount > 0 => deposit
		a.balance += amount
	} else {
		fmt.Println(fmt.Errorf("account inactive"))
		wg.Done()
		return 0, fmt.Errorf("account inactive")
	}

	fmt.Println("deposited: ", a.balance)
	wg.Done()
	return a.balance, nil
}

func (a *Account) Close(wg *sync.WaitGroup) (int64, error) {
	a.Lock()

	defer a.Unlock()

	if a.status {
		a.status = false
		fmt.Println("account closed")
		wg.Done() // marks the completion of goroutine
		return a.balance, nil
	}
	wg.Done()
	return 0, fmt.Errorf("the account is already closed")
}

func main() {
	var wg sync.WaitGroup

	a, _ := Open(1000)

	wg.Add(1)
	go a.Deposit(3000, &wg)

	wg.Add(1)
	go a.Balance(&wg)

	wg.Add(1)
	go a.Deposit(-100, &wg)

	wg.Add(1)
	go a.Balance(&wg)

	time.Sleep(2 * time.Second)
	wg.Add(1)
	a.Close(&wg)

	wg.Wait()
}
