package main

import "fmt"

func main() {
	sl := make([]int, 2, 3)

	sl[0] = 90
	sl[1] = 80

	fmt.Println(sl)

	sl2 := sl[:]

	sl = append(sl, 1, 60)

	fmt.Println(sl2)

	sl2[1] = 200

	fmt.Println(sl, sl2)
}
