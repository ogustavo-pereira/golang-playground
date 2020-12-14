package main

import "fmt"

func main() {
	var address string
	var number int
	var valid bool

	var (
		name string
		age int
		old bool
	)

	var cep = "\n35500-000"
	street := "\nRua Sao lorenzo"

	address = "\nBelo Horizonte"
	name = "\nJosé"

	const constante = "\nNão Muda isso"
	const (
		constante2 = "\nIsso não muda"
	)

	numero1, numero2 := 10, 20

	fmt.Print(numero1 * numero2)
	fmt.Print(numero1 + numero2)


	fmt.Print(address, number, valid, name, age, old, cep, street, constante, constante2)
}