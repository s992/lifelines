package main

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"path"

	_ "github.com/mattn/go-sqlite3"

	"github.com/s992/logger/internal/config"
	"github.com/s992/logger/internal/generated/db"
	"github.com/s992/logger/internal/server"
)

//go:embed sql/schema.sql
var ddl string

//go:embed dist/client/*
var clientFiles embed.FS

func main() {
	config.InitEnv()

	queries, err := initDb(ddl, config.Env.DbDir)
	if err != nil {
		panic(err)
	}

	if err := server.Run(&server.ServerConfig{
		ClientFiles: clientFiles,
		Port:        config.Env.Port,
		Queries:     queries,
	}); err != nil {
		panic(err)
	}
}

func initDb(ddl string, dir string) (*db.Queries, error) {
	ctx := context.Background()

	dbPath := path.Clean(fmt.Sprintf("%s/logger.sqlite3", dir))
	database, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?_foreign_keys=on", dbPath))
	if err != nil {
		return nil, err
	}

	if _, err := database.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	return db.New(database), nil
}
