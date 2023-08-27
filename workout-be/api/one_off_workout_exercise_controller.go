package api

import (
	"log"
	"net/http"

	db "github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createOneOffExerciseLogRequest struct {
	WorkoutID       int64              `json:"workout_id"`
	ExerciseName    string             `json:"exercise_name"`
	Description     string             `json:"description"`
	MuscleGroupName db.Musclegroupenum `json:"muscle_group_name" binding:"required,oneof=Chest Back Legs Shoulders Arms Abs Cardio"`
}

type getOneOffExerciseLogRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updateOneOffExerciseLogRequest struct {
	ID              int32              `json:"id"`
	WorkoutID       int64              `json:"workout_id"`
	Description     string             `json:"description"`
	MuscleGroupName db.Musclegroupenum `json:"muscle_group_name" binding:"required,oneof=Chest Back Legs Shoulders Arms Abs Cardio"`
}

func (server *Server) createOneOffExerciseLog(ctx *gin.Context) {
	var req createOneOffExerciseLogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.CreateOneOffWorkoutExerciseParams{
		WorkoutID:       req.WorkoutID,
		ExerciseName:    req.ExerciseName,
		Description:     req.Description,
		MuscleGroupName: req.MuscleGroupName,
	}
	exercise, err := server.store.CreateOneOffWorkoutExercise(ctx, arg)

	if err != nil {
		errorCode := db.ErrorCode(err)
		if errorCode == db.ForeignKeyViolation || errorCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, exercise)
}

func (server *Server) getOneOffExerciseLog(ctx *gin.Context) {
	var req getOneOffExerciseLogRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	exercise, err := server.store.GetOneOffWorkoutExercise(ctx, int32(req.ID))
	if err != nil {
		if err == db.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	// Check if the profile being updated belongs to the authenticated user
	ctx.JSON(http.StatusOK, exercise)
}

func (server *Server) updateOneOffExerciseLog(ctx *gin.Context) {
	var req updateOneOffExerciseLogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	log.Print(req)
	arg := db.UpdateOneOffWorkoutExerciseParams{
		ID:              req.ID,
		WorkoutID:       req.WorkoutID,
		Description:     req.Description,
		MuscleGroupName: req.MuscleGroupName,
	}

	exercise, err := server.store.UpdateOneOffWorkoutExercise(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, exercise)
}
