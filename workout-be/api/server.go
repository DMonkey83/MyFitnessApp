package api

import (
	"fmt"

	"github.com/DMonkey83/MyFitnessApp/workout-be/config"
	db "github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/workout-be/token"
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
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
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
	return server.router.Run(address)
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
