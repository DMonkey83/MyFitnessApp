package gapi

import (
	"fmt"

	"github.com/DMonkey83/MyFitnessApp/config"
	db "github.com/DMonkey83/MyFitnessApp/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/pb"
	"github.com/DMonkey83/MyFitnessApp/token"
)

// Server struct  to serve new Grpc requests
type Server struct {
	pb.UnimplementedMyFitnessAppServer
	config     config.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server.
func NewServer(config config.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
