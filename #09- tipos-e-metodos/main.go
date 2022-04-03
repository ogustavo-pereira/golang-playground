package main

import "fmt"

func main() {
	var l List

	l.init()
	l.show()
}

type List []interface{}

func (l *List) init() {
	*l = []interface{}{
		10,
		"Ola",
		1.9,
		false,
	}
}

func (l List) show() {
	fmt.Println(l...)
}
