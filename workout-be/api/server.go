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

	// Login
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// userProfiles
	authRoutes.POST("/userProfiles", server.createUserProfile)
	authRoutes.GET("/userProfiles/:username", server.getUserProfile)
	authRoutes.GET("/userProfiles", server.getUserProfile)
	authRoutes.PATCH("/userProfiles", server.updateUserProfile)

	// workouts
	authRoutes.POST("/workouts", server.createWorkout)
	authRoutes.GET("/workouts/:workout_id", server.getWorkout)
	authRoutes.GET("/workouts", server.getWorkout)
	authRoutes.PATCH("/workouts", server.updateWorkout)

	// exercises
	authRoutes.POST("/exercises", server.createExercise)
	authRoutes.GET("/exercises/:exercise_id", server.getExercise)
	authRoutes.GET("/exercises", server.getExercise)
	authRoutes.PATCH("/exercises", server.updateExercise)

	// Sets
	authRoutes.POST("/sets", server.createSet)
	authRoutes.GET("/sets/:id", server.getSet)
	authRoutes.GET("/sets", server.getSet)
	authRoutes.PATCH("/sets", server.updateSet)

	// weight entry
	authRoutes.POST("/weight", server.createWeightEntry)
	authRoutes.GET("/weight/:weight_id", server.getWeightEntry)
	authRoutes.GET("/weight", server.getWeightEntry)
	authRoutes.PATCH("/weight", server.updateWeightEntry)

	// Max reps goal
	authRoutes.POST("/repsgoal", server.createMaxRepGoal)
	authRoutes.GET("/respgoal/:username/:exercise_id/:goal_id", server.getMaxRepGoal)
	authRoutes.GET("/respgoal", server.getMaxRepGoal)
	authRoutes.PATCH("/repsgoal", server.updateMaxRepGoal)

	// max weight goal
	authRoutes.POST("/weightgoal", server.createMaxWeightGoal)
	authRoutes.GET("/weigghtgoal/:username/:exercise_id/:goal_id", server.getMaxWeightGoal)
	authRoutes.GET("/weigthgoal", server.getMaxWeightGoal)
	authRoutes.PATCH("/weightgoal", server.updateMaxWeightGoal)

	server.router = router

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
