package main

import (
	"fmt"
)

// Interface genérica para implementar os métodos de enfileiramento
type Queue []interface{}

// Adiciona um item ao fim da fila
func (q *Queue) Enqueue(item interface{}) {
	*q = append(*q, item)
}

// Remove o primeiro item da fila e o retorna
func (q *Queue) Dequeue() interface{} {
	if len(*q) == 0 {
		return nil
	}
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func main() {
	queue := Queue{}

	// Itens a serem adicionados na Queue
	items := []string{
		"Pizza",
		"Hamburger",
		"Churrasco",
	}

	// Adicionando os itens na ordem da lista
	for _, item := range items {
		fmt.Println("Input:", item)
		queue.Enqueue(item)
	}

	fmt.Println()

	// Removendo os itens em ordem de chegada na lista
	fmt.Println("Output:", queue.Dequeue())
	fmt.Println("Output:", queue.Dequeue())
	fmt.Println("Output:", queue.Dequeue())
}
