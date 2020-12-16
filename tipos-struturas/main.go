package main

import "fmt"

type age int

func (a *age) Envelhecer() {
	*a++
}

type user struct {
	Name string
	Age  age
	Live bool
}

func main() {
	u := user{
		Name: "Gustavo",
		Age:  19,
		Live: true,
	}

	a := user{
		"gabriel",
		20,
		false,
	}

	a.Age.Envelhecer()
	fmt.Println(u, a)
}
