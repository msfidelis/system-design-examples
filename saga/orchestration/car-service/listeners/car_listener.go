package listeners

import (
	"context"
	"fmt"
	"sync"

	clients "orchestrator/pkg/kafka"
)

func NewCarListener() {
	var wg sync.WaitGroup

	consumers := 2
	bootstrapServers := "0.0.0.0:9092"
	topic := "car"
	consumerGroup := "car-service"

	for i := 0; i < consumers; i++ {
		wg.Add(1)
		var consumerID = i + 1
		consumer := clients.GetConsumer(bootstrapServers, topic, consumerGroup, consumerID, false)
		consumerName := fmt.Sprintf("%v-%v", consumerGroup, consumerID)

		fmt.Printf("[Car Service] Rent a Car %v\n", consumerName)

		go func() {
			for {
				m, err := consumer.ReadMessage(context.Background())
				if err != nil {
					wg.Done()
					break
				}
				fmt.Printf("[Key] %s | [Value] %s\n\n\n", m.Key, m.Value)
			}
			wg.Done()
		}()

	}
	wg.Wait()
}
