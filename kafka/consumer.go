package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"main/entities"
	"main/pkg/database"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	// Configuração do Consumidor
	topic := "ecommerce_nova_venda"
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":             "localhost:29092", // Endereço dos Brokers
		"group.id":                      "faturamento",     // Consumer Group
		"auto.offset.reset":             "earliest",        // Controle do Offset
		"partition.assignment.strategy": "cooperative-sticky",
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// Subscreve o consumidor ao tópico
	c.SubscribeTopics([]string{topic}, nil)

	// Looping de consumo
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Mensagem recebida: %s\n", string(msg.Value))

			// Converte os dados da transação em JSON
			transaction := entities.Transacao{}
			err := json.Unmarshal([]byte(msg.Value), &transaction)
			if err != nil {
				fmt.Printf("Erro ao converter o evento para uma struct: %v\n", err)
				return
			}

			// Instancia o Database
			ctx := context.Background()
			db := database.GetDB()

			// Inicia uma transaction
			tx, err := db.BeginTx(ctx, &sql.TxOptions{})
			if err != nil {
				fmt.Printf("Erro ao iniciar a transação: %v\n", err)
				return
			}

			// Insere o registro da transação no database
			_, err = tx.NewInsert().
				Model(&transaction).
				Exec(ctx)

			// Autorollback
			defer func() {
				if err != nil {
					fmt.Println(err)
					tx.Rollback()
				}
			}()

			// Commita a transação
			err = tx.Commit()
			if err != nil {
				fmt.Printf("Erro ao commitar a transação: %v\n", err)
				return
			}

			// Faz o commit do offset manualmente
			_, commitErr := c.CommitMessage(msg)
			if commitErr != nil {
				fmt.Printf("Erro no commit do offset: %v\n", commitErr)
			} else {
				fmt.Println("Offset commitado com sucesso")
			}

			fmt.Printf("Mensagem processada: %v\n", transaction)
		} else {
			fmt.Printf("Erro ao consumir: %v (%v)\n", err, msg)
			break
		}
	}

	c.Close()
}
