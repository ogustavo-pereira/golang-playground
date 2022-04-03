package main

import "fmt"

func main() {
	// Condicional Simples
	
	a := true
	

	if a {
		fmt.Println("IF")
	} else {
		fmt.Println("ELSE")
	}

	b := 2
	
	if b == 2 {
		fmt.Println("IF")
	} else {
		fmt.Println("ELSE")
	}

	if b != 2 {
		fmt.Println("IF")
	} else {
		fmt.Println("ELSE")
	}

	if b > 3 {
		fmt.Println("IF")
	} else {
		fmt.Println("ELSE")
	}

	if b < 3 {
		fmt.Println("IF")
	} else {
		fmt.Println("ELSE")
	}

	if b <= 2 {
		fmt.Println("IF")
	} else {
		fmt.Println("ELSE")
	}

	if b >= 2 {
		fmt.Println("IF")
	} else {
		fmt.Println("ELSE")
	}

	// Condicional Switch
	c := 3

	switch c {
	case 1:
		fmt.Println("Não Entrou")
	case 2:
		fmt.Println("Entrou é igual a 2")
	default:
		fmt.Println("Default")
	}

	if b == 2 && b < 3 {
		fmt.Println("IF")
	} else {
		fmt.Println("ELSE")
	}


	if b == 2 || c == 3 {
		fmt.Println("IF")
	} else {
		fmt.Println("ELSE")
	}


}
