package api

import (
	"net/http"

	db "github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createExerciseSetRequest struct {
	ExerciseLogID        int64 `json:"exercise_log_id"`
	SetNumber            int32 `json:"set_number"`
	WeightLifted         int32 `json:"weight_lifted"`
	RepetitionsCompleted int32 `json:"repetitions_completed"`
}

type getExerciseSetRequest struct {
	SetID int64 `uri:"set_id" binding:"required,min=1"`
}

type updateExerciseSetRequest struct {
	SetID                int64 `json:"set_id"`
	WeightLifted         int32 `json:"weight_lifted"`
	RepetitionsCompleted int32 `json:"repetitions_completed"`
}

func (server *Server) createExerciseSet(ctx *gin.Context) {
	var req createExerciseSetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.CreateExerciseSetParams{
		SetNumber:            req.SetNumber,
		ExerciseLogID:        req.ExerciseLogID,
		WeightLifted:         req.WeightLifted,
		RepetitionsCompleted: req.RepetitionsCompleted,
	}
	set, err := server.store.CreateExerciseSet(ctx, arg)

	if err != nil {
		errorCode := db.ErrorCode(err)
		if errorCode == db.ForeignKeyViolation || errorCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, set)
}

func (server *Server) getExerciseSet(ctx *gin.Context) {
	var req getExerciseSetRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	set, err := server.store.GetExerciseSet(ctx, req.SetID)
	if err != nil {
		if err == db.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	// Check if the profile being updated belongs to the authenticated user
	ctx.JSON(http.StatusOK, set)
}

func (server *Server) updateExerciseSet(ctx *gin.Context) {
	var req updateExerciseSetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.UpdateExerciseSetParams{
		SetID:                req.SetID,
		WeightLifted:         req.WeightLifted,
		RepetitionsCompleted: req.RepetitionsCompleted,
	}

	set, err := server.store.UpdateExerciseSet(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, set)
}
