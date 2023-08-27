package api

import (
	"errors"
	"net/http"
	"time"

	db "github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/workout-be/token"
	"github.com/gin-gonic/gin"
)

type createWorkoutRequest struct {
	Username            string          `json:"username"`
	WorkoutDate         time.Time       `json:"workout_date"`
	WorkoutDuration     string          `json:"workout_duration"`
	Notes               string          `json:"notes"`
	FatigueLevel        db.Fatiguelevel `json:"fatigue_level"`
	TotalCaloriesBurned int32           `json:"total_calories_burned"`
	TotalDistance       int32           `json:"total_distance"`
	TotalRepetitions    int32           `json:"total_repetitions"`
	TotalSets           int32           `json:"total_sets"`
	TotalWeightLifted   int32           `json:"total_weight_lifted"`
}

type getWorkoutRequest struct {
	WorkoutID int64 `uri:"workout_id" binding:"required,min=1"`
}

type updateWorkoutRequest struct {
	WorkoutID           int64           `json:"workout_id"`
	WorkoutDate         time.Time       `json:"workout_date"`
	WorkoutDuration     string          `json:"workout_duration"`
	Notes               string          `json:"notes"`
	FatigueLevel        db.Fatiguelevel `json:"fatigue_level"`
	TotalSets           int32           `json:"total_sets"`
	TotalDistance       int32           `json:"total_distance"`
	TotalRepetitions    int32           `json:"total_repetitions"`
	TotalWeightLifted   int32           `json:"total_weight_lifted"`
	TotalCaloriesBurned int32           `json:"total_calories_burned"`
}

func (server *Server) createWorkout(ctx *gin.Context) {
	var req createWorkoutRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateWorkoutParams{
		Username:        authPayload.Username,
		WorkoutDate:     req.WorkoutDate,
		WorkoutDuration: req.WorkoutDuration,
		Notes:           req.Notes,
	}
	workout, err := server.store.CreateWorkout(ctx, arg)

	if err != nil {
		errorCode := db.ErrorCode(err)
		if errorCode == db.ForeignKeyViolation || errorCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, workout)
}

func (server *Server) getWorkout(ctx *gin.Context) {
	var req getWorkoutRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	workout, err := server.store.GetWorkout(ctx, req.WorkoutID)
	if err != nil {
		if err == db.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	// Check if the profile being updated belongs to the authenticated user
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if workout.Username != authPayload.Username {
		err := errors.New("profile doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(err))
	}
	ctx.JSON(http.StatusOK, workout)
}

func (server *Server) updateWorkout(ctx *gin.Context) {
	var req updateWorkoutRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.UpdateWorkoutParams{
		WorkoutID:           req.WorkoutID,
		WorkoutDate:         req.WorkoutDate,
		WorkoutDuration:     req.WorkoutDuration,
		Notes:               req.Notes,
		FatigueLevel:        req.FatigueLevel,
		TotalSets:           req.TotalSets,
		TotalDistance:       req.TotalDistance,
		TotalRepetitions:    req.TotalRepetitions,
		TotalWeightLifted:   req.TotalWeightLifted,
		TotalCaloriesBurned: req.TotalCaloriesBurned,
	}

	workout, err := server.store.UpdateWorkout(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, workout)
}
