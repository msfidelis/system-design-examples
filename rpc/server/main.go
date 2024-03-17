package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Args struct {
	Peso float64
}

// Calculo de Recomendação de Consumo de Proteínas
// Baseado no peso informado
type Proteinas float64

func (p *Proteinas) Recomendacao(args Args, reply *float64) error {
	// Calcula o o consumo de proteina recomendado e devolve para o objeto de respotsa
	*reply = args.Peso * 2
	return nil
}

func main() {

	// Registrando o serviço RPC na porta 1234
	proteina := new(Proteinas)
	rpc.Register(proteina)

	ln, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Falha ao ouvir na porta:", err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Falha ao aceitar:", err)
			continue
		}
		// Servindo a rotina RPC
		go rpc.ServeConn(conn)
	}

}
