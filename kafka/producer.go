package main

import (
	"encoding/json"
	"fmt"

	"main/entities"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-faker/faker/v4"
)

func main() {
	topic := "ecommerce_nova_venda"
	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers":      "localhost:29092",
			"acks":                   "1", // Configuração para garantir durabilidade sem degradar muito a performance
			"batch.size":             1,   // Ajuste conforme a carga e teste
			"linger.ms":              0,   // Ajuste conforme a necessidade de latência e throughput
			"queue.buffering.max.ms": 0,   //
			"compression.type":       "lz4",
		})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	deliveryChan := make(chan kafka.Event, 100) // Buffer para evitar bloqueio na produção

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

	for i := 0; i < 100000; i++ {
		message := entities.Transacao{}
		err := faker.FakeData(&message)
		if err != nil {
			fmt.Println(err)
		}
		value, err := json.Marshal(message)
		if err != nil {
			fmt.Println(err)
			return
		}

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(value),
		}, deliveryChan)
	}

	p.Flush(15 * 1000)  // Aguarda até 15 segundos para entregar todas as mensagens
	close(deliveryChan) // Fecha o canal após completar a produção e o flush
}
