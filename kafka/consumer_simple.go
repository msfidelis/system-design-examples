package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	// Nome do Tópico que enviaremos os eventos
	topic := "ecommerce_nova_venda"

	// Configuração do Consumidor
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":             "localhost:29092",    // Endereço dos Brokers
		"group.id":                      "faturamento",        // Consumer Group
		"auto.offset.reset":             "earliest",           // Controle do Offset
		"partition.assignment.strategy": "cooperative-sticky", // Partition Assignment
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// Subscreve o consumidor ao tópico
	c.SubscribeTopics([]string{topic}, nil)

	// Looping de consumo
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Mensagem recebida: %s\n", string(msg.Value))

			// Faz o commit do offset manualmente
			_, commitErr := c.CommitMessage(msg)
			if commitErr != nil {
				fmt.Printf("Erro no commit do offset: %v\n", commitErr)
			} else {
				fmt.Println("Offset commitado com sucesso")
			}

		} else {
			fmt.Printf("Erro ao consumir o evento: %v (%v)\n", err, msg)
			break
		}
	}

	c.Close()
}
