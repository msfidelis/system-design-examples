package main

import (
	"context"
	"log"
	"orchestrator/listeners"
	"orchestrator/pkg/migrations"
	"orchestrator/pkg/tracer"
	"os"
	"os/signal"
	"sync"
)

func main() {

	// Graceful Shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	serviceName := "orchestrator"
	serviceVersion := "0.1.0"

	tp := tracer.InitTracer(serviceName, serviceVersion)

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	// Run Migrations
	migrations.Run()

	var wg sync.WaitGroup

	// Initial New Transaction Listener
	wg.Add(1)
	go func() {
		defer wg.Done()
		listeners.NewTransactionListener(ctx, serviceName)
	}()

	wg.Wait()

}
