package main

import (
	"context"
	"log"

	"github.com/golang-documented-todo-api/app/cmd"
	"github.com/golang-documented-todo-api/app/datasources"
	"github.com/golang-documented-todo-api/app/datasources/db"
	"github.com/golang-documented-todo-api/app/pkg/env"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := pgxpool.New(ctx, env.Get().Database.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	db := db.NewDatabase(conn)

	log.Fatal(cmd.RunServer(ctx, &datasources.DataSources{DB: db}))
}
