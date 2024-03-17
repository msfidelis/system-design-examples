package migrations

import (
	"fmt"
	"main/pkg/database"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Driver do banco de dados
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func DatabaseMigration() {

	consumerName := "DatabaseMigration"

	// time.Sleep(10 * time.Second)

	db := database.GetPGX()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Printf("[%s] Erro ao criar a config com o postgres %v:\n", consumerName, err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations", // Certifique-se de que este caminho está correto e acessível
		"postgres", driver)
	if err != nil {
		fmt.Printf("[%s] Erro criar a instância de migração %v:\n", consumerName, err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Printf("[%s] Erro ao aplicar as migrations %v:\n", consumerName, err)
	}

	fmt.Printf("[%s] Migrations aplicadas com sucesso:\n", consumerName)
}
