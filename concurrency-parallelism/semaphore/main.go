package main

import (
	"fmt"
	"sync"
	"time"
)

// Função para simular o tempo de preparo de cada atividade do churrasco
func grelhar(item int, tempoPreparo int) {
	fmt.Printf("Preparando o alimento %v...\n", item)
	time.Sleep(time.Duration(tempoPreparo) * time.Second)
	fmt.Printf("Alimento %v preparado, desocupando a grelha...\n", item)
}

func main() {

	var wg sync.WaitGroup

	// Numero maximo de goroutines / threads / alimentos que podem ser assados
	var capacidadeDaGrelha = 3
	// Alimentos disponíveis pra colocar na grelha
	var alimentosChurrasco = 10

	// Criando canal que tenha o tamanho maximo igual ao tamanho da grelha
	semaforo := make(chan struct{}, capacidadeDaGrelha)

	// Inicia o processo de assar os alimentos disponíveis
	for i := 0; i < alimentosChurrasco; i++ {

		wg.Add(1)              // Adiciona 1 atividade ao contador do WaitGroup
		semaforo <- struct{}{} // Adquire 1 espaço no semáforo adicionando 1 item ao channel

		// Começando a assar os alimentos na churrasqueira
		go func(i int) {
			alimento := i + 1

			preparar(alimento, 2) // Inicia o preparo da comida
			<-semaforo            // Libera um espaço no semaforo quando terminar o preparo
			wg.Done()             // Termina uma ativilidade no contador do WaitGroup

		}(i)
	}

	wg.Wait() // Espera todos os alimentos serem preparados
	fmt.Println("Acabou o churrasco :/")

}
