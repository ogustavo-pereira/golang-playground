package main

import "fmt"

func main() {
	// Array
	tabuada := [10]int{0, 5, 10}
	fmt.Println("ARRAY", tabuada)
	fmt.Println("TAMANHO DO ARRAY", len(tabuada))

	// Map
	user := map[string]string{
		"name":    "Eliseu",
		"Gustavo": "1",
	}

	fmt.Println("MAP", user)

	user["age"] = "19"

	fmt.Println("MAP", user)
	fmt.Println("Idade", user["age"])
	fmt.Println("Tamanho do MAP", len(user))

	//Slice

	slice := []int{}
	slice = append(slice, 90, 11, 50, 30)
	fmt.Println("Slice", slice)
}
