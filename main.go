package main

import (
	"context"
	"database/sql"
	_ "embed"

	_ "github.com/mattn/go-sqlite3"

	"github.com/s992/logger/internal/config"
	"github.com/s992/logger/internal/generated/db"
	"github.com/s992/logger/internal/server"
)

//go:embed sql/schema.sql
var ddl string

func main() {
	config.InitEnv()

	queries, err := initDb(ddl)
	if err != nil {
		panic(err)
	}

	if err := server.Run(&server.ServerConfig{
		Port:    config.Env.Port,
		Queries: queries,
	}); err != nil {
		panic(err)
	}
}

func initDb(ddl string) (*db.Queries, error) {
	ctx := context.Background()

	database, err := sql.Open("sqlite3", "./db.sqlite3?_foreign_keys=on")
	if err != nil {
		return nil, err
	}

	if _, err := database.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	return db.New(database), nil
}
