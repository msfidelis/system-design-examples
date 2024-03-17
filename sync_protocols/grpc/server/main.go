package main

import (
	"context"
	"fmt"
	imc "main/service/imc"
	"math"
	"net"

	"google.golang.org/grpc"
)

// Cria um servico baseado no protobuf informado
type service struct {
	imc.IMCServiceServer
}

// Método utilizado para calcular o IMC com a altura e peso informados
func (s *service) Calcular(ctx context.Context, in *imc.IMCRequest) (*imc.IMCResponse, error) {
	fmt.Println("Iniciando Calculo")
	result := (in.Weight / (in.Height * in.Height)) * 1
	result = math.Round(result*100) / 100
	return &imc.IMCResponse{Result: float64(result)}, nil
}

func main() {
	// Alocando a porta 50051 para o servidor
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("Falha ao servir na porta 50051:", err)
		return
	}

	// Instancia um servidor gRPC
	s := grpc.NewServer()

	// Registra o serviço de calculo no servidor gRPC
	fmt.Println("Registrando o serviço de Calculo de IMC no server gRPC")
	imc.RegisterIMCServiceServer(s, &service{})

	// Instancia o servidor na porta alocada
	if err := s.Serve(lis); err != nil {
		fmt.Println("Falha criar o servidor gRPC:", err)
		return
	}
}
