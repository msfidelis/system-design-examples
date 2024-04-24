package main

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

// Exemplo de implementação de uma default exchange
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

	// Criação da Exchange
	err = ch.ExchangeDeclare(
		"ecommerce.nova.venda.faturamento", // Nome da exchange
		"topic",                            // Tipo da exchangem - topic
		true,                               // durable
		false,                              // auto-deleted
		false,                              // internal
		false,                              // no-wait
		nil,                                // arguments
	)
	if err != nil {
		fmt.Println("Falha ao construir a exchange", err)
		return
	}

	// Criação de uma Queue de faturamento de vendas
	// de prioridade default, onde em teoria a maior parte das
	// mensagens será enviada
	queueDefault, err := ch.QueueDeclare(
		"queue.faturamento", // Nome da fila
		true,                // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		fmt.Println("Falha ao criar a queue", err)
		return
	}

	// Associando a Queue até a Exchange
	// e informando a binding key para roteamento
	err = ch.QueueBind(
		queueDefault.Name,                  // Nome da fila
		"faturamento.prioridade.default",   // Binding key de roteamento - chave de prioridade default
		"ecommerce.nova.venda.faturamento", // Nome da exchange
		false,                              // no-wait
		nil,                                // arguments
	)
	if err != nil {
		fmt.Println("Falha associar a queue a exchange", err)
		return
	}

	// Criação de uma Queue de faturamento de vendas
	// de prioridade alta, onde somente os clientes de maior volume
	// financeiro será destinada
	queuePrioridade, err := ch.QueueDeclare(
		"queue.faturamento.prioritario", // Nome da fila
		true,                            // durable
		false,                           // delete when unused
		false,                           // exclusive
		false,                           // no-wait
		nil,                             // arguments
	)
	if err != nil {
		fmt.Println("Falha ao criar a queue", err)
		return
	}

	// Associando a Queue até a Exchange
	// e informando a binding key para roteamento
	err = ch.QueueBind(
		queuePrioridade.Name,               // Nome da fila
		"faturamento.prioridade.alta",      // Binding key de roteamento - chave de prioridade alta
		"ecommerce.nova.venda.faturamento", // Nome da exchange
		false,                              // no-wait
		nil,                                // arguments
	)
	if err != nil {
		fmt.Println("Falha associar a queue a exchange", err)
		return
	}

	// Criação de uma Queue que receberá todas as mensagens, independente da prioridade
	// A intenção é receber todos os pedidos de faturamento e enviar para um
	// suposto analitico
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

	// Associando a queue na exchange, todas as mensagems que forem enviadas com o pattern
	// faturamento.prioridade.* será enviada para essa fila independente da prioridade informada
	err = ch.QueueBind(
		queueLake.Name,                     // Nome da fila
		"faturamento.prioridade.*",         // Binding key de roteamento - chave de prioridade alta
		"ecommerce.nova.venda.faturamento", // Nome da exchange
		false,                              // no-wait
		nil,                                // arguments
	)
	if err != nil {
		fmt.Println("Falha associar a queue a exchange", err)
		return
	}

	for i := 0; i < 3000000000; i++ {

		routingKey := "faturamento.prioridade.default"
		if rand.Float64() < 0.1 { // mock para dar 10% de chance de uma mensagem ser encaminhada para a queue prioritária
			routingKey = "faturamento.prioridade.alta"
		}

		id := uuid.New()
		// Mensagem simples
		body := fmt.Sprintf("id:%v:%v", routingKey, id)

		// Publicando a mensagem na exchange usando a routing key de default/prioritario
		err = ch.Publish(
			"ecommerce.nova.venda.faturamento", // exchange
			routingKey,                         // routing key (binding key)
			false,                              // mandatory
			false,                              // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		if err != nil {
			fmt.Println("Falha ao publicar a mensagem na exchange", err)
		}
		fmt.Printf("Mensagem de faturamento enviada para a queue %v: %v\n", routingKey, body)
	}

}
