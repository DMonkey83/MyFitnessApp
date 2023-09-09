package api

import (
	"errors"
	"log"
	"net/http"
	"time"

	db "github.com/DMonkey83/MyFitnessApp/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/token"
	"github.com/DMonkey83/MyFitnessApp/util"
	"github.com/gin-gonic/gin"
)

type createWorkoutRequest struct {
	Username            string            `json:"username" binding:"required"`
	WorkoutDate         time.Time         `json:"workout_date" binding:"required" `
	WorkoutDuration     string            `json:"workout_duration" binding:"required"`
	Notes               string            `json:"notes" binding:"required"`
	FatigueLevel        util.Fatiguelevel `json:"fatigue_level" binding:"required,fatigue_level"`
	TotalCaloriesBurned int32             `json:"total_calories_burned" binding:"required"`
	TotalDistance       int32             `json:"total_distance" binding:"required"`
	TotalRepetitions    int32             `json:"total_repetitions" binding:"required"`
	TotalSets           int32             `json:"total_sets" binding:"required"`
	TotalWeightLifted   int32             `json:"total_weight_lifted" binding:"required"`
}

type getWorkoutRequest struct {
	WorkoutID int64 `uri:"workout_id" binding:"required,min=1"`
}

type updateWorkoutRequest struct {
	WorkoutID           int64             `json:"workout_id"`
	WorkoutDate         time.Time         `json:"workout_date"`
	WorkoutDuration     string            `json:"workout_duration"`
	Notes               string            `json:"notes"`
	FatigueLevel        util.Fatiguelevel `json:"fatigue_level" binding:"fatigue_level"`
	TotalSets           int32             `json:"total_sets"`
	TotalDistance       int32             `json:"total_distance"`
	TotalRepetitions    int32             `json:"total_repetitions"`
	TotalWeightLifted   int32             `json:"total_weight_lifted"`
	TotalCaloriesBurned int32             `json:"total_calories_burned"`
}

type listWorkoutsRequest struct {
	Limit    int32  `form:"limit" binding:"required,min=1"`
	Offset   int32  `form:"offset" binding:"required,min=5,max=10"`
	Username string `form:"username" binding:"required,min=1"`
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
		FatigueLevel:        db.Fatiguelevel(req.FatigueLevel),
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

func (server *Server) listWorktouts(ctx *gin.Context) {
	var req listWorkoutsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	// Check if the profile being updated belongs to the authenticated user
	if authPayload.Username != req.Username {
		log.Println("auth", authPayload.Username, req.Username)
		log.Println("req", ctx.Param("username"))
		err := errors.New("profile doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(err))
		return
	}
	arg := db.ListWorkoutsParams{
		Limit:    req.Limit,
		Offset:   (req.Offset - 1) * req.Limit,
		Username: authPayload.Username,
	}
	exercises, err := server.store.ListWorkouts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}

	ctx.JSON(http.StatusOK, exercises)
}
