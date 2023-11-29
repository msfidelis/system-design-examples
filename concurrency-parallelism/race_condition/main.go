package main

import (
	"fmt"
	"sync"
	"time"
)

// Variável compartilhada para contar os alimentos grelhados
var grelhados int = 0

// Função para simular o tempo de preparo de um alimento na churrasqueira
func grelhar() {
	grelhados++
	time.Sleep(time.Millisecond * 100)
}

func main() {

	// Wait Group para esperar todas as goroutines terminarem
	var wg sync.WaitGroup

	// Alimentos disponíveis pra colocar na grelha
	var alimentosChurrasco = 100

	// Simula a concorrência pela grelha
	for i := 0; i < alimentosChurrasco; i++ {
		wg.Add(1) // Adiciona 1 alimento ao contador do WaitGroup
		go func() {
			grelhar()       // Inicia o processo de grelhar o alimento
			defer wg.Done() // Remove um contador do wait group da grelha
		}()
	}

	wg.Wait() // Espera que todas as goroutines chamem Done()

	fmt.Println("Total de itens grelhados na churrasqueira:", grelhados)

}
