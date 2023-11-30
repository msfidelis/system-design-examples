package main

import (
	"fmt"
	"time"

	"github.com/go-zookeeper/zk"
)

type PedidoDeCompra struct {
	Id         string
	Item       string
	Quantidade float64
}

// Função mock para exemplificar a chegada de alguma mensagem
func consomeMensagem() PedidoDeCompra {
	return PedidoDeCompra{
		Id:         "123456",
		Item:       "pão de alho",
		Quantidade: 4,
	}
}

// Função mock para exemplificar o processamento de uma mensagem
func processaMensagem(pedido PedidoDeCompra) bool {
	fmt.Println("Processando pedido:", pedido.Id)
	time.Sleep(1 * time.Second)
	return true
}

func main() {

	// Conecta ao ZooKeeper
	conn, _, err := zk.Connect([]string{"0.0.0.0"}, 1000*time.Second)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// looping de consumo
	NovoPedido := consomeMensagem()
	mutexKey := fmt.Sprintf("/%v", NovoPedido.Id)

	// Verifica se o Znode de lock já existe
	exists, _, err := conn.Exists(mutexKey)
	if err != nil || exists == true {
		fmt.Println("Mutex travado para o recurso", mutexKey)
		return
	}

	// Criando um lock para o registro
	acl := zk.WorldACL(zk.PermAll) // Permissões abertas, ajuste conforme necessário
	path, err := conn.Create(mutexKey, []byte{}, zk.FlagEphemeral, acl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Mutex criado para o recurso", mutexKey)

	// Processa o registro
	success := processaMensagem(NovoPedido)
	if !success {
		return
	}

	fmt.Println("Pedido processado:", mutexKey)

	// Libera o Mutex manualmente
	conn.Delete(path, -1)
	fmt.Println("Mutex liberado para o recurso:", mutexKey)

	// Caso a sessão com o zookeeper acabe, todos os locks gerados pela conexão serão liberados.
	time.Sleep(50 * time.Second)
}
