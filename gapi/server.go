package gapi

import (
	"crypto/ed25519"
	"encoding/pem"
	"fmt"
	"os"

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

// NewServer function creates a new gRPC server.
func NewServer(config config.Config, store db.Store) (*Server, error) {
	privateKey, publicKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}

	// Save the private key to a file (keep this secure)
	privateFile, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()
	pem.Encode(privateFile, &pem.Block{Type: "PRIVATE KEY", Bytes: privateKey})

	// Save the public key to a file (you can share this)
	publicFile, err := os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()
	pem.Encode(publicFile, &pem.Block{Type: "PUBLIC KEY", Bytes: publicKey})
	tokenMaker, err := token.NewPasetoMaker(ed25519.PublicKey(config.TokenKey))
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
