package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Atividade struct {
	Nome        string
	Tempo       int
	Responsavel int
}

// Função para simular o tempo de preparo de cada atividade do churrasco - Agor aceita uma lista de atividades
func preparar(atividades []Atividade, churrasco chan<- Atividade, amigo int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, atividade := range atividades {
		atividade.Responsavel = amigo // Identifica um responsável pela tarefa
		fmt.Printf("Amigo %v começou a prepação de %s...\n", amigo, atividade.Nome)
		time.Sleep(time.Duration(atividade.Tempo) * time.Second)
		churrasco <- atividade
	}
}

func main() {

	// Canal das atividades que compõe o churrasco
	churrasco := make(chan Atividade)

	// Wait Group para esperar todas as goroutines terminarem
	var wg sync.WaitGroup

	// Recuperando o número de CPU's (pessoas) disponíveis para ajudar no churrasco
	numCPU := runtime.NumCPU()

	fmt.Printf("Número de CPU's (amigos) pra ajudar no churrasco: %v.\n", numCPU)

	// Lista de Atividades do Churrasco - Atividade / Tempo de Execução / Amigo Responsável
	tarefas := []Atividade{
		{"picanha", 5, 0},
		{"costela", 7, 0},
		{"linguica", 3, 0},
		{"salada", 2, 0},
		{"gelar cerveja", 1, 0},
		{"organizar geladeira", 1, 0},
		{"queijo", 3, 0},
		{"caipirinha", 2, 0},
		{"panceta", 4, 0},
		{"espetinhos", 3, 0},
		{"abacaxi", 3, 0},
		{"limpar piscina", 1, 0},
		{"molhos", 2, 0},
		{"pão de alho", 4, 0},
		{"arroz", 4, 0},
		{"farofa", 4, 0},
	}

	fmt.Printf("Número tarefas do churrasco: %v.\n", len(tarefas))

	// Dividindo lista de tarefas do churrasco entre os CPU's (amigos) disponíveis
	// Efetuando o balanceamento arrendondando a divisão sempre pra cima para evitar
	// que alguém CPU (amigo) fique sem fazer nada :)
	sliceSize := (len(tarefas) + numCPU - 1) / numCPU
	fmt.Printf("Número tarefas pra cada CPU (amigo): %v.\n", sliceSize)

	// Dividindo as tarefas entre as CPUs (amigos)
	for i := 0; i < len(tarefas); i += sliceSize {
		end := i + sliceSize
		amigo := end / 2
		if end > len(tarefas) {
			end = len(tarefas)
		}
		wg.Add(1)
		go preparar(tarefas[i:end], churrasco, amigo, &wg)
	}

	go func() {
		wg.Wait()        // Espera que todas as goroutines chamem Done()
		close(churrasco) // Fecha o canal após todas as atividades do churrasco terminarem
	}()

	for atividade := range churrasco {
		fmt.Printf("Amigo %v terminou de preparar %s...\n", atividade.Responsavel, atividade.Nome)
	}

}
