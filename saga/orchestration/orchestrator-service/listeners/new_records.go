package listeners

import (
	"context"
	"fmt"
	clients "orchestrator/pkg/kafka"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
)

func NewTransactionListener(ctx context.Context, serviceName string) {
	var wg sync.WaitGroup

	consumers := 1
	bootstrapServers := "0.0.0.0:9092"
	topic := "new_transaction"
	consumerGroup := "orchestrator"

	tr := otel.Tracer(serviceName)

	for i := 0; i < consumers; i++ {
		wg.Add(1)
		var consumerID = i + 1
		consumer := clients.GetConsumer(bootstrapServers, topic, consumerGroup, consumerID, false)
		consumerName := fmt.Sprintf("%v-%v", consumerGroup, consumerID)

		fmt.Printf("[Orchestrator - New Transaction] Starting consumer %v\n", consumerName)

		go func() {
			for {
				m, err := consumer.ReadMessage(ctx)
				if err != nil {
					fmt.Println(err)
					// break
				}
				// Initial
				startTime := time.Now()

				// Extrair o contexto do trace dos cabeçalhos da mensagem Kafka
				carrier := propagation.MapCarrier{}
				for _, h := range m.Headers {
					carrier.Set(string(h.Key), string(h.Value))
				}
				ctxEvent := otel.GetTextMapPropagator().Extract(ctx, carrier)
				_, span := tr.Start(ctxEvent, "Initial Orchestration")

				span.SetAttributes(
					attribute.String("Reservation", string(m.Key)),
					attribute.String("Event", string(m.Value)),
				)

				fmt.Printf("[Key] %s | [Value] %s\n\n\n", m.Key, m.Value)

				time.Sleep(time.Millisecond * 200)

				elapsedTime := time.Since(startTime)
				span.SetAttributes(
					attribute.String("Process time", string(elapsedTime)),
				)

				span.End()

			}
			wg.Done()
		}()

	}
	wg.Wait()
}
