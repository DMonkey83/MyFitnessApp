package api

import "github.com/gin-gonic/gin"

func (server *Server) SetupRouter() {
	router := gin.New()
	routes := router.Use(loggingMiddleware())
	routes.POST("/users", server.createUser)
	routes.POST("/users/login", server.loginUser)
	routes.POST("/tokens/renew_access", server.renewAccessToken)

	// Login
	authRoutes := router.Group("/api/").Use(authMiddleware(server.tokenMaker), loggingMiddleware())

	// userProfiles
	authRoutes.POST("/userProfiles", server.createUserProfile)
	authRoutes.GET("/userProfiles/:username", server.getUserProfile)
	authRoutes.PATCH("/userProfiles", server.updateUserProfile)

	// workout plan
	authRoutes.POST("/workout-plan", server.createPlan)
	authRoutes.GET("/workout-plan/:plan_id/:username", server.getPlan)
	authRoutes.PATCH("/workout-plan", server.updatePlan)

	// available workout plan
	authRoutes.POST("/available-plans", server.createAvailablePlan)
	authRoutes.GET("/available-plans/:plan_id", server.getAvailablePlan)
	authRoutes.GET("/available-plans", server.listAllAvailablePlans)
	authRoutes.GET("/available-plans-by-creator", server.listAllAvailablePlansByCreator)
	authRoutes.PATCH("/available-plan", server.updateAvailablePlan)

	// workouts
	authRoutes.POST("/workouts", server.createWorkout)
	authRoutes.GET("/workouts/:workout_id", server.getWorkout)
	authRoutes.GET("/workouts", server.listWorktouts)
	authRoutes.PATCH("/workouts", server.updateWorkout)

	// workouts
	authRoutes.POST("/workout-logs", server.createWorkoutLog)
	authRoutes.GET("/workout-logs/:log_id", server.getWorkoutLog)
	authRoutes.GET("/workout-logs", server.listWorktoutLogs)
	authRoutes.PATCH("/workout-logs", server.updateWorkoutLog)

	// exercises
	authRoutes.POST("/exercise-logs", server.createExerciseLog)
	authRoutes.GET("/exercise-logs/:exercise_id", server.getExerciseLog)
	authRoutes.GET("/exercise-logs", server.listAllExerciseLogs)
	authRoutes.PATCH("/exercise-logs", server.updateExerciseLog)

	// exercises
	authRoutes.POST("/exercises", server.createExercise)
	authRoutes.GET("/exercises/:exercise_id", server.getExercise)
	authRoutes.GET("/muscle_group_exercises", server.listMuscleGroupExercises)
	authRoutes.GET("/equipment_exercises", server.listEquipmentExercises)
	authRoutes.GET("/exercises", server.listAllExercises)
	authRoutes.PATCH("/exercises", server.updateExercise)

	// exercises
	authRoutes.POST("/available-exercises", server.createAvailablePlanExercise)
	authRoutes.GET("/available-exercises/:id", server.getAvailablePlanExercise)
	authRoutes.GET("/available-exercises", server.listAllAvailablePlanExercises)
	authRoutes.PATCH("/available-exercises", server.updateAvailablePlanExercise)

	// exercises
	authRoutes.POST("/one_off-exercises", server.createOneOffExerciseLog)
	authRoutes.GET("/one-off-exercises/:id/:username", server.getOneOffExerciseLog)
	authRoutes.GET("/one-off-exercises", server.listOneOffExerciseLogs)
	authRoutes.PATCH("/one-off-exercises", server.updateOneOffExerciseLog)

	// Sets
	authRoutes.POST("/sets", server.createSet)
	authRoutes.GET("/sets/:id", server.getSet)
	authRoutes.GET("/sets", server.getSet)
	authRoutes.PATCH("/sets", server.updateSet)

	// Exercise Log Sets
	authRoutes.POST("/exercise_sets", server.createExerciseSet)
	authRoutes.GET("/exercise_sets/:set_id", server.getExerciseSet)
	authRoutes.GET("/exercise_sets", server.listAllExerciseLogSets)
	authRoutes.PATCH("/exercise_sets", server.updateExerciseSet)

	// weight entry
	authRoutes.POST("/weight", server.createWeightEntry)
	authRoutes.GET("/weight/:weight_id", server.getWeightEntry)
	authRoutes.GET("/weight", server.getWeightEntry)
	authRoutes.PATCH("/weight", server.updateWeightEntry)

	// Max reps goal
	authRoutes.POST("/repsgoal", server.createMaxRepGoal)
	authRoutes.GET("/respgoal/:username/:exercise_id/:goal_id", server.getMaxRepGoal)
	authRoutes.GET("/respgoal", server.listMaxReps)
	authRoutes.PATCH("/repsgoal", server.updateMaxRepGoal)

	// max weight goal
	authRoutes.POST("/weightgoal", server.createMaxWeightGoal)
	authRoutes.GET("/weigghtgoal/:username/:exercise_id/:goal_id", server.getMaxWeightGoal)
	authRoutes.GET("/weigthgoal", server.listMaxWeight)
	authRoutes.PATCH("/weightgoal", server.updateMaxWeightGoal)

	server.router = router
}
