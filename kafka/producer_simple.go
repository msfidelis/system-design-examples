package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	// Nome do Tópico que enviaremos os eventos
	topic := "ecommerce_nova_venda"

	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers":      "localhost:29092",
			"acks":                   "1",   // Configuração para garantir durabilidade sem degradar muito a performance
			"batch.size":             1,     // Tamanho do Batch Size - No caso enviaremos 1 item por vez
			"linger.ms":              0,     // Ajuste conforme a necessidade de latência e throughput
			"queue.buffering.max.ms": 0,     // Enfileiramento de Mensagem
			"compression.type":       "lz4", // Tipo de Compressão da mensagem
		})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	// Buffer para evitar bloqueio na produção
	deliveryChan := make(chan kafka.Event, 100)
	go func() {
		for e := range deliveryChan {
			m := e.(*kafka.Message)
			if m.TopicPartition.Error != nil {
				fmt.Printf("Falha na entrega: %v\n", m.TopicPartition.Error)
			} else {
				fmt.Printf("Mensagem entregue em %v\n", m.TopicPartition)
			}
		}
	}()

	// Looping para envio dos Eventos de forma contínua
	for i := 0; i < 100000; i++ {
		message := "Exemplo de um evento"

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny, // Seleção aleatória da partição
			},
			Value: []byte(message), // Payload com o conteúdo do evento
		}, deliveryChan)
	}

	p.Flush(15 * 1000)  // Aguarda até 15 segundos para entregar todas as mensagens
	close(deliveryChan) // Fecha o canal após completar a produção e o flush
}
