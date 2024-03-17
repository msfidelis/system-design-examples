package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

var onceDBPGX sync.Once
var pgxInstance *sql.DB

var onceBun sync.Once
var BunInstance *bun.DB

func getDBUrl() string {
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	schema := os.Getenv("DATABASE_DB")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, schema)
}

func GetPGX() *sql.DB {
	onceDBPGX.Do(func() {
		var err error
		config, err := pgx.ParseConfig(getDBUrl())
		if err != nil {
			panic(err)
		}
		config.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

		pgxInstance = stdlib.OpenDB(*config)
		pgxInstance.SetMaxOpenConns(100)
		pgxInstance.SetMaxIdleConns(100)
	})
	return pgxInstance
}

func GetDB() *bun.DB {
	onceBun.Do(func() {
		conn := GetPGX()
		// conn := GetDBConn()
		BunInstance = bun.NewDB(conn, pgdialect.New(), bun.WithDiscardUnknownColumns())
	})
	return BunInstance
}
