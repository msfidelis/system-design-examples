package main

import (
	"context"
	"fakeproducer/pkg/fake"
	clients "fakeproducer/pkg/kafka"
	"fakeproducer/pkg/tracer"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.opentelemetry.io/otel"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"

	guuid "github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

func main() {

	// Graceful Shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	serviceName := "external-stimulus"
	serviceVersion := "0.1.0"

	tp := tracer.InitTracer(serviceName, serviceVersion)
	tr := otel.Tracer(serviceName)

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	// Kafka Setup
	bootstrapServers := "0.0.0.0:9092"
	topic := "new_transaction"
	batchSize := 0

	// Healthcheck
	go func() {
		port := "8080"
		http.HandleFunc("/healthcheck", healthCheckHandler)
		log.Printf("Server running On: %s", port)
		log.Printf("Enabling healthcheck on /healthcheck")
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}()

	producer := clients.GetProducer(bootstrapServers, topic, batchSize, 1, false, "murmur")

	for {
		id := guuid.New().String()
		message := fake.MockPayload()

		spanCtx, span := tr.Start(ctx, "New travel reservation event")

		span.SetAttributes(
			attribute.String("Reservation", id),
			attribute.String("Event", message),
		)

		err := produceMessage(producer, spanCtx, topic, id, message)
		log.Println("Message produced:", id)

		if err != nil {
			span.SetAttributes(
				attribute.String("error", "true"),
				attribute.String("message", err.Error()),
			)
			log.Println(err.Error())
		}

		span.End()
		time.Sleep(500 * time.Millisecond)

	}

}

// Simple Probe Healthcheck
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// Produce message with OpenTelemetry Context
func produceMessage(producer *kafka.Writer, ctx context.Context, topic string, key string, message string) error {

	// Childspan da produção da mensagem
	tracer := otel.Tracer("kafka-producer")
	_, childSpan := tracer.Start(ctx, "produceMessage")
	defer childSpan.End()

	childSpan.SetAttributes(
		attribute.String("Reservation", key),
		attribute.String("Event", message),
	)

	// Preparar cabeçalhos para propagação do contexto
	carrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	headers := make([]kafka.Header, 0, len(carrier))
	for k, v := range carrier {
		headers = append(headers, kafka.Header{Key: k, Value: []byte(v)})
	}

	msg := kafka.Message{
		Key:     []byte(key),
		Value:   []byte(message),
		Headers: headers,
	}

	err := producer.WriteMessages(ctx, msg)
	if err != nil {
		childSpan.RecordError(err)
		return err
	}
	return nil
}
