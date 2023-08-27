package api

import (
	"net/http"

	db "github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAvailablePlanExerciseRequest struct {
	ExerciseName string `json:"exercise_name"`
	PlanID       int64  `json:"plan_id"`
	Sets         int32  `json:"sets"`
	RestDuration string `json:"rest_duration"`
	Notes        string `json:"notes"`
}

type getAvailablePlanExerciseRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updateAvailablePlanExerciseRequest struct {
	ID           int64  `json:"id"`
	Notes        string `json:"notes"`
	Sets         int32  `json:"sets"`
	RestDuration string `json:"rest_duration"`
}

func (server *Server) createAvailablePlanExercise(ctx *gin.Context) {
	var req createAvailablePlanExerciseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.CreateAvailablePlanExerciseParams{
		ExerciseName: req.ExerciseName,
		PlanID:       req.PlanID,
		Sets:         req.Sets,
		RestDuration: req.RestDuration,
		Notes:        req.Notes,
	}
	exercise, err := server.store.CreateAvailablePlanExercise(ctx, arg)

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

func (server *Server) getAvailablePlanExercise(ctx *gin.Context) {
	var req getAvailablePlanExerciseRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	exercise, err := server.store.GetAvailablePlanExercise(ctx, req.ID)
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

func (server *Server) updateAvailablePlanExercise(ctx *gin.Context) {
	var req updateAvailablePlanExerciseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.UpdateAvailablePlanExerciseParams{
		ID:           req.ID,
		Notes:        req.Notes,
		Sets:         req.Sets,
		RestDuration: req.RestDuration,
	}

	exercise, err := server.store.UpdateAvailablePlanExercise(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, exercise)
}
