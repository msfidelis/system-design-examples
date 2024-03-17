package main

import (
	"context"
	"log"
	"time"

	imc "main/service/imc"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Falha ao conectar ao servidor gRPC: %v", err)
	}
	defer conn.Close()
	client := imc.NewIMCServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Executa uma chamada gRPC para o servidor calcular o IMC
	peso := 90.5
	altura := 1.77
	r, err := client.Calcular(ctx, &imc.IMCRequest{
		Weight: peso,
		Height: altura,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("O IMC de uma pessoa com %v de peso e %v de altura é de: %v", peso, altura, r.GetResult())
}
