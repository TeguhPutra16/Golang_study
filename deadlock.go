package main

import (
	"fmt"
	"sync"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Nama    string
	Balance int
}

func (User *UserBalance) Lock() {
	User.Mutex.Lock()

}

func (User *UserBalance) Unlock() {
	User.Mutex.Unlock()

}

func (User *UserBalance) Change(amount int) {
	User.Balance = User.Balance + amount

}

func Transfer(User1 *UserBalance, User2 *UserBalance, amount int) {
	User1.Lock()
	fmt.Println("Lock 1", User1.Nama)
	User1.Change(-amount)
	// User1.Unlock()

	time.Sleep(1 * time.Second)

	User2.Lock()
	fmt.Println("Lock 2", User2.Nama)
	User2.Change(amount)
	// User2.Unlock()

	time.Sleep(1 * time.Second)

	User1.Unlock()
	User2.Unlock()
}

func main() {

	User1 := UserBalance{
		Nama:    "Teguh",
		Balance: 1000000,
	}
	User2 := UserBalance{
		Nama:    "Celin",
		Balance: 1000000,
	}

	go Transfer(&User1, &User2, 200000)
	go Transfer(&User2, &User1, 100000)

	time.Sleep(3 * time.Second)

	fmt.Println(User1, User2)
}
