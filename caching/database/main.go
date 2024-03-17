package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {

	// Conexão com o Redis.
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Conexão com o Banco de Dados
	mysqlDSN := "usuario:senha@tcp(localhost:3306)/produtos"
	db, err := sql.Open("mysql", mysqlDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ID do pedido que desejamos buscar
	pedidoID := "1"

	// Busca no cache pela chave criada
	valor, err := rdb.Get(ctx, "produto:"+pedidoID).Result()

	// Verifica se o pedido está ou não em cache
	if err == rdb.Nil {
		fmt.Println("Produto não encontrado no cache")

		// Se não estiver em cache, busca no database
		query := `SELECT valor FROM pedidos WHERE id = ?`

		err := db.QueryRow(query, pedidoID).Scan(&valor)
		if err != nil {
			log.Fatal(err)
		}

		// Armazena o resultado no cache Redis para consultas futuras
		err = rdb.Set(ctx, "pedido:"+pedidoID, valor, 0).Err()
		if err != nil {
			log.Fatal(err)
		}

		// Exibe o valor
		fmt.Println("Pedido recuperado do banco de dados e armazenado no cache:", valor)
	} else {
		fmt.Println("Pedido recuperado do cache:", valor)
	}

}
