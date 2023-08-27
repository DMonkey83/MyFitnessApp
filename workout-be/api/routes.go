package api

import "github.com/gin-gonic/gin"

func (server *Server) SetupRouter() {
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

	// workout plan
	authRoutes.POST("/workout-plan", server.createPlan)
	authRoutes.GET("/workout-plan/:plan_id/:username", server.getPlan)
	authRoutes.PATCH("/workout-plan", server.updatePlan)

	// available workout plan
	authRoutes.POST("/available-plans", server.createAvailablePlan)
	authRoutes.GET("/available-plans/:plan_id", server.getAvailablePlan)
	authRoutes.PATCH("/available-plan", server.updateAvailablePlan)

	// workouts
	authRoutes.POST("/workouts", server.createWorkout)
	authRoutes.GET("/workouts/:workout_id", server.getWorkout)
	authRoutes.GET("/workouts", server.getWorkout)
	authRoutes.PATCH("/workouts", server.updateWorkout)

	// workouts
	authRoutes.POST("/workout-logs", server.createWorkoutLog)
	authRoutes.GET("/workout-logs/:log_id", server.getWorkoutLog)
	authRoutes.GET("/workout-logs", server.getWorkoutLog)
	authRoutes.PATCH("/workout-logs", server.updateWorkoutLog)

	// exercises
	authRoutes.POST("/exercise-logs", server.createExerciseLog)
	authRoutes.GET("/exercise-logs/:exercise_id", server.getExerciseLog)
	authRoutes.PATCH("/exercise-logs", server.updateExerciseLog)

	// exercises
	authRoutes.POST("/exercises", server.createExercise)
	authRoutes.GET("/exercises/:exercise_id", server.getExercise)
	authRoutes.GET("/exercises", server.getExercise)
	authRoutes.PATCH("/exercises", server.updateExercise)

	// exercises
	authRoutes.POST("/available-exercises", server.createAvailablePlanExercise)
	authRoutes.GET("/available-exercises/:id", server.getAvailablePlanExercise)
	authRoutes.PATCH("/available-exercises", server.updateAvailablePlanExercise)

	// exercises
	authRoutes.POST("/one_off-exercises", server.createOneOffExerciseLog)
	authRoutes.GET("/one-off-exercises/:id", server.getOneOffExerciseLog)
	authRoutes.PATCH("/one-off-exercises", server.updateOneOffExerciseLog)

	// Sets
	authRoutes.POST("/sets", server.createSet)
	authRoutes.GET("/sets/:id", server.getSet)
	authRoutes.GET("/sets", server.getSet)
	authRoutes.PATCH("/sets", server.updateSet)

	// Sets
	authRoutes.POST("/exercise_sets", server.createExerciseSet)
	authRoutes.GET("/exercise_sets/:set_id", server.getExerciseSet)
	authRoutes.PATCH("/exercise_sets", server.updateExerciseSet)

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
