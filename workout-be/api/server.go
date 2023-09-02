package api

import (
	"crypto/ed25519"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/DMonkey83/MyFitnessApp/workout-be/config"
	db "github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/workout-be/token"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     config.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer function
func NewServer(config config.Config, store db.Store) (*Server, error) {
	privateKey, publicKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}

	// Save the private key to a file (keep this secure)
	privateFile, err := os.Create("tls/private.pem")
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()
	pem.Encode(privateFile, &pem.Block{Type: "PRIVATE KEY", Bytes: privateKey})

	// Save the public key to a file (you can share this)
	publicFile, err := os.Create("tls/public.pem")
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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("goal", validGoal)
		v.RegisterValidation("completion", validCompletion)
		v.RegisterValidation("rating", validRating)
		v.RegisterValidation("equipment", validEquipment)
		v.RegisterValidation("weight_unit", validWeightUnit)
		v.RegisterValidation("difficulty", validDifficulty)
		v.RegisterValidation("fatigue_level", validFatigueLevel)
		v.RegisterValidation("muscle_group", validMuscleGroup)
		v.RegisterValidation("is_public", validVisibility)
	}

	server.SetupRouter()
	return server, nil
}

func (server *Server) Start(address string) error {

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://127.0.0.1:3000"} // Replace with your frontend's URL
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Credentials", "Access-Control-Allow-Origin"}
	config.AllowCredentials = true

	server.router.Use(cors.New(config))
	server.router.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(204)
	})

	return server.router.Run(address)
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
