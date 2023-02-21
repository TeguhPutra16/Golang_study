package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	pool := sync.Pool{
		New: func() any {
			return "New" //Default Condition ketika pool kosong
		},
	}

	pool.Put("Teguh")
	pool.Put("Putra")
	pool.Put("Butarbutar")

	for i := 0; i < 10; i++ {

		go func() {

			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)

		}()

	}
	time.Sleep(3 * time.Second)
	fmt.Println("selesai")

}
