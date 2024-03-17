package listeners

import (
	"context"
	"fmt"
	clients "orchestrator/pkg/kafka"
	"sync"
)

func NewCompensationListener() {
	var wg sync.WaitGroup

	consumers := 2
	bootstrapServers := "0.0.0.0:9092"
	topic := "compensation"
	consumerGroup := "orchestrator"

	for i := 0; i < consumers; i++ {
		wg.Add(1)
		var consumerID = i + 1
		consumer := clients.GetConsumer(bootstrapServers, topic, consumerGroup, consumerID, false)
		consumerName := fmt.Sprintf("%v-%v", consumerGroup, consumerID)

		fmt.Printf("[Orchestrator - Compensation] Starting consumer %v\n", consumerName)

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
