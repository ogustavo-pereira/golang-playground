/*
Author: Gustavo Henrique Pereira
Demonstração do algoritimo de ordenação selection sort
1. Instale o Go https://golang.org/doc/install
2. Dentro da pasta do algoritimo execute 'go run ./codigo.go'
*/

package main

import "fmt"

/*
* Criando o tipo do grafo.
* index número do verticie
* next é array de números de verticies interligados
 */
type Graph struct {
	index int
	next  []int
}

// Troca os valores de pos1 com a de pos2
func tradePosition(graph []Graph, pos1 int, pos2 int) []Graph {
	firstValue := graph[pos1]
	secondValue := graph[pos2]

	graph[pos2] = firstValue
	graph[pos1] = secondValue

	return graph
}

// Ordenation with selection sort
func selectionSort(graph []Graph) []Graph {

	// Imprimindo Grafo
	fmt.Println("Grafo desordenado:", graph)

	i := 0

	// Realiza um loop em todo grafo
	for i < len(graph) {
		j := i + 1

		aux := graph[i]
		auxIndex := i

		/*
		* Realiza um loop para validar se existe um vertice com menos
		* graus
		 */
		for j < len(graph) {
			cur := graph[j]

			if len(aux.next) < len(cur.next) {
				aux = cur
				auxIndex = j
			}

			j++
		}

		// Se Existir um vertice com menos graus que o atual é feito a troca.
		if i != auxIndex {
			graph = tradePosition(graph, i, auxIndex)
		}

		// increment i
		i++
	}

	fmt.Println("Grafo ordenado:", graph)

	return graph
}

func main() {

	// Criando Grafo
	graph := []Graph{
		{
			index: 1,
			next:  []int{3, 4},
		},
		{
			index: 2,
			next:  []int{6, 9},
		},
		{
			index: 3,
			next:  []int{1},
		},
		{
			index: 4,
			next:  []int{1, 5},
		},
		{
			index: 5,
			next:  []int{4, 7, 8},
		},
		{
			index: 6,
			next:  []int{2},
		},
		{
			index: 7,
			next:  []int{5, 10},
		},
		{
			index: 8,
			next:  []int{5, 9, 10},
		},
		{
			index: 9,
			next:  []int{2, 8},
		},
		{
			index: 10,
			next:  []int{7, 8},
		},
	}

	selectionSort(graph)
}
