package api

import (
	"fmt"
	"github.com/DMonkey83/MyFitnessApp/workout-be/config"
	db "github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/workout-be/token"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     config.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer function
func NewServer(config config.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/userProfiles", server.createUserProfile)
	authRoutes.GET("/userProfiles/:id", server.getUserProfile)
	authRoutes.GET("/userProfiles", server.getUserProfile)
	authRoutes.PATCH("/userProfiles", server.updateUserProfile)

	server.router = router

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
