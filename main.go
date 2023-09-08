package main

import (
	"context"
	"net"
	"net/http"
	"os"

	"github.com/DMonkey83/MyFitnessApp/api"
	"github.com/DMonkey83/MyFitnessApp/config"
	db "github.com/DMonkey83/MyFitnessApp/db/sqlc"
	_ "github.com/DMonkey83/MyFitnessApp/doc/statik"
	"github.com/DMonkey83/MyFitnessApp/gapi"
	"github.com/DMonkey83/MyFitnessApp/pb"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	runDBMigration(config.MigrationURL, config.DBSource)
	store := db.NewStore(connPool)
	go runGatewayServer(config, store)
	runGrpcServer(config, store)
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}

func runGrpcServer(config config.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server:")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMyFitnessAppServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}
	log.Printf("start GPRC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}
}

func runGatewayServer(config config.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server:")
	}

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterMyFitnessAppHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register handler server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create statik fs")
	}

	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))

	mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}
	log.Printf("start HTTP server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start HTTP gateway server")
	}
}

func runGinServer(config config.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server:")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("error")
	}
}
