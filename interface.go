package main

import "fmt"

type person struct {
	Nama string
}
type animal struct {
	Nama string
}

type Hasname interface {
	GetName() string
}

func SayHello(i Hasname) {
	fmt.Println("hello", i.GetName())
}

func (p person) GetName() string {
	return p.Nama
}

func main() {
	var teguh person
	teguh.Nama = "teguh"
	SayHello(teguh)
}
