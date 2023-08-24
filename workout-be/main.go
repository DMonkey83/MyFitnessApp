package main

import (
	"context"
	"log"

	"github.com/DMonkey83/MyFitnessApp/workout-be/api"
	"github.com/DMonkey83/MyFitnessApp/workout-be/config"
	"github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(connPool)
	server := api.NewServer(config, store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("error", err)
	}
}
