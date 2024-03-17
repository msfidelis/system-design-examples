package services

import (
	"context"
	"database/sql"
	"fmt"
	"main/entities"
	"main/pkg/database"
)

func CreatePost(p entities.Post) (bool, error) {
	ctx := context.Background()
	db := database.GetDB()

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		fmt.Printf("Erro ao iniciar a transação: %v\n", err)
		return false, err
	}

	_, err = tx.NewInsert().
		Model(&p).
		Exec(ctx)

	err = tx.Commit()
	if err != nil {
		fmt.Printf("Erro ao fazer commit da transação: %v\n", err)
		return false, err
	}

	// Garantir que um rollback será feito em caso de erro
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	return true, nil
}
