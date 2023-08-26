package api

import (
	"errors"
	"log"
	"net/http"
	"time"

	db "github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/workout-be/token"
	"github.com/gin-gonic/gin"
)

type createWorkoutRequest struct {
	WorkoutDate     time.Time `json:"workout_date"`
	WorkoutDuration string    `json:"workout_duration"`
	Notes           string    `json:"notes"`
}

type getWorkoutRequest struct {
	WorkoutID int64 `uri:"workout_id" binding:"required,min=1"`
}

type updateWorkoutRequest struct {
	Username        string    `json:"username" binding:"required"`
	WokroutID       int64     `json:"workout_id" binding:"required,min=1"`
	WorkoutDate     time.Time `json:"workout_date"`
	WorkoutDuration string    `json:"workout_duration" binding:"min=2"`
	Notes           string    `json:"notes"`
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

	log.Print(req)
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	// Check if the profile being updated belongs to the authenticated user
	if authPayload.Username != req.Username {
		err := errors.New("profile doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(err))
		return
	}

	arg := db.UpdateWorkoutParams{
		Username:        authPayload.Username,
		WorkoutID:       req.WokroutID,
		WorkoutDate:     req.WorkoutDate,
		WorkoutDuration: req.WorkoutDuration,
		Notes:           req.Notes,
	}

	workout, err := server.store.UpdateWorkout(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, workout)
}
