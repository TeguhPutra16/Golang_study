package main

import (
	"fmt"
	"sync"
)

func main() {
	group := &sync.WaitGroup{}
	data := &sync.Map{}

	for i := 0; i < 100; i++ {

		go AddtoMap(data, i, group)

	}

	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})

}

func AddtoMap(data *sync.Map, value int, group *sync.WaitGroup) {

	defer group.Done()
	group.Add(1)
	data.Store(value, value)

}
