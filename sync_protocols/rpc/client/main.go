package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	Peso float64
}

func main() {
	client, err := rpc.Dial("tcp", "0.0.0.0:1234")
	if err != nil {
		fmt.Println("Falha ao conectar:", err)
		return
	}
	var reply float64
	args := Args{Peso: 85.00}

	fmt.Println("Iniciando a chamada RPC para o serviço Proteinas.Recomendacao")
	err = client.Call("Proteinas.Recomendacao", args, &reply)
	if err != nil {
		fmt.Println("Erro na chamada:", err)
		return
	}

	fmt.Printf("O consumo de proteínas adequado para o peso de %v kg é de %vg por dia\n", args.Peso, reply)
}
