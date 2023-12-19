package main

import (
	"fmt"
	"sync"
	"time"
)

// Abstração de um Mecanismo Least Request
type LeastRequest struct {
	hosts    []string   // Lista de Hosts disponíveis para balanceamento
	requests []int      // Contagem de requisições ativas para cada host
	mutex    sync.Mutex // Mutex para operações thread-safe
}

// Inicializa um novo balanceador Least Request
func NewLeastRequest(hosts []string) *LeastRequest {
	return &LeastRequest{
		hosts:    hosts,
		requests: make([]int, len(hosts)),
	}
}

// Retorna o host com o menor número de requisições ativas
func (lr *LeastRequest) getHost() string {
	lr.mutex.Lock()
	defer lr.mutex.Unlock()

	minIndex := 0
	minRequests := lr.requests[0]

	// Encontra o host com o menor número de requisições ativas
	for i, reqs := range lr.requests {
		if reqs < minRequests {
			minIndex = i
			minRequests = reqs
		}
	}

	// Incrementa a contagem de requisições para o host selecionado
	lr.requests[minIndex]++

	return lr.hosts[minIndex]
}

func main() {
	// Lista de hosts disponíveis
	hosts := []string{
		"http://host1.com",
		"http://host2.com",
		"http://host3.com",
	}

	// Inicia o mecanismo Least Request
	leastRequest := NewLeastRequest(hosts)

	// Simula 30 Requests
	for i := 0; i < 30; i++ {
		host := leastRequest.getHost()
		fmt.Printf("Requisição %d direcionada para: %s\n", i+1, host)
	}

	// Simula um pequeno delay para permitir que as goroutines terminem
	time.Sleep(5 * time.Second)

	fmt.Println("Distribuição de requisições executadas:", leastRequest.requests)
}
