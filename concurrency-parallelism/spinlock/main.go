package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type SpinLock struct {
	state int32
}

// Método do SpinLock para "travar a grelha"
// Cria um loop ativo (spin) que aguarda o valor do state ter o valor de 1
// Quando entrar na condição de `Unlock()`, libera o runtime para executar outras goroutines
func (s *SpinLock) Lock() {
	for !atomic.CompareAndSwapInt32(&s.state, 0, 1) {
		runtime.Gosched() // Permite que outras goroutines sejam executadas
	}
}

// Método do SpinLock para "destravar a grelha"
// Seta o valor da propriedade `state` para 0
func (s *SpinLock) Unlock() {
	atomic.StoreInt32(&s.state, 0)
}

// Função para simular o tempo de preparo de cada atividade do churrasco
func grelhar(amigo int, lock *SpinLock, wg *sync.WaitGroup) {
	fmt.Printf("Amigo %d está esperando para usar a grelha\n", amigo)
	lock.Lock() // Trava as demais Goroutines para usar a grelha

	fmt.Printf("Amigo %d está grelhando seu almoço\n", amigo)
	time.Sleep(1 * time.Second) // Simulando o tempo para grelhar

	fmt.Printf("Amigo %d terminou de usar a grelhar\n", amigo)
	lock.Unlock() // Destrava a grelha para a proxima goroutine (amigo)

	defer wg.Done() // Decrementa o contador de Goroutines
}

func main() {
	var wg sync.WaitGroup
	var lock SpinLock

	// Define a quantidade de gente no churrasco
	var amigosNoChurrasco = 10

	// Cria go routines para colocar todos os amigos para esperar a grelha liberar
	for i := 1; i <= amigosNoChurrasco; i++ {
		wg.Add(1)                 // Incrementa o contador dos WaitGroups das Goroutines
		go grelhar(i, &lock, &wg) // Inicia a preparação da comida
	}

	wg.Wait()
	fmt.Println("O churrasco terminou :/")

}
