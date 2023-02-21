package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)

}
