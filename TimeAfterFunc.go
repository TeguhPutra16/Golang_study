package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	group := sync.WaitGroup{}

	group.Add(1)
	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})

	fmt.Println(time.Now())
	group.Wait()

}
