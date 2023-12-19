package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Abstração de um Mecanismo Random
type RandomBalancer struct {
	hosts  []string // Lista de Hosts disponíveis para balanceamento
	mutex  sync.Mutex
	random *rand.Rand // Gerador de números aleatórios
}

// Inicializa um novo balanceador Random
func NewRandomBalancer(hosts []string) *RandomBalancer {
	src := rand.NewSource(time.Now().UnixNano())
	return &RandomBalancer{
		hosts:  hosts,
		random: rand.New(src),
	}
}

// Retorna um host aleatório
func (r *RandomBalancer) getHost() string {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Simplesmente calcula um número aleatório entre 0 e o len(r.hosts)
	randomIndex := r.random.Intn(len(r.hosts))
	return r.hosts[randomIndex]
}

func main() {
	// Lista de hosts disponíveis
	hosts := []string{
		"http://host1.com",
		"http://host2.com",
		"http://host3.com",
	}

	// Inicia o mecanismo Random
	randomBalancer := NewRandomBalancer(hosts)

	// Simula 30 requisições
	for i := 0; i < 30; i++ {
		host := randomBalancer.getHost()
		fmt.Printf("Requisição %d direcionada para: %s\n", i+1, host)
	}
}
