package main

import (
	"context"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type PedidoDeCompra struct {
	Id         string
	Item       string
	Quantidade float64
}

// Função mock para exemplificar a chegada de alguma mensagem
func consomeMensagem() PedidoDeCompra {
	return PedidoDeCompra{
		Id:         "12345",
		Item:       "pão de alho",
		Quantidade: 4,
	}
}

// Função mock para exemplificar o processamento de uma mensagem
func processaMensagem(pedido PedidoDeCompra) bool {
	fmt.Println("Processando pedido:", pedido.Id)
	time.Sleep(1 * time.Second)
	return true
}

func main() {

	var ctx = context.Background()

	// Create a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// looping de consumo
	NovoPedido := consomeMensagem()
	mutexKey := NovoPedido.Id

	// Verifica se o Lock já existe
	lock, _ := client.Get(ctx, mutexKey).Result()
	if lock != "" {
		fmt.Println("Mutex travado para o recurso", mutexKey)
		return
	}

	// Criando um lock para o registro por 10 segundos
	err := client.Set(ctx, mutexKey, "locked", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Mutex criado para o recurso por 10s:", mutexKey)

	// Processa o registro
	success := processaMensagem(NovoPedido)
	if !success {
		return
	}

	fmt.Println("Pedido processado:", mutexKey)

	// Libera o Mutex
	_, err = client.Del(ctx, mutexKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Mutex liberado para o recurso:", mutexKey)

}
