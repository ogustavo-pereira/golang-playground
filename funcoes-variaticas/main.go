package main

import "fmt"

var i = 10

func main() {
	fmt.Println(soma(12, 13, 14, 52))
	fakeLoop()
}

// Funções Variaticas
func soma(numbers ...int) int {
	var total int

	for _, number := range numbers {
		total += number
	}

	return total
}

// Função Recurciva
func fakeLoop() {
	i--
	if i > 0 {
		fakeLoop()
		fmt.Println("LOOP")
	}
}
