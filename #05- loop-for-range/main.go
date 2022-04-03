package main

import "fmt"

func main() {

	// For Convencional
	anr := [...]int{1, 2, 3, 4}

	for i := 0; i < len(anr); i++ {
		fmt.Println(i)
		fmt.Println(anr[i])
	}

	fmt.Println("FOR RANGE")
	arr := [...]int{1, 2, 3, 4}
	for key, value := range arr {
		fmt.Println(key * value)
	}

	user := map[string]string{
		"name": "Eliseu",
		"nick": "Zeus",
	}

	for key, value := range user {
		fmt.Printf("O campo \"%s\" tem o valor igual a \"%s\"\n", key, value)
	}

	languages := []string{
		"Go",
		"JS",
		"Rust",
		"Python",
	}

	for _, value := range languages {
		fmt.Println(value)
	}
}
