package api

import (
	"log"
	"net/http"

	db "github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createExerciseLogRequest struct {
	LogID                int64  `json:"log_id"`
	ExerciseName         string `json:"exercise_name"`
	SetsCompleted        int32  `json:"sets_completed"`
	RepetitionsCompleted int32  `json:"repetitions_completed"`
	WeightLifted         int32  `json:"weight_lifted"`
	Notes                string `json:"notes"`
}

type getExerciseLogRequest struct {
	ExerciseLogID int64 `uri:"exercise_log_id" binding:"required,min=1"`
}

type updateExerciseLogRequest struct {
	ExerciseLogID        int64  `json:"exercise_log_id"`
	SetsCompleted        int32  `json:"sets_completed"`
	RepetitionsCompleted int32  `json:"repetitions_completed"`
	WeightLifted         int32  `json:"weight_lifted"`
	Notes                string `json:"notes"`
}

func (server *Server) createExerciseLog(ctx *gin.Context) {
	var req createExerciseLogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.CreateExerciseLogParams{
		ExerciseName:         req.ExerciseName,
		LogID:                req.LogID,
		SetsCompleted:        req.SetsCompleted,
		RepetitionsCompleted: req.RepetitionsCompleted,
		WeightLifted:         req.WeightLifted,
		Notes:                req.Notes,
	}
	exercise, err := server.store.CreateExerciseLog(ctx, arg)

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

func (server *Server) getExerciseLog(ctx *gin.Context) {
	var req getExerciseLogRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	exercise, err := server.store.GetExerciseLog(ctx, req.ExerciseLogID)
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

func (server *Server) updateExerciseLog(ctx *gin.Context) {
	var req updateExerciseLogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	log.Print(req)
	arg := db.UpdateExerciseLogParams{
		ExerciseLogID:        req.ExerciseLogID,
		SetsCompleted:        req.SetsCompleted,
		RepetitionsCompleted: req.RepetitionsCompleted,
		WeightLifted:         req.WeightLifted,
		Notes:                req.Notes,
	}

	exercise, err := server.store.UpdateExerciseLog(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, exercise)
}
