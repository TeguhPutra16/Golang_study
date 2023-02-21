package main

import (
	"fmt"
	"sync"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})
var group = sync.WaitGroup{}

func waitCondition(value int) {

	group.Add(1)
	defer group.Done()
	cond.L.Lock()
	cond.Wait()
	fmt.Println("Hasil", value)
	cond.L.Unlock()
}

func main() {

	for i := 0; i < 10; i++ {
		go waitCondition(i)

	}

	// go func() {
	// 	for i := 0; i < 10; i++ {

	// 		cond.Signal()
	// 		time.Sleep(1 * time.Second)

	// 	}

	// }()

	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast()

	}()

	group.Wait()

}
