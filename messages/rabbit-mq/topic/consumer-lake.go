package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	if err != nil {
		fmt.Println("Falha ao conectar com o broker", err)
		return
	}
	defer conn.Close()

	// Criando um canal
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Falha ao abrir um canal com o broker", err)
		return
	}
	defer ch.Close()

	// Criação de uma Queue de faturamento de vendas
	// de prioridade alta, onde somente os clientes de maior volume
	// financeiro será destinada
	queueLake, err := ch.QueueDeclare(
		"queue.faturamento.datalake", // Nome da fila
		true,                         // durable
		false,                        // delete when unused
		false,                        // exclusive
		false,                        // no-wait
		nil,                          // arguments
	)
	if err != nil {
		fmt.Println("Falha ao criar a queue", err)
		return
	}

	msgs, err := ch.Consume(
		queueLake.Name, // queue
		"data-lake",    // consumer tag
		false,          // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("[DataLake] enviando o pedido %v: %v\n", queueLake.Name, string(d.Body))
			d.Ack(true)
		}
	}()

	fmt.Println("[DataLake] Aguardando por mensagens")

	<-forever
}
