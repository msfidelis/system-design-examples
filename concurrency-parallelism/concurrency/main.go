package main

import (
	"fmt"
	"sync"
	"time"
)

type Atividade struct {
	Nome  string
	Tempo int
}

// Função para simular o tempo de preparo de cada atividade do churrasco
func preparar(item string, tempoPreparo int, churrasco chan<- string) {
	fmt.Printf("Preparando %s...\n", item)
	time.Sleep(time.Duration(tempoPreparo) * time.Second)
	churrasco <- item
}

func main() {

	// Canal das atividades que compõe o churrasco
	churrasco := make(chan string)

	// Wait Group para esperar todas as goroutines terminarem
	var wg sync.WaitGroup

	// Lista de Atividades do Churrasco
	tarefas := []Atividade{
		{"picanha", 5},
		{"costela", 7},
		{"linguica", 3},
		{"salada", 2},
		{"bebidas", 1},
		{"churrasqueira", 2},
		{"queijo", 3},
	}

	for _, tarefa := range tarefas {
		wg.Add(1)                                         // Adiciona 1 atividade ao contador do WaitGroup
		go preparar(tarefa.Nome, tarefa.Tempo, churrasco) // Inicia o preparo de uma atividade do churrasco numa goroutine
	}

	go func() {
		wg.Wait()        // Espera que todas as goroutines chamem Done()
		close(churrasco) // Fecha o canal após todas as atividades do churrasco terminarem
		fmt.Println("\nChurrasco terminou :/")
	}()

	for item := range churrasco {
		fmt.Printf("%s foi preparado.\n", item)
		wg.Done() // Remove um contador do wait group
	}

}
