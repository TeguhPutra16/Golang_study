package main

import (
	"fmt"
	"sync"
)

func main() {
	group := &sync.WaitGroup{}
	balance := Account{}

	for i := 0; i < 1000; i++ {
		go func() {
			group.Add(1)
			defer group.Done()

			for j := 0; j < 100; j++ {

				balance.AddBalance(1)
				fmt.Println(balance.GetBalacne())

			}

		}()

	}
	group.Wait()
	res := balance.GetBalacne()
	fmt.Println(res)
}

type Account struct {
	RWmutex sync.RWMutex
	Balance int
}

func (account *Account) AddBalance(amount int) {
	account.RWmutex.Lock()
	account.Balance = account.Balance + amount
	account.RWmutex.Unlock()
}

func (account *Account) GetBalacne() int {
	account.RWmutex.RLock()
	balance := account.Balance
	account.RWmutex.RUnlock()
	return balance
}
