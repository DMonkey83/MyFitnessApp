package api

import (
	"github.com/DMonkey83/MyFitnessApp/workout-be/config"
	db "github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config config.Config
	store  db.Store
	router *gin.Engine
}

// NewServer function
func NewServer(config config.Config, store db.Store) *Server {
	server := &Server{config: config, store: store}
	router := gin.Default()

	router.POST("/userProfiles", server.createUserProfile)
	router.GET("/userProfiles/:id", server.getUserProfile)
	router.GET("/userProfiles", server.getUserProfile)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
