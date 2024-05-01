package main

import (
	"fmt"
)

type Stack []interface{}

// Push adiciona um item ao topo da pilha
func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

// Pop remove o item do topo da pilha e o retorna
func (s *Stack) Pop() interface{} {
	if len(*s) == 0 {
		return nil
	}
	index := len(*s) - 1
	item := (*s)[index]
	*s = (*s)[:index]
	return item
}

func main() {
	stack := Stack{}

	// Itens a serem adicionados na Stack
	items := []string{
		"Pizza",
		"Hamburger",
		"Churrasco",
	}

	// Adicionando os itens na pilha
	for _, item := range items {
		fmt.Println("Input:", item)
		stack.Push(item)
	}

	fmt.Println()

	// Removendo os itens da pilha
	fmt.Println("Output:", stack.Pop())
	fmt.Println("Output:", stack.Pop())
	fmt.Println("Output:", stack.Pop())
}
