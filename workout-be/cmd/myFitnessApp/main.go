package main

import (
	"context"
	"log"

	"github.com/DMonkey83/MyFitnessApp/workout-be/config"
	server "github.com/DMonkey83/MyFitnessApp/workout-be/internal/api"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	configuration, err := config.LoadConfig("./")
	connPool, err := pgxpool.New(context.Background(), configuration.DBSource)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	server := server.NewServer(configuration, connPool)

	print("HEllo")

	if err := server.Start(configuration.Port); err != nil {
		panic(err)
	}

}
