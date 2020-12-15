package main

import "fmt"

func main() {
	result := soma(10, 20)
	fmt.Println(result)

	result = sub(10, 20)
	fmt.Println(result)

	result = mult(10, 20)
	fmt.Println(result)

	resultDiv := div(10, 20)
	fmt.Println(resultDiv)

	maiorQn, msg := maiorQ(10, 20)
	fmt.Println(maiorQn, msg)
}

func soma(n1, n2 int) int {
	fmt.Println("Iniciando Soma")
	return n1 + n2
}

func sub(n1, n2 int) int {
	fmt.Println("Iniciando Subtracao")
	return n1 - n2
}

func mult(n1, n2 int) int {
	fmt.Println("Iniciando Multiplicação")
	return n1 * n2
}

func div(n1, n2 float64) float64 {
	fmt.Println("Iniciando Divisão")

	if n2 == 0 {
		return -1.00
	}

	return n1 / n2
}

func maiorQ(n1 int, n2 int) (int, string) {
	if n1 > n2 {
		return n1, "n1 maior que n2"
	} else {
		return n2, "n2 maior que n1"
	}
}
